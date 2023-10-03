package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол", "пятка"}
	m := getAnagrams(words)
	fmt.Println(m)
}

func getAnagrams(words []string) map[string][]string {
	anagramsMap := make(map[string][]string)
	tmpMap := make(map[string]bool)

	//Перебираем все слова
	for _, word := range words {
		tmpWord := toValidWord(word)
		
		//Проверям было ли добавлено слово
		_, ok := anagramsMap[tmpWord]
		if !ok && tmpMap[word] == false {
			anagramsMap[tmpWord] = []string{word}
		} else if tmpMap[word] == false {
			anagramsMap[tmpWord] = append(anagramsMap[tmpWord], word)
		}

		tmpMap[word] = true
	}

	//Удаляем ненужные слова
	for k, v := range anagramsMap {
		delete(anagramsMap, k)
		if len(v) > 1 {
			anagramsMap[v[0]] = v 
		}	
	}

	return anagramsMap
}

//Приводим строку к нижнему регистру и сортируем
func toValidWord(word string) string {
	var tmp strings.Builder
	for _, ch := range word {
		tmp.WriteString(string(unicode.ToLower(ch)))
	}

	s := strings.Split(tmp.String(), "")
    sort.Strings(s)

    return strings.Join(s, "")
}
