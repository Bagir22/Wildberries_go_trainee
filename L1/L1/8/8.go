package main

import (
	"fmt"
)

func main() {
	var num int64 = 22 //Начальное число
	i := 4       //Индекс бита
	value := 1   //Новый бит

	result := BitToOne(num, i)
	fmt.Printf("Исходное число: %d - %b\n", num, num)
	fmt.Printf("Установка бита %d в %d: %d - %b\n", i, value, result, result)

	i = 3
	value = 0
	result = BitToZero(num, i)
	fmt.Printf("Исходное число: %d - %b\n", num, num)
	fmt.Printf("Установка %d бита в %d: %d - %b\n", i, value, result, result)
}

func BitToOne(num int64, i int) int64 {
	mask := int64(1 << (i - 1))
	return num | mask
}

func BitToZero(num int64, i int) int64 {
	mask := int64(^(1 << (i - 1)))
	return num & mask
}