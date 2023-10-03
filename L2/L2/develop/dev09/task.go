package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	oFlag string
}

var (
	oFlag string
)

func main() {
	//go run task.go -o output google.com  			=> output.html
	//go run task.go google.com						=> index.html

	flag.StringVar(&oFlag, "o", "index.html", "change output file")
	flag.Parse()

	url := flag.Arg(0) //Получаем url сайта 
	if url == "" {
		log.Fatal("Input site url")
	}

	Wget(url, oFlag)
}

func Wget(url, filename string) {
	//Если нет префикса, то добавляем
	if !strings.HasPrefix(filename, "http://") || !strings.HasPrefix(filename, "https://"){
		url = fmt.Sprint("http://", url)
	}

	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		log.Fatalf("Response status code: %v\nResponse error: %v\n", response.StatusCode, err)
	}
	defer response.Body.Close()

	//Если в выходном файле нет суфикса .html
	if !strings.HasSuffix(filename, ".html") {
		filename = fmt.Sprint(filename, ".html")
	}

	//Удаляем пробелы
	filename = strings.TrimSpace(filename)

	fmt.Printf("Output file: %v\nUrl: %v\n", filename, url)
	
	outFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) //0666 => (wiki) File-system permissions - read & write
	defer outFile.Close()
	if err != nil {
		log.Fatalf("Out file error: %v\n", err)
	}

	//Копируем ответ в выходной файл
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatalf("Can't download %v page: %v\n", url, err)
	}
}
