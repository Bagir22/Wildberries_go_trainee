package main

import (
	"fmt"
)

func main() {
	//Создание слайса с разными типами
	values := []interface{}{10, "String", true, make(chan int)}

	for _, v := range values {
		//Вызов функции getType по которой мы получаем тип значения
		getType(v)
	}
}

func getType(value interface{}) {
	//Использование switch для определения типа значения
	switch value.(type) {
	case int:
		fmt.Printf("%v - int\n", value)
	case string:
		fmt.Printf("%v - string\n", value)
	case bool:
		fmt.Printf("%v - bool\n", value)
	case chan int:
		fmt.Printf("%v - chan int\n", value)
	}
}
