//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Примеры переменных различных типов
	checkType(42)
	checkType("О ноу, поко икс 3......boom.")
	checkType(true)

	// Пример переменной типа канал
	myChannel := make(chan int)
	checkType(myChannel)
}

func checkType(value interface{}) {
	// Получаем тип переменной в интерфейсе
	t := reflect.TypeOf(value)

	// Выводим информацию о типе
	fmt.Printf("Значение: %v\n", value)
	fmt.Printf("Тип: %v\n", t)

	// Проверяем конкретные типы и выводим соответствующее сообщение
	// Возможно, что тип переменной может быть интерфейсом, который реализует несколько типов
	switch value.(type) {
	case int:
		fmt.Println("Это int.")
	case string:
		fmt.Println("Это string.")
	case bool:
		fmt.Println("Это bool.")
	case chan int:
		fmt.Println("Это канал int.")
	default:
		fmt.Println("Не удалось определить тип.")
	}
	fmt.Println()
}
