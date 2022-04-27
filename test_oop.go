package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) eat() {
	fmt.Println("eat....")
}

func (person Person) sleep() {
	fmt.Println("sleep")
}

func (person Person) work() {
	fmt.Println("word")
}

func main() {
	person := Person{"zhw", 2}
	person.eat()
	person.sleep()
	person.work()
}
