package main

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	«цепочка вызовов» позволяет избежать жесткой привязки отправителя запроса к получателю. 
	Все возможные обработчики запроса образуют цепочку, а сам запрос перемещается по этой цепочке.

	+
	1. Паттерн позволяет отделить отправителя запроса от получателя, что уменьшает зависимость между ними и делает код более гибким и расширяемым.
	-
	1. Нет гарантии, что запрос будет обработан каким-либо обработчиком. Если ни один из обработчиков не может обработать запрос, он останется необработанным.
*/

import "fmt"

type Inspector interface {
	CheckQuality(dish *Dish)
	SetNext(inspector Inspector)
}

type ChefInspector struct {
	next Inspector
}

func (ci *ChefInspector) SetNext(inspector Inspector) {
	ci.next = inspector
}

func (ci *ChefInspector) CheckQuality(dish *Dish) {
	if dish.CookingStatus == "raw" {
		fmt.Println("Chef Inspector: Dish is raw. Cooking a bit longer.")
		dish.CookingStatus = "medium"
	}
	if ci.next != nil {
		ci.next.CheckQuality(dish)
	}
}

type SousChefInspector struct {
	next Inspector
}

func (sci *SousChefInspector) SetNext(inspector Inspector) {
	sci.next = inspector
}

func (sci *SousChefInspector) CheckQuality(dish *Dish) {
	if dish.CookingStatus == "medium" {
		fmt.Println("Sous Chef Inspector: Dish is medium. Ready to serve.")
	} else if sci.next != nil {
		sci.next.CheckQuality(dish)
	}
}

type Dish struct {
	CookingStatus string
}

func main() {
	chefInspector := &ChefInspector{}
	sousChefInspector := &SousChefInspector{}

	chefInspector.SetNext(sousChefInspector)

	dish := &Dish{CookingStatus: "raw"}

	chefInspector.CheckQuality(dish)
}
