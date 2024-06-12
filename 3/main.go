package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	array := []int64{2, 4, 6, 8, 10}
	var (
		sum int64
		wg  sync.WaitGroup
	)
	for _, value := range array {
		wg.Add(1)
		go func(value int64) {
			defer wg.Done()
			atomic.AddInt64(&sum, value*value)
		}(value)
	}
	wg.Wait()
	fmt.Println(sum)
}
