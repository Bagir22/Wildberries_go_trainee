package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var timeoutFlag int

func main() {
	/*
		go build go-telnet.go

		~~~~
		Run example:
		./go-telnet --timeout=10 localhost 8080
		~~~~
	*/

	flag.IntVar(&timeoutFlag, "timeout", 10, "server connection timeout")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("Invalid number of arguments")
	}

	//Получаем host и port
	host, port := flag.Arg(0), flag.Arg(1)
	address := fmt.Sprintf("%s:%s", host, port)
	//Устанавливаем соединение с портом
	conn, err := net.DialTimeout("tcp", address, time.Second*time.Duration(timeoutFlag))
	if err != nil {
		time.Sleep(time.Second * time.Duration(timeoutFlag))
		log.Fatal(err)
	}

	defer conn.Close()

	//Создание канала для прерывания по Ctrl+C
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT)

	//stdin => socket
	go func() {
		//Отправляем в stdin 
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		interrupt <- syscall.SIGINT //Прерываем
	}()

	//socket => stdout
	go func() {
		//Отправляем в stdout 
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		interrupt <- syscall.SIGINT ///Прерываем
	}()

	//Ожидание триггера на каналах и реакция на них
	<-interrupt //Ожидаем сигнала об ошибке
	fmt.Println("\nConnection closed")
}
