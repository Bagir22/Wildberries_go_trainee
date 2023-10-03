package main

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Комманда позволяет инкапсулировать запрос на выполнение определенного действия в виде отдельного объект

	+
	1. Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	-
	1. Усложняет код программы из-за введения множества дополнительных классов.
*/ 

import "fmt"

type Command interface {
	Execute()
}

type LightOnCommand struct {
	Light *Light
}

func (c *LightOnCommand) Execute() {
	c.Light.TurnOn()
}

type LightOffCommand struct {
	Light *Light
}

func (c *LightOffCommand) Execute() {
	c.Light.TurnOff()
}

type Light struct {
	State string
}

func (l *Light) TurnOn() {
	l.State = "on"
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	l.State = "off"
	fmt.Println("Light is off")
}

type RemoteControl struct {
	Command Command
}

func (rc *RemoteControl) PressButton() {
	rc.Command.Execute()
}

func main() {
	light := &Light{}

	lightOnCommand := &LightOnCommand{Light: light}
	lightOffCommand := &LightOffCommand{Light: light}

	remoteControl := &RemoteControl{Command: lightOnCommand}
	remoteControl.PressButton()

	remoteControl.Command = lightOffCommand
	remoteControl.PressButton()
}