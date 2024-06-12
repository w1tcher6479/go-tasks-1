package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

var justString string

/*
Во время выполнения функции someFunc() создается строка размером в 1 КБ.
Затем из этой строки берутся только первые 100 байт и присваиваются глобальной переменной justString.
Но сама исходная строка остается в памяти и не может быть освобождена, так как на нее все еще ссылается глобальная переменная justString.

Чтобы исправить это, нужно скопировать часть строки в новую переменную.
*/
func someFunc() {
	v := createHugeString(1 << 10)

	byteSlice := make([]byte, 100)
	copy(byteSlice, v[:100])
	v = ""

	justString = string(byteSlice)
	fmt.Println(justString)
}

func main() {
	someFunc()
}
