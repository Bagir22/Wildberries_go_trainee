package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fFlag string
	dFlag string
	sFlag bool
)

func init() {
	flag.StringVar(&fFlag, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&dFlag, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&sFlag, "s", false, "только строки с разделителем")
}

func main() {
	flag.Parse()

	//go run task.go -d=, -f=1 -s in.txt
	//cut -d, -f1 in.txt

	//Разделяем на колонки
	columns := []int{}
	if len(fFlag) != 0 {
		var flags []string
		if strings.Contains(fFlag, ",") {
			flags = strings.Split(fFlag, ",")
		} else if strings.Contains(fFlag, "-") {
			flags = strings.Split(fFlag, "-")
		} else {
			addToColumns(fFlag, columns)
		}

		for _, v := range flags {
			addToColumns(v, columns)
		}
	}

	inFile, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("Can't open file")
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text() //Читаем по предложениям
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		//Если sFlag и в строке нет разделителя, то пропускаем
		if sFlag && !strings.Contains(line, dFlag) {
			continue
		}

		//Разбиваем строку по разделителю
		fields := strings.Split(line, dFlag)
		
		if len(columns) != 0 {
			resLine := ""
			for _, v := range columns {
				if v-1 >= 0 && v-1 < len(fields) {
					resLine += fields[v-1]
				} else {
					fmt.Println(line)
				}				
			}
			fmt.Println(resLine)
		} else {
			for _, v := range fields {
				fmt.Print(v)
			}
			fmt.Println()
		}
	}
}

func addToColumns(flag string, columns []int) {
	column, err := strconv.Atoi(flag)
	if err != nil {
		fmt.Println("Can't parse colum")
	} else {
		columns = append(columns, column)
	}
}
