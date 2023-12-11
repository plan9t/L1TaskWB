// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// СОздаём канал
	ch := make(chan int)

	// Запрос количества воркеров
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	// Создание WaitGroup для отслеживания завершения работы воркеров
	var wg sync.WaitGroup

	// Ожидание сигнала завершения программы Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Главный поток записывает данные в канал
	go func() {
		for i := 1; ; i++ {

			select {
			case ch <- i:
				// fmt.Println("Данные записаны: ", i) // Debug
				// Данные успешно записаны в канал
			case <-c:
				// Получен сигнал завершения, закрытие канала и завершение программы
				close(ch)
				fmt.Println("Канал CH закрыт!")
				wg.Wait() // Ждём пока все воркеры отработают и довыведут данные
				os.Exit(0)
			}
		}
	}()

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Добавляем 1 к счётчику
		go worker(i, &wg, ch)
	}

	// Ожидание завершения всех воркеров
	wg.Wait()
}

// Воркер
func worker(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for {

		// Чтение данных из канала
		data, ok := <-ch
		if !ok {
			// Канал закрыт, выход из цикла
			fmt.Printf("Воркер %d завершил работу\n", id)
			return
		}
		// time.Sleep(time.Millisecond)
		// Вывод данных в stdout
		fmt.Printf("Воркер %d: %d\n", id, data)
		time.Sleep(time.Millisecond * 300) //
	}
}
