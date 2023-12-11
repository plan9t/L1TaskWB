//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	withChannel()
	contextClose()
	wgClose()
	selectDefaultClose()
	runtimeGoexitClose()
	flagClose()
	timeAfterClose()
	panicClose() // panic последняя потому что exit code 2
}

// 1. Закрытие через канал (способ из предыдущих заданий)
func withChannel() {
	stopperChannel := make(chan bool)

	go func() {
		for {
			select {
			case <-stopperChannel:
				fmt.Println("Горутина завершена.")
				return
			default:
				fmt.Println("Выполняется работа в горутине.")
				time.Sleep(time.Second)
			}
		}
	}()

	// Время для работы горутины
	time.Sleep(3 * time.Second)

	// Останавливаем горутину, отправляя true канал
	stopperChannel <- true

	// Время для завершения горутины
	time.Sleep(1 * time.Second)

}

// Закрытие через контекст
func contextClose() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена.")
				return
			default:
				fmt.Println("Выполняется работа в горутине.")
				time.Sleep(time.Second)
			}
		}
	}()

	// Даем горутине работать в течение некоторого времени
	time.Sleep(3 * time.Second)

	// Останавливаем горутину, вызывая cancel()
	cancel()

	// Даем немного времени для завершения горутины
	time.Sleep(1 * time.Second)
}

// waitGroup
func wgClose() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done() // После работы горутины дикремент счётчика
			// Работа горутины
		}()
		wg.Add(1)
	}

	// Главная горутина блокируется до завершения всех горутин
	wg.Wait()
	fmt.Println("Все горутины завершены.")

}

func selectDefaultClose() {
	done := make(chan bool)

	go func() {
		// Работа горутины
		time.Sleep(2 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Горутина завершена.")
	default:
		fmt.Println("Выполняется работа в основной горутине.")
	}
}

// Закрытие через рантайм Goexit. Goexit завершает работу конкретной корутины в которой вызывется этот метод
func runtimeGoexitClose() {
	go func() {
		// Работа горутины
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина завершена.")
		runtime.Goexit()
	}()

	// Выполняется работа в основной горутине
	time.Sleep(1 * time.Second)
	fmt.Println("Выполняется работа в основной горутине после 1 секунды.")
}

func flagClose() {

	var doneFlag bool
	var mutex sync.Mutex

	go func() {
		for {
			mutex.Lock()
			if doneFlag {
				mutex.Unlock()
				fmt.Println("Горутина завершена.")
				return
			}
			mutex.Unlock()

			fmt.Println("Выполняется работа в горутине.")
			time.Sleep(time.Second)
		}
	}()

	// Даем горутине работать в течение некоторого времени
	time.Sleep(3 * time.Second)

	// Устанавливаем флаг для завершения горутины
	mutex.Lock()
	doneFlag = true
	mutex.Unlock()

	// Даем немного времени для завершения горутины
	time.Sleep(1 * time.Second)
}

// panic close (тоже способ, хоть и не очень)
func panicClose() {
	go func() {
		defer func() {
			fmt.Println("Отложенная функция перед завершением горутины.")
		}()

		for {
			fmt.Println("Выполняется работа в горутине.")
			time.Sleep(time.Second)
		}
	}()

	// Даем горутине работать в течение некоторого времени
	time.Sleep(3 * time.Second)

	// Завершаем выполнение горутины с помощью panic
	panic("Горутина завершена.")
}

// timeAfter close
func timeAfterClose() {
	go func() {
		for {
			select {
			case <-time.After(3 * time.Second):
				fmt.Println("Горутина завершена.")
				return
			default:
				fmt.Println("Выполняется работа в горутине.")
				time.Sleep(time.Second)
			}
		}
	}()

	// Даем горутине работать

	// Даем немного времени для завершения горутины
	time.Sleep(5 * time.Second)
}
