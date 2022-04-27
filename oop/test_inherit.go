package oop

/**
继承
*/

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (animal Animal) Eat() {
	fmt.Println("eat")
}

func (animal Animal) Sleep() {
	fmt.Println("sleep")
}

func (animal Animal) Work() {
	fmt.Println("work")
}

type Bird struct {
	Animal
}

type Snake struct {
	Animal
	Say string
}
