package decimalwrap2

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func TFunc() {
	v, _ := decimal.NewFromString("-123.4567")
	fmt.Printf("v1.2.0: %s\n", v)
}
