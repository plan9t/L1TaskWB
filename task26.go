// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

package main

import (
	"fmt"
	"strings"
)

func main() {
	testString1 := "abcd"
	testString2 := "abCdefAaf"
	testString3 := "aabcd"

	result1 := areAllCharactersUnique(testString1)
	result2 := areAllCharactersUnique(testString2)
	result3 := areAllCharactersUnique(testString3)

	fmt.Printf("%s — %t\n", testString1, result1)
	fmt.Printf("%s — %t\n", testString2, result2)
	fmt.Printf("%s — %t\n", testString3, result3)
}

// areAllCharactersUnique возвращает true, если все символы в строке уникальны, иначе false
func areAllCharactersUnique(s string) bool {
	charSet := make(map[rune]bool)

	// Приводим строку к нижнему регистру перед проверкой уникальности
	lowercaseString := strings.ToLower(s)

	for _, char := range lowercaseString {
		if charSet[char] {
			return false // символ уже встречался, строка не уникальна
		}
		charSet[char] = true
	}

	return true // все символы уникальны
}
