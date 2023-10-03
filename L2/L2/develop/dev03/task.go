package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	/*Без дополнительных*/
	kFlag := flag.Int("k", 1, "Sort column")
	nFlag := flag.Bool("n", false, "Numeric sort")
	rFlag := flag.Bool("r", false, "Reverse")
	uFlag := flag.Bool("u", false, "Unique")

	flag.Parse()

	//Мапа для уникальных строк
	uniqueLine := make(map[string]bool) 
	lines := []string{}

	//Чтение из файлов
	for _, fileName := range flag.Args() {
		fmt.Println(flag.Args())
		 //Открываем файл
		inFile, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Can't open file")
			continue
		}
		fmt.Println(inFile.Name())
		defer inFile.Close()

		reader := bufio.NewReader(inFile)
		for {
			//Читаем построчно
			line, err := reader.ReadString('\n') 
			if err != nil && err != io.EOF {
				fmt.Println(err)
				break
			}

			//Очищаем строку от непечатных символов
			line = strings.TrimFunc(line, func(r rune) bool {
				return !unicode.IsGraphic(r) && unicode.IsSpace(r)
			})

			//Если нужны только уникальные строки
			if *uFlag && !uniqueLine[line] { 
				lines = append(lines, line)
				uniqueLine[line] = true
			} else if !*uFlag {
				lines = append(lines, line)
			}

			if err == io.EOF {
				break
			}
		}
	}

	sort.Slice(lines, func(i, j int) bool {
		if *nFlag && *kFlag > 1 { //Если есть n и k флаги
			firstNum, errFirst := strconv.Atoi(strings.Split(lines[i], " ")[*kFlag])
			secondNum, errSecond := strconv.Atoi(strings.Split(lines[j], " ")[*kFlag])
			if errFirst != nil && errSecond != nil {
				return lines[i] < lines[j]
			}
			return secondNum < firstNum
		} else if *nFlag { //Если есть только n флаг
			firstNum, errFirst := strconv.Atoi(strings.Split(lines[i], " ")[0])
			secondNum, errSecond := strconv.Atoi(strings.Split(lines[j], " ")[0])
			if errFirst != nil && errSecond != nil {
				return lines[i] < lines[j]
			}
			return secondNum < firstNum
		} else if *kFlag > 1 { //Если есть только k флаг
			firstSplit := strings.Split(lines[i], " ")
			secondSplit := strings.Split(lines[j], " ")
			if len(firstSplit) > *kFlag-1 && len(secondSplit) > *kFlag-1 {
				return firstSplit[*kFlag-1] < secondSplit[*kFlag-1]
			}
			return lines[i] < lines[j]
		}

		return lines[i] < lines[j]
	})

	outFile, err := os.OpenFile("out.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) //0666 => (wiki) File-system permissions - read & write
	defer outFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	//Пишем в файл в правильном  порядке
	if *rFlag {
		for i := len(lines) - 1; i >= 0; i-- {
			fmt.Fprintf(outFile, "%v\n", lines[i])
		}
	} else {
		for i := 0; i < len(lines); i++ {
			fmt.Fprintf(outFile, "%v\n", lines[i])
		}
	}
}
