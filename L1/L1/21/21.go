package main

import (
	"fmt"
)

//Интерфейс для работы с градусами Цельсия
type CelsiusTemperature interface {
	GetCelsius() float64
}

//Структура для градусов Цельсия
type Celsius struct {
	temperature float64
}

//Получаем температуру в градусах Цельсия
func (c *Celsius) GetCelsius() float64 {
	return c.temperature
}

//Адаптер для конвертации из градусов Фаренгейта в градусы Цельсия
type FahrenheitAdapter struct {
	fahrenheit float64
}

func (fa *FahrenheitAdapter) GetCelsius() float64 {
	celsius := FahrenheitToCelsius(fa.fahrenheit)
	return celsius
}


func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5.0 / 9.0
}


func main() {
	celsiusTemp := Celsius{ temperature: 25.0 }
	fmt.Printf("Температура в градусах Цельсия: %.f\n", celsiusTemp.GetCelsius())

	fahrenheitTemp := FahrenheitAdapter{ fahrenheit: 77.0 }
	fmt.Printf("Температура в градусах Фаренгейта, преобразованная в граудусф Цельсия: %.f\n", fahrenheitTemp.GetCelsius())
}


