package main

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Фабричный метод определяет общий интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов

	+
	1. позволяет добавлять новые типы продуктов, не изменяя существующий код. Это делает код более расширяемым и гибким.
	2. скрывает детали процесса создания объектов
	
	-
	1. Введение фабричных методов и их подклассов может увеличить сложность кода, особенно если есть много разных продуктов и фабрик.
*/

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type CheesePizza struct{}
type PepperoniPizza struct{}
type VeggiePizza struct{}

func (p *CheesePizza) Prepare() { fmt.Println("Preparing Cheese Pizza") }
func (p *CheesePizza) Bake()    { fmt.Println("Baking Cheese Pizza") }
func (p *CheesePizza) Cut()     { fmt.Println("Cutting Cheese Pizza") }
func (p *CheesePizza) Box()     { fmt.Println("Boxing Cheese Pizza") }

func (p *PepperoniPizza) Prepare() { fmt.Println("Preparing Pepperoni Pizza") }
func (p *PepperoniPizza) Bake()    { fmt.Println("Baking Pepperoni Pizza") }
func (p *PepperoniPizza) Cut()     { fmt.Println("Cutting Pepperoni Pizza") }
func (p *PepperoniPizza) Box()     { fmt.Println("Boxing Pepperoni Pizza") }

func (p *VeggiePizza) Prepare() { fmt.Println("Preparing Veggie Pizza") }
func (p *VeggiePizza) Bake()    { fmt.Println("Baking Veggie Pizza") }
func (p *VeggiePizza) Cut()     { fmt.Println("Cutting Veggie Pizza") }
func (p *VeggiePizza) Box()     { fmt.Println("Boxing Veggie Pizza") }

type PizzaFactory interface {
	CreatePizza() Pizza
}

type CheesePizzaFactory struct{}
type PepperoniPizzaFactory struct{}
type VeggiePizzaFactory struct{}

func (f *CheesePizzaFactory) CreatePizza() Pizza { return &CheesePizza{} }
func (f *PepperoniPizzaFactory) CreatePizza() Pizza { return &PepperoniPizza{} }
func (f *VeggiePizzaFactory) CreatePizza() Pizza { return &VeggiePizza{} }

func main() {
	pizzaFactories := []PizzaFactory{
		&CheesePizzaFactory{},
		&PepperoniPizzaFactory{},
		&VeggiePizzaFactory{},
	}

	for _, factory := range pizzaFactories {
		pizza := factory.CreatePizza()
		pizza.Prepare()
		pizza.Bake()
		pizza.Cut()
		pizza.Box()
		fmt.Println("---")
	}
}
