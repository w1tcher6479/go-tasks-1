package main

import (
	"fmt"
	"slices"
)

// Функция remove принимает слайс s и индекс i, заменяет элемент в позиции i последним элементом слайса и возвращает слайс без последнего элемента.
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	slice = slices.Delete(slice, 1, 2)
	fmt.Println(slice) // [1 3 4 5]

	slice = remove(slice, 1)
	fmt.Println(slice) // [1 5 4]
}
