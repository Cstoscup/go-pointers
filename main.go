package main

import "fmt"

func main() {
	var name = "John"
	var ptr *string

	// assign the memory address of name to the pointer
	ptr = &name

	fmt.Println("Value", *ptr)
}
