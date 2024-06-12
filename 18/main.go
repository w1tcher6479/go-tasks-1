package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

// Функция Increment увеличивает значение счетчика на 1.
func (c *Counter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

// Value возвращает текущее значение счетчика.
func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter.Value())
}
