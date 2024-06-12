package main

import (
	"fmt"
)

// Функция BinarySearch ищет элемент в отсортированном в порядке возрастания слайсе, состоящем только из целочисленных элементов, и возвращает индекс, в котором элемент найден.
// Функция также возвращает bool, указывающий, действительно ли элемент найден в этом слайсе.
func BinarySearch(slice []int, target int) (int, bool) {
	left, right := 0, len(slice)
	for left < right {
		mid := (left + right) >> 1
		if slice[mid] == target {
			return mid, true
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right, right < len(slice) && (slice[right] == target)
}

func main() {
	slice := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	fmt.Println(BinarySearch(slice, 56))
}
