package main

import "fmt"

func main() {

	d := Dog{Name: "Rufus"}
	c := Cat{Name: "Meowington"}

	displayAnimal(d)
	displayAnimal(c)
}

type Dog struct {
	Name string
}

func (d Dog) makeSound() string {
	return "Dog barks"
}

type Cat struct {
	Name string
}

func (c Cat) makeSound() string {
	return "Cat meows"
}

type Animal interface {
	makeSound() string
}

func displayAnimal(a Animal) {
	fmt.Println(a.makeSound())
}
