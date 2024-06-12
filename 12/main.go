package main

import "fmt"

func main() {
	var (
		seq = []string{"cat", "cat", "dog", "cat", "tree"}
		set = make(map[string]struct{})
	)
	for _, key := range seq {
		set[key] = struct{}{}
	}
	fmt.Println(set)
}
