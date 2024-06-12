package main

import (
	"fmt"
	"sync"
)

var (
	m    = map[int]int{}
	lock = sync.Mutex{}
)

// Функция write записывает значение value по ключу key в мапу m.
func write(key int, value int) {
	lock.Lock()
	defer lock.Unlock()
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			write(i, i)
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}
