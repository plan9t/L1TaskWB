// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	// Изначальные числа
	a, b := 6, 9

	// Вывод до замены мест
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Основной алгоритм
	a = a + b // a = 15
	b = a - b // b = 6
	a = a - b // a = 9

	// Вывод
	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

	// Решение через XOR
	fmt.Println("Решение через исключающее или (xor)")
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

}
