package main

import (
	"fmt"
)

func main() {
	exsist := "главрыба — абырвалг"
	reversed := reverseString(exsist)
	fmt.Println("Исходная строка:", exsist)
	fmt.Println("Перевернутая строка:", reversed)
}

func reverseString(s string) string {
	//Конвертируем строку в руны
	runes := []rune(s)

	//Переворачиваем руны в массиве
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
