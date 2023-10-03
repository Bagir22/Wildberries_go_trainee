package main

import "fmt"

//Родительская структура
type Human struct {
	name string
	age int
}

//Методы Human
func (h *Human) printName() {
	fmt.Printf("He's name is %v\n", h.name)
}

func (h *Human) printAge() {
	fmt.Printf("%v is %v years old\n", h.name, h.age)
}

//Дочерняя структура
type Action struct {
	Human
}

func main() {
	h := Action {
		Human {
			"Danila", 21,
		},
	}

	h.printName()
	h.printAge()

	/*
		Дочерняя стуктура на прямую не реализует методы Human, но встаривает эту структуру, посе чего структуре Action доступны методы Human 
	*/
}
