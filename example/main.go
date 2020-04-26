package main

import (
	"fmt"

	bigNumber "github.com/LiMoMoMo/bigNumber"
)

func main() {
	number1, err := bigNumber.New(bigNumber.DECIMAL, "123412341234")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(number1.BinaryStr(), number1.DecimalStr())

	number2, err := bigNumber.New(bigNumber.BINARY, "100011000100110110110111111100110100001001101001110")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(number2.BinaryStr(), number2.DecimalStr())
}
