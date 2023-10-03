package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	AFlag = flag.Int("A", 0, "печатать +N строк после совпадения")
	BFlag = flag.Int("B", 0, "печатать +N строк до совпадения")
	CFlag = flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	cFlag = flag.Bool("c", false, "количество строк")
	iFlag = flag.Bool("i", false, "игнорировать регистр")
	vFlag = flag.Bool("v", false, "вместо совпадения, исключать")
	FFlag = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	nFlag = flag.Bool("n", false, "печатать номер строки")
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("No file name")
	}

	pattern := flag.Arg(0)

	var text []string
	inFile, err := os.Open(os.Args[len(os.Args)-1]) //Открываем файл
	if err != nil {
		fmt.Println("Can't open file")
	}

	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	for {
		line, err := reader.ReadString('.') //Читаем по предложениям
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		text = append(text, line)

		if err == io.EOF {
			break
		}
	}

	flagAfter, flagBefore := 0, 0
	if *CFlag != 0 {
		flagAfter = *CFlag
		flagBefore = *CFlag
	} else {
		if *AFlag != 0 {
			flagAfter = *AFlag
		}

		if *BFlag != 0 {
			flagBefore = *BFlag
		}
	}

	count := 0
	printedLines := make(map[int]bool)
	for i, sentence := range text {
		if *iFlag {
			pattern = strings.ToLower(pattern)
			sentence = strings.ToLower(sentence)
		}

		var found bool
		// Поиск совпадений
		if *FFlag && sentence == pattern { //Обработка флага -F
			found = true
		} else {
			found = strings.Contains(sentence, pattern)
		}

		//Обработка флага -v
		if *vFlag {
			found = !found
		}

		//Вывод совпадений
		if found {
			if !*cFlag {
				// Обработка флага -B
				for j := i - flagBefore; j < i; j++ {
					if j >= 0 {
						output(j+1, text[j], printedLines)
					}
				}

				output(i+1, sentence, printedLines)
				// Обработка флага -A
				for j := i + 1; j <= i+flagAfter; j++ {
					if j < len(text) {
						output(i+1, text[j], printedLines)
					}
				}
			} else {
				count++
			}
		}
	}

	if *cFlag {
		fmt.Printf("%d\n", count)
	}
}

func output(lineNumber int, str string, pl map[int]bool) {
	if *nFlag && pl[lineNumber] == false {
		fmt.Printf("%d: %s\n", lineNumber, str)
		pl[lineNumber] = true
	} else if pl[lineNumber] == false {
		fmt.Printf("%s\n", str)
		pl[lineNumber] = true
	}
}
