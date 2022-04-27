package oop

import "fmt"

type Man struct {
	name string
	age  int
}

func NewMan(name string, age int) (*Man, error) {
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}

	if age <= 0 {
		return nil, fmt.Errorf("age 不能小雨0")
	}

	return &Man{name, age}, nil
}
