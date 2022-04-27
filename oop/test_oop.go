package oop

/**
类下面的方法
*/

import "fmt"

type Woman struct {
	Name string
	Age  int
}

func (woman Woman) Eat() {
	fmt.Println("eat....")
}

func (woman Woman) Sleep() {
	fmt.Println("sleep")
}

func (woman Woman) Work() {
	fmt.Println("work")
}
