package main

import (
	"fmt"
	"math"
)

//Структура точки
type Point struct {
	x, y int
}

//Создание новой точки
func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

//Вычисления расстояния между двумя точками
func DistnaceBetweenPoints(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	distance := DistnaceBetweenPoints(point1, point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
