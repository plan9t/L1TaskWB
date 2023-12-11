// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создание канала
	ch := make(chan int)

	// Создание группы WaitGroup для отслеживания завершения горутин
	var wg sync.WaitGroup

	// Запуск горутины для отправки данных в канал
	wg.Add(1)
	go sendData(ch, &wg)

	// Запуск горутины для считывания данных из канала
	wg.Add(1)
	go receiveData(ch, &wg)

	// Ждем N секунд
	N := 5
	time.Sleep(time.Duration(N) * time.Second)

	// Закрываем канал и ожидаем завершения горутин
	//close(ch)
	wg.Wait()

	fmt.Println("Программа завершена.")
}

// Функция для отправки данных в канал
func sendData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("Отправлены данные: ", i)
		time.Sleep(500 * time.Millisecond) // Задержка в полсекунды
	}

	close(ch) // Закрываем канал после отправки данных
	fmt.Println("Горутина для отправки завершена.")
}

// Функция для считывания данных из канала
func receiveData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		data, ok := <-ch
		if !ok {
			// Канал закрыт, выходим из цикла
			fmt.Println("Горутина для считывания данных завершена.")
			return
		}

		fmt.Println("Принято из канала:", data)
	}
}
