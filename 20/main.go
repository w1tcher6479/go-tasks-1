package main

import (
	"fmt"
	"strings"
)

// Функция reverseWords принимает строку, переворачивает порядок слов и возвращает результат в виде строки.
func reverseWords(str string) string {
	// Разделить строку на слова
	words := strings.Fields(str)
	// Перевернуть порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// Собрать слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun hello world hi"
	fmt.Println(reverseWords(str))
}
