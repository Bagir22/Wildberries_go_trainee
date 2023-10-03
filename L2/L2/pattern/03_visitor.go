package main

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Посетитель позволяет добавлять в программу новые операции, не изменяя классы объектов, над которыми эти операции могут выполняться.

	+
	1. Упрощает добавление операций, работающих со сложными структурами объектов.
	-
	1. Может привести к нарушению инкапсуляции элементов.
*/

import "fmt"

type AnimalVisitor interface {
	VisitDog(dog *Dog)
	VisitCat(cat *Cat)
}

type SoundVisitor struct{}

func (sv *SoundVisitor) VisitDog(dog *Dog) {
	fmt.Println("Dog says: Woof!")
}

func (sv *SoundVisitor) VisitCat(cat *Cat) {
	fmt.Println("Cat says: Meow!")
}

type Animal interface {
	Accept(visitor AnimalVisitor)
}

type Dog struct{}

func (d *Dog) Accept(visitor AnimalVisitor) {
	visitor.VisitDog(d)
}

type Cat struct{}

func (c *Cat) Accept(visitor AnimalVisitor) {
	visitor.VisitCat(c)
}

func main() {
	animals := []Animal{&Dog{}, &Cat{}}

	soundVisitor := &SoundVisitor{}
	for _, animal := range animals {
		animal.Accept(soundVisitor)
	}
}
