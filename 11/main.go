package main

import (
	"fmt"
)

// Функция SimpleGeneric принимает два слайса типа T и возвращает слайс с элементами, которые присутствуют в обоих слайсах.
func SimpleGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

// Функция HashGeneric принимает два слайса типа T и возвращает слайс с элементами, которые присутствуют в обоих слайсах.
func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

// Функция containsGeneric проверяет, содержит ли слайс b элемент e.
func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(SimpleGeneric([]int{11, 12, 13}, []int{11, 14, 12, 16, 9}))
	fmt.Println(HashGeneric([]int{11, 12, 13}, []int{11, 14, 12, 16, 9}))
}
