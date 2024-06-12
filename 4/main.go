package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Функция PrintFromChannel получает значения из канала и выводит их.
func PrintFromChannel(ch <-chan int, workerID int) {
	for v := range ch {
		fmt.Printf("Получено число %d от воркера %d\n", v, workerID)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int)

	var N int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&N)

	handler := make(chan os.Signal, 1)
	signal.Notify(handler, syscall.SIGINT, os.Interrupt)

	for i := 0; i < N; i++ {
		go PrintFromChannel(ch, i)
	}

	for {
		select {
		case <-handler:
			close(ch)
			close(handler)
			return
		default:
			ch <- rand.Intn(10)
		}
	}
}
