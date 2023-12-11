//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

package main

import (
	"fmt"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для хранения групп температур
	temperatureGroups := make(map[int][]float64)

	// Группируем температуры по первой цифре
	for _, temp := range temperatures {
		// Определяем группу по первой цифре
		firstDigitGroup := (int(temp) / 10) * 10
		temperatureGroups[firstDigitGroup] = append(temperatureGroups[firstDigitGroup], temp)
	}

	// Выводим результат
	fmt.Println(temperatureGroups)
}
