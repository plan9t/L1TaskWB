// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("START")

	// Использование собственной функции sleep
	sleep(3)

	fmt.Println("END")
}

// sleep функция задержки, принимает количество секунд для ожидания
func sleep(seconds int) {
	startTime := time.Now()                                        // Время старта = текущее время
	endTime := startTime.Add(time.Duration(seconds) * time.Second) // К времени старта добавляется указаная задержка. Используется в цикле

	for time.Now().Before(endTime) { // Цикл для ожидания указанного количества секунд
		// ожидание.
	}
}
