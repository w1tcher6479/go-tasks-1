package main

import (
	"fmt"
)

// Функция Reverse возвращает строку в обратном порядке.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(Reverse("главрыба"))
}
