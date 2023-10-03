package main

import (
	"fmt"
	"strings"
)

func checkString(s string) bool {
	//Создаем мапу которая проверят был ли такой же символ в строке
	seen := make(map[rune]bool)
	for _, char := range s {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range strs {
		//Проверям строку в нижнем регистре
		result := checkString(strings.ToLower(str))
		fmt.Printf("Строка: %s, Результат: %v\n", str, result)
	}
}
