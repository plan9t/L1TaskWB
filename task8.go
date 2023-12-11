// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func main() {
	var number int64
	var position uint
	var value uint

	// Ввод данных
	fmt.Print("Введите число (int64): ")
	fmt.Scan(&number)

	fmt.Print("Введите позицию бита: ")
	fmt.Scan(&position)

	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&value)

	// Установить бит
	result := setBit(number, position, value)

	// Вывести результат
	fmt.Printf("Результат: %d\n", result)
}

func setBit(number int64, position uint, value uint) int64 {
	// Очиститка бита на position
	mask := ^(1 << position) // Все биты в 1 кроме position
	number &= int64(mask)    //  Все биты number не меняются кроме бита на позиции position. Он сбрасывается в 0

	// Установка бита в новое значение
	number |= int64(value) << position // Бит на позиции position устанавливается в value. Все ост. остаются

	return number
}
