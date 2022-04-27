package oop

/**
类实现接口，多态
*/

import "fmt"

type Pet interface {
	Eat()
	Sleep()
}

type Dog struct {
	Name string
	Age  int
}

func (dog Dog) Eat() {
	fmt.Println("dog eat")
}

func (dog Dog) Sleep() {
	fmt.Println("dog sleep")
}

type Cat struct {
	Name string
	Age  int
}

func (cat Cat) Eat() {
	fmt.Println("cat eat")
}

func (cat Cat) Sleep() {
	fmt.Println("cat sleep")
}

type Person struct {
	Name string
}

func (person Person) Play(pet Pet) {
	pet.Eat()
	pet.Sleep()
}
