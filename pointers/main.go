package main

import "fmt"

type dog struct {
	name  string
	breed string
	age   int
}

func main() {
	var dog1 dog = dog{"Fido", "German Shepherd", 5}
	fmt.Println(dog1)

	changeDogAge(&dog1)
	fmt.Println(dog1)
}

func changeDogAge(dog *dog) {
	dog.age = dog.age + 1
}
