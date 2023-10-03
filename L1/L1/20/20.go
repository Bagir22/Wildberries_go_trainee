package main

import (
	"fmt"
	"strings"
)

func main() {
	exsist := "snow dog sun — sun dog snow"
	reversed := reverseWords(exsist)
	fmt.Println("Исходная строка:", exsist)
	fmt.Println("Перевернутые слова:", reversed)
}

func reverseWords(s string) string {
	//Получаем отдельные слова
	words := strings.Fields(s)
	//Создаем слайс длинной = количеству слов 
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	return strings.Join(reversedWords, " ")
}
