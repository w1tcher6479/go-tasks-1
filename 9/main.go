package main

import "fmt"

// Функция Read принимает произвольное количество целых чисел и возвращает канал, из которого можно читать эти числа.
func Read(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

// Функция Power принимает канал с целыми числами и возвращает канал, в который будут записываться квадраты этих чисел.
func Power(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	ch := Read(1, 2, 3, 4, 5)
	pow := Power(ch)
	for v := range pow {
		fmt.Print(v, " ")
	}
}
