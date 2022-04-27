package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}

	if age <= 0 {
		return nil, fmt.Errorf("age 不能小雨0")
	}

	return &Person{name, age}, nil
}

func main() {
	person, error := NewPerson("zhw", 25)
	if error == nil {
		fmt.Println(person)
	}
}
