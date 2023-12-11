package main

import "fmt"

func main() {
	// Исходные множества
	set1 := []int{1, 2, 3, 4, 5, 7}
	set2 := []int{3, 5, 6, 7}

	// Вызов функции для нахождения пересечения
	result := intersection(set1, set2)

	// Вывод результата
	fmt.Println("Пересечение множеств:", result)
}

// Функция для поиска пересечения множетсв
func intersection(set1, set2 []int) []int {
	var result []int

	// Создаем карту для быстрого поиска элементов в set2
	set2Map := make(map[int]bool)
	// Проходим по каждому элементу массива set2 и добавляем его в карту set2Map в качестве ключа со значением true.
	// Так можно быстро проверять принадлежность элемента к множеству set2.
	for _, val := range set2 {
		set2Map[val] = true
	}

	// Проходим по каждому элементу массива set1 и проверяем, присутствует ли он в set2Map.
	//Если элемент присутствует, мы добавляем его в результат.
	for _, val := range set1 {
		if set2Map[val] {
			result = append(result, val)
		}
	}

	return result
}
