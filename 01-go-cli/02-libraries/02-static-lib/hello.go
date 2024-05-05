package main

import (
	"C"
	"fmt"
)

func main() {}

//export Hello
func Hello() {
	fmt.Println("Hello from Go Static Library")
}
