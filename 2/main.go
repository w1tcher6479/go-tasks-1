package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, value := range array {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			fmt.Println(value * value)
		}(value)
	}
	wg.Wait()
}
