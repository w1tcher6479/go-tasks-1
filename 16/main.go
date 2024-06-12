package main

import "fmt"

// Функция partition разбивает массив на две части относительно опорного элемента (pivot) и возвращает модифицированный массив и индекс опорного элемента.
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high] // выбираем последний элемент в качестве опорного
	i := low           // индекс для отслеживания места, куда поместить следующий элемент меньше pivot

	// проходим по массиву и перемещаем элементы, меньшие pivot, в начало
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

// Функция quickSort рекурсивно сортирует массив с помощью алгоритма быстрой сортировки.
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high) // разбиваем массив на две части
		arr = quickSort(arr, low, p-1)     // рекурсивно сортируем левую часть
		arr = quickSort(arr, p+1, high)    // рекурсивно сортируем правую часть
	}
	return arr
}

// Функция quickSortStart запускает быструю сортировку на всем массиве.
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	var array = []int{57, 32, 11, 87, 34, -51, -20, -13, 0, 39}
	fmt.Println(quickSortStart(array))
}
