package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var input string
	fmt.Scanln(&input)

	result, err := RLEDecode(input)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}

func RLEDecode(str string) (string, error) {
	//Если пустая строка возращаем пустую строку
	if len(str) == 0 {
		return "", nil
	}

	//Если первый символ - число, то возразаем ошибку
	runes := []rune(str)
	if unicode.IsDigit(runes[0]) {
		return "", errors.New("Invalid string")
	}

	result := ""

	for i := 0; i < len(str); i++ {
		if unicode.IsLetter(runes[i]) || runes[i] == '\\' {
			if runes[i] == '\\' {
				i++
			}

			//Обработка повторений
			if i+1 < len(str) && unicode.IsDigit(runes[i+1]) {
				prev := runes[i]
				count := int(runes[i+1]) - '0'
				i++

				for i+1 < len(str) && unicode.IsDigit(runes[i+1]) {
					i++
					count = count*10 + int(runes[i+1]) - '0'
				}

				//Добавление повторений в result
				for j := 0; j < count; j++ {
					result += string(prev)
				}
			} else {
				//Добавление символов без повторений
				result += string(runes[i])
			}
		}
	}

	return result, nil
}
