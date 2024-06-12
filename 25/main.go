package main

import (
	"time"
)

func Sleep(x int) {
	if x <= 0 {
		return
	}
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	Sleep(5)
}
