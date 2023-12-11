//1. Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import (
	"fmt"
	"strconv"
)

// Базовая структура Human
type Human struct {
	firstName string
	lastName  string
	age       int
	country   string
}

// Второстепенная структура Action, в которой мы реализуем наследование Human
type Action struct {
	Human              // Встраиваем "родительскую" структуру
	Description string // Какое-то описание
}

// Метод для Human
func (h Human) sayInfo() string {
	return "Hi! My name is " + h.firstName + " " + h.lastName + ". I am " + strconv.Itoa(h.age) + " years old. I from " + h.country + "!"
}

// Доп. метод лля Action. Здесь используется поле из родительской струтуры
func (a Action) doAction() string {
	return "I am " + a.firstName + "! I do it!!!"
}

func main() {
	// Создание экземпляра структуры
	me := Human{
		firstName: "Artyom",
		lastName:  "Vdovenko",
		age:       23,
		country:   "Russia",
	}

	// Создание дочернего от Human экземпляра структуры Action
	myAction := Action{
		Human:       me,
		Description: "Action description",
	}

	// Тестовые выводы
	fmt.Println(me.sayInfo())
	fmt.Println(myAction.doAction())

}
