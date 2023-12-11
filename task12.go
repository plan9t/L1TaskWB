package main

import "fmt"

func main() {
	// Создаем множество строк
	stringSet := make(map[string]bool)

	// Последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Добавляем каждую строку в множество
	for _, str := range sequence {
		stringSet[str] = true
	}

	// Выводим содержимое множества
	fmt.Println("Множество строк:", stringSet)

	// Пример проверки принадлежности строки к множеству
	checkString := "dog"
	if stringSet[checkString] {
		fmt.Printf("%s принадлежит множеству\n", checkString)
	} else {
		fmt.Printf("%s не принадлежит множеству\n", checkString)
	}
}
