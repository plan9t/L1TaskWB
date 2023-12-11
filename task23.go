// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	// Пример слайса
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	indexToRemove := 2

	// Проверка на корректный индекс
	if indexToRemove >= 0 && indexToRemove < len(slice) {
		// Удаление элемента. От начала и до(не включая) indexToRemove, после добаляются все элементы от indexToRemove+1
		// ... разворачивает слайс (для append)
		slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)

		// Output
		fmt.Println("Слайс после удаления:", slice)
	} else {
		fmt.Println("Некорректный индекс для удаления")
	}
}
