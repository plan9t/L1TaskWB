// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "snow dog sun — sun dog snow"

	// Печать исходной строки
	fmt.Println("Исходная строка:", inputString)

	// Переворачиваем слова в строке
	reversedString := reverseWords(inputString)

	// Печать результата
	fmt.Println("Перевернутые слова:", reversedString)
}

func reverseWords(input string) string {
	words := strings.Fields(input) // Разделяем строку на слова
	reversedWords := make([]string, len(words))

	// Переворачиваем каждое слово
	for i, word := range words {
		reversedWords[len(words)-1-i] = word // Вставляем перевернутые слова в обратном порядке
	}

	// Собираем перевернутую строку из перевернутых слов
	reversedString := strings.Join(reversedWords, " ")

	return reversedString
}
