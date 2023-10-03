package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2<<20 + 10)
	b := big.NewInt(2<<20 + 5)

	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Printf("Сумма: %d\n", big.NewInt(0).Add(a, b))

	fmt.Printf("Разность: %d\n", big.NewInt(0).Sub(a, b))

	fmt.Printf("Произведение: %d\n", big.NewInt(0).Mul(a, b))

	fmt.Printf("Деление: %d\n", big.NewInt(0).Div(a, b))
}
