// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

func main() {
	// Создание двух точек
	point1 := NewPoint(0.0, 0.0)
	point2 := NewPoint(47.0, 0.0)

	// Вычисление расстояния между точками
	distance := Distance(point1, point2)

	// Вывод результата
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

// Point представляет структуру для хранения координат точки.
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance вычисляет расстояние между двумя точками. По теореме Пифагора
func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}
