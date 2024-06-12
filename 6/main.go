package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Функция manageWithChannels демонстрирует использование каналов для управления выполнением горутины.
func manageWithChannels() {
	fmt.Println("Канал")

	ch := make(chan struct{})

	go func() {
		defer fmt.Println("Горутина остановлена.")
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина выполняется:", i)
			time.Sleep(time.Second)
		}
	}()

	<-ch
}

// Функция manageWithContext демонстрирует использование контекста для управления выполнением горутины.
func manageWithContext() {
	fmt.Println("Контекст")
	ctx, cancel := context.WithCancel(context.Background())

	defer fmt.Println("Горутина остановлена.")
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина выполняется...")
				time.Sleep(time.Second)
			}
		}
	}()

	fmt.Println("Нажмите Enter, чтобы остановить горутину.")
	fmt.Scanln()
}

// Функция manageWithWaitGroup демонстрирует использование WaitGroup для управления выполнением горутины.
func manageWithWaitGroup() {
	fmt.Println("WaitGroup")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина выполняется:", i)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("\nГорутина остановлена.")
}

func main() {
	manageWithChannels()
	manageWithContext()
	manageWithWaitGroup()
}
