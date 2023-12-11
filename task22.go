// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация огромных чисел a и b
	a := new(big.Int)
	b := new(big.Int)

	// Установка значений для a и b
	a.SetString("10000000000000000000000000000000004", 10)
	b.SetString("10000000000000000000000000000000002", 10)

	// Умножение
	mulResult := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s\n", mulResult.String())

	// Деление
	divResult := new(big.Int).Div(a, b)
	fmt.Printf("Деление: %s\n", divResult.String())

	// Сложение
	addResult := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s\n", addResult.String())

	// Вычитание
	subResult := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s\n", subResult.String())
}
