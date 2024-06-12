package main

import "fmt"

// Устанавливает бит в позиции pos в числе n.
func setBit(n int64, pos uint) int64 {
	return n | (1 << pos)
}

// Сбрасывает бит в позиции pos в числе n.
func clearBit(n int64, pos uint) int64 {
	return n &^ (1 << pos)
}

func main() {
	var num int64 = 10
	fmt.Printf("%b\n", num)
	fmt.Printf("%b\n", setBit(num, 0))
	fmt.Printf("%b\n", clearBit(num, 0))
}
