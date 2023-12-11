// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode

package main

import "fmt"

func main() {
	inputString := "plan9t"

	// Печать исходной строки
	fmt.Println("Исходная строка:", inputString)

	// Переворачиваем строку
	reversedString := reverseString(inputString)

	// Печать перевернутой строки
	fmt.Println("Перевернутая строка:", reversedString)
}

func reverseString(input string) string {
	// Преобразуем строку в руны
	runes := []rune(input) // rune - 32-битное число и предназначен для хранения одиночного символа Unicode

	// Длина строки
	length := len(runes)

	// Переворачиваем строку, меняя местами символы от начала к концу
	// Используются два указателя i и j, которые идут навстречу друг другу
	// Каждый символ в позиции i меняется местами с символом в позиции j
	// Цикл продолжается, пока i не станет больше или равно j
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
