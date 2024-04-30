package main

import "fmt"

func main() {
	name := "John"
	ptr := &name

	fmt.Println("Value", *ptr)
}
