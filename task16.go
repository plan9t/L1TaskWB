// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
)

func main() {
	arr := []string{"c", "a", "b"}
	fmt.Println(quickSort(arr))

}

func quickSort(arr []string) []string {
	if len(arr) < 2 { // Если длина массива меньше 2, то массив уже отсортирован, и функция возвращает его без изменений
		return arr
	}

	pivot := arr[0] // Объявление пивота первым элементом массива

	var less, greater []string //  Объявление двух пустых срезов, less для чисел меньше или равных пивоту, и greater для чисел больше пивота

	for _, num := range arr[1:] { // Второй элемент массива, т.к первый пивот
		if num <= pivot { //
			less = append(less, num) // Разделение элементов на два среза: меньшие или равные пивоту (less) и большие пивота (greater)
		} else { //
			greater = append(greater, num) //
		}
	}

	result := append(quickSort(less), pivot)       //  Рекурсивный вызов quickSort для среза less, после чего к результату добавляется пивот
	result = append(result, quickSort(greater)...) // К результату добавляется результат рекурсивного вызова quickSort для среза greater.
	// "три точки" (...) распаковывают элементы среза quickSort(greater) и передают их
	//в функцию append вместо создания нового среза
	return result
}
