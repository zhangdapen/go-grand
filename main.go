package main

import (
	"fmt"
	"go-grand/oop"
)

func main() {
	testConstruct()
	testInherit()
	testOop()
	testInterface()
}

func testConstruct() {
	man, error := oop.NewMan("zhw", 18)
	if error == nil {
		fmt.Println(*man)
	}
}

func testInherit() {
	bird := oop.Bird{
		oop.Animal{"乌鸦", 1},
	}

	snake := oop.Snake{
		oop.Animal{"菜花蛇", 1},
		"sisisi",
	}

	bird.Eat()
	bird.Sleep()
	bird.Work()
	snake.Eat()
	snake.Sleep()
	snake.Work()
}

func testOop() {
	woman := oop.Woman{"htt", 1}
	woman.Eat()
	woman.Sleep()
	woman.Work()
}

func testInterface() {
	person := oop.Person{"zhw"}
	person.Play(oop.Cat{"中华田园猫", 1})
	person.Play(oop.Dog{"哈士奇", 2})
}
