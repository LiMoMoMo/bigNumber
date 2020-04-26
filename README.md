# bigNumber
Golang big number.
You can use this repository handle big number.

# Use
## Decimal to Binary
```
    import (
        bigNumber "github.com/LiMoMoMo/bigNumber"
        "fmt"
    )
    number, err := bigNumber.New(bigNumber.DECIMAL, "123412341234")
    fmt.Println(number.BinaryStr(), number.DecimalStr())
```

## Binary to Decimal
```
    import (
        bigNumber "github.com/LiMoMoMo/bigNumber"
        "fmt"
    )
    number, err := bigNumber.New(bigNumber.BINARY, "100011000100110110110111111100110100001001101001110")
    fmt.Println(number.BinaryStr(), number.DecimalStr())
```

# reference
[double-dabble](https://en.wikipedia.org/wiki/Double_dabble)
