package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := big.NewInt(0).SetString("1000000000000000000000", 10)
	b, _ := big.NewInt(0).SetString("1000000000000000000000", 10)

	calculate := big.NewInt(0)
	fmt.Println("Sum: ", calculate.Add(a, b))
	fmt.Println("Sub: ", calculate.Sub(a, b))
	fmt.Println("Mul: ", calculate.Mul(a, b))
	fmt.Println("Div: ", calculate.Div(a, b))
}
