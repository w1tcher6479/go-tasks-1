package main

import (
	"fmt"
	"strings"
)

// Функция isUnique принимает строку str и возвращает bool, указывающий, содержит ли строка уникальные символы.
func isUnique(str string) bool {
	// Преобразование строки к нижнему регистру
	lowerStr := strings.ToLower(str)

	// Используем множество для отслеживания встреченных символов
	charSet := make(map[rune]struct{})

	for _, char := range lowerStr {
		if _, ok := charSet[char]; ok {
			return false
		}
		charSet[char] = struct{}{}
	}
	return true
}

func main() {
	example := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, str := range example {
		fmt.Printf("%s — %t\n", str, isUnique(str))
	}
}
