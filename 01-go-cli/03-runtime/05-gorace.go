package main

import (
	"fmt"
	"sync"
)

var j int

var wg sync.WaitGroup

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println(j)
}

func increment() {
	j++
	wg.Done()
}
