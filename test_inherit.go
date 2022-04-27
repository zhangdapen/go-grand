package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (animal Animal) eat() {
	fmt.Println("eat")
}

func (animal Animal) sleep() {
	fmt.Println("sleep")
}

func (animal Animal) work() {
	fmt.Println("work")
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
	say string
}

func main() {

	cat := Cat{
		Animal{"中华田园猫", 1},
	}

	dog := Dog{
		Animal{"藏獒", 2},
		"汪汪汪",
	}

	cat.eat()
	cat.sleep()
	cat.work()

	dog.eat()
	dog.sleep()
	dog.work()
	fmt.Println(dog.say)
}
