package main

import (
	"fmt"
	"reflect"
)

type MyFloat float64
type Point struct {
	X, Y int
}

// Функция determineType определяет и выводит тип переменной v типа interface{}.
func determineType(v interface{}) {
	fmt.Println(reflect.TypeOf(v).Kind())
}

func main() {
	types := []interface{}{42, "Hello World!", true, MyFloat(1), Point{0, 0}}
	for _, val := range types {
		determineType(val)
	}
}
