package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// Функция EuclidianDistance вычисляет евклидово расстояние между двумя точками p1 и p2.
func EuclidianDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}

func main() {
	point1 := NewPoint(1.3, 2.5)
	point2 := NewPoint(2.7, 6.9)
	fmt.Println(EuclidianDistance(point1, point2))
}
