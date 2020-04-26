package bignumber

import "errors"

type BCD struct {
	data []byte
}

func NewBCD(input []byte) (*BCD, error) {
	if len(input) != 4 {
		return nil, errors.New("input argument error")
	}
	d := BCD{
		data: input,
	}
	return &d, nil
}

func (b *BCD) Value() int {
	result := 0
	if b.data[0] == '1' {
		result += 8
	}
	if b.data[1] == '1' {
		result += 4
	}
	if b.data[2] == '1' {
		result += 2
	}
	if b.data[3] == '1' {
		result += 1
	}
	return result
}

func (b *BCD) Plus(arg int) error {
	result := b.Value() + arg
	if result > 15 {
		return errors.New("stack over flow")
	}
	for i := 0; i < 4; i++ {
		b.data[i] = '0'
	}
	for result > 0 {
		if result >= 8 {
			result -= 8
			b.data[0] = '1'
			continue
		}
		if result >= 4 {
			result -= 4
			b.data[1] = '1'
			continue
		}
		if result >= 2 {
			result -= 2
			b.data[2] = '1'
			continue
		}
		if result >= 1 {
			result -= 1
			b.data[3] = '1'
			continue
		}
	}
	return nil
}

func devide(arg string, divisor int) string {
	result := make([]byte, 0)
	idx := 0
	temp := int(arg[idx] - '0')
	for temp < divisor {
		idx++
		temp = temp*10 + int(arg[idx]-'0')
	}
	for len(arg) > idx {
		result = append(result, byte(temp/divisor+int('0')))
		idx++
		if idx >= len(arg) {
			break
		}
		temp = (temp%divisor)*10 + int(arg[idx]-'0')
	}
	if len(result) == 0 {
		return "0"
	}
	return string(result)
}

type InitType int

const (
	DECIMAL InitType = iota
	BINARY
)

type BigNumber struct {
	decimalData []byte
	binaryData  []byte
}

func (b *BigNumber) toBinary() {
	arg := string(b.decimalData)
	oddOreven := func(arg string) bool {
		last := arg[len(arg)-1]
		switch last {
		case '1':
			return false
		case '3':
			return false
		case '5':
			return false
		case '7':
			return false
		case '9':
			return false
		default:
			return true
		}
	}
	for arg != "1" {
		if oddOreven(arg) {
			b.binaryData = append([]byte{'0'}, b.binaryData...)
		} else {
			b.binaryData = append([]byte{'1'}, b.binaryData...)
		}
		arg = devide(arg, 2)
	}
	b.binaryData = append([]byte{'1'}, b.binaryData...)
}

// double-dabble.
func (number *BigNumber) toDecimal() {
	input := number.binaryData
	N := 0
	//
	a := len(input) / 4
	b := len(input) % 4
	if b != 0 {
		preadd := make([]byte, 4-b)
		for i := 0; i < 4-b; i++ {
			preadd[i] = '0'
		}
		input = append(preadd, input...)
		a++
	}
	N = len(input)

	//
	pre := make([]byte, 4*(a+(N/3)))
	for i := 0; i < 4*(a+(N/3)); i++ {
		pre[i] = '0'
	}
	input = append(pre, input...)

	// define copy function
	copy := func(src []byte, dest []byte) {
		for i := 0; i < len(src); i++ {
			dest[i] = src[i]
		}
	}

	//
	count := 0
	lens := len(input)
	for count < N {
		//
		for i := 0; i < a+(N/3); i++ {
			bcd, _ := NewBCD(input[i*4 : i*4+4])
			if bcd.Value() > 4 {
				bcd.Plus(3)
			}
		}
		//
		copy(input[1:lens], input[0:lens-1])
		input[lens-1] = '0'
		count++
	}
	//
	flag := false
	for i := 0; i < a+(N/3); i++ {
		bcd, _ := NewBCD(input[i*4 : i*4+4])
		if bcd.Value() == 0 && !flag {
			continue
		}
		flag = true
		number.decimalData = append(number.decimalData, byte(bcd.Value()+'0'))
	}
}

func (b *BigNumber) BinaryStr() string {
	return string(b.binaryData)
}

func (b *BigNumber) DecimalStr() string {
	return string(b.decimalData)
}

func New(t InitType, input string) (*BigNumber, error) {
	b := BigNumber{
		decimalData: make([]byte, 0),
		binaryData:  make([]byte, 0),
	}
	checkDecimalInput := func(arg string) bool {
		for _, v := range arg {
			if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}
	checkBinaryInput := func(arg string) bool {
		for _, v := range arg {
			if v != '0' && v != '1' {
				return false
			}
		}
		return true
	}
	switch t {
	case DECIMAL:
		if !checkDecimalInput(input) {
			return nil, errors.New("input error")
		}
		b.decimalData = []byte(input)
		b.toBinary()
	case BINARY:
		if !checkBinaryInput(input) {
			return nil, errors.New("input error")
		}
		b.binaryData = []byte(input)
		b.toDecimal()
	}
	return &b, nil
}
