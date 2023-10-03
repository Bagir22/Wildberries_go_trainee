package main

import (
	"fmt"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//мапа хранящая температуры по шагу
	grouped := make(map[int][]float64)

	step := 10

	for _, temp := range temps {
		key := Round(int(temp), step) //Получаем ключ
		grouped[key] = append(grouped[key], temp) //Добавляем температуру в список 
	}

	for k, v := range grouped {
		//Выводим ключ и список температур по нему
		fmt.Printf("%d: %v\n", k, v)
	}
}

func Round(x, unit int) int {
	return int(x / unit * unit)
}
