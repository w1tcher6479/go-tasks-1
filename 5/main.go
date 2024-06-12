package main

import (
	"fmt"
	"time"
)

const N = 3

// Функция send отправляет в канал целые числа, начиная с 0, с интервалом в 500 миллисекунд.
func send(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}

// Функция receive читает из канала и выводит значения на экран.
func receive(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	timeoutCh := time.After(N * time.Second)

	go send(ch)
	go receive(ch)

	<-timeoutCh

	fmt.Println("Программа завершилась.")
}
