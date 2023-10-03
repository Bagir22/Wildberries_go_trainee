package main

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

	+
	1. Позволяет создавать продукты пошагово.
	2. Позволяет использовать один и тот же код для создания различных продуктов.

	-
	1. Усложняет код программы из-за введения дополнительных классов.
	2. В случае, если у разных продуктов есть схожие части конструирования, может возникнуть дублирование кода в различных строителях.
*/

import "fmt"

type Sandwich struct {
	Bread     string
	Meat      string
	Cheese    string
	Veggies   []string
	Sauce string
}

func (s *Sandwich) String() string {
	veggies := ""
	if len(s.Veggies) > 0 {
		veggies = " with " + s.joinItems(s.Veggies)
	}
	return fmt.Sprintf("Sandwich with %s bread,%s meat, %s cheese%s, and %s", s.Bread, s.Meat, s.Cheese, veggies, s.Sauce)
}

func (s *Sandwich) joinItems(items []string) string {
	return fmt.Sprintf("%s", items)
}

type SandwichBuilder interface {
	ChooseBread()
	AddMeat()
	AddCheese()
	AddVeggies()
	AddSauce()
	GetSandwich() *Sandwich
}

type VeggieSandwichBuilder struct {
	sandwich *Sandwich
}

func NewVeggieSandwichBuilder() *VeggieSandwichBuilder {
	return &VeggieSandwichBuilder{sandwich: &Sandwich{}}
}

func (vb *VeggieSandwichBuilder) ChooseBread() {
	vb.sandwich.Bread = "whole wheat"
}

func (vb *VeggieSandwichBuilder) AddMeat() {
}

func (vb *VeggieSandwichBuilder) AddCheese() {
	vb.sandwich.Cheese = "cheddar"
}

func (vb *VeggieSandwichBuilder) AddVeggies() {
	vb.sandwich.Veggies = []string{"onion", "tomato"}
}

func (vb *VeggieSandwichBuilder) AddSauce() {
	vb.sandwich.Sauce = "garlic Sauce"
}

func (vb *VeggieSandwichBuilder) GetSandwich() *Sandwich {
	return vb.sandwich
}

type SandwichDirector struct {
	builder SandwichBuilder
}

func (sa *SandwichDirector) CreateSandwich() {
	sa.builder.ChooseBread()
	sa.builder.AddMeat()
	sa.builder.AddCheese()
	sa.builder.AddVeggies()
	sa.builder.AddSauce()
}

func main() {
	veggieBuilder := NewVeggieSandwichBuilder()
	director := &SandwichDirector{builder: veggieBuilder}

	director.CreateSandwich()
	sandwich := veggieBuilder.GetSandwich()

	fmt.Println(sandwich)
}
