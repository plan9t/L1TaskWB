//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создание каналов для записи исходных данных и результатов работы
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Горутина для записи чисел в первый канал
	wg.Add(1)
	go produceNumbers(inputChannel, &wg)

	// Горутина для умножения чисел и записи во второй канал
	wg.Add(1)
	go multiplyByTwo(inputChannel, outputChannel, &wg)

	// Горутина для вывода результата в stdout
	wg.Add(1)
	go consumeResults(outputChannel, &wg)

	// Ожидаем завершения всех горутин
	wg.Wait()
}

// Функция для записи в канал
func produceNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch) // После отрабоки функции канал закрывается
	defer wg.Done() // Счётчик -1

	numbers := []int{1, 2, 3, 4, 5} // Исходный массив с данными

	for _, num := range numbers {
		ch <- num // Запись в канал
	}
}

// Функция для умножения чисел и записи во второй канал
func multiplyByTwo(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer close(output) // Закрытие канала после отработки
	defer wg.Done()     // Счётчик -1

	for num := range input {
		result := num * 2
		output <- result // Запись в канал
	}
}

// Функция для вывода результата в stdout
func consumeResults(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Счётчик -1

	for result := range ch {
		fmt.Println(result) // Вывод в stdout
	}
}
