// Реализовать бинарный поиск встроенными методами языка

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 11))
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2 // Средний индекс входного массива

		if arr[mid] == target {
			return mid // Элемент найден, возвращаем его индекс
		} else if arr[mid] < target {
			low = mid + 1 // Искомый элемент в верхней половине
		} else {
			high = mid - 1 // Искомый элемент в нижней половине
		}
	}

	return -1 // Если такого элемента в массиве не будет
}
