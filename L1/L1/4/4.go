package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Считываем количество горутин
	var workerCount int  
	fmt.Scanln(&workerCount)

	ch := make(chan int)

	for i := 0; i < workerCount; i++ {
		go ReadFromChannel(ch, i) //Запуск workerCount горутин читающих с канала 
	}

	go WriteToChannel(ch) //Запуск горутины отправлящей даные в канал

	//Создаем канал для сигналов от операционной системы
	sigs := make(chan os.Signal, 1)
	// Создаем канал для сигнала завершения программы
	done := make(chan bool, 1)

	//Настраиваем обработку сигналов SIGINT (Ctrl+C)
	signal.Notify(sigs, syscall.SIGINT)

	//// Запускаем горутину для обработки сигналов
	go func() {
		//Ожидаем получения сигнала из канала sigs
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		// Отправляем сигнал завершения в канал done
		done <- true
	}()
	
	// Ожидаем сигнала завершения из канала done
	<-done
}

//Бесконечная функция для записи случайных чисел в канал
func WriteToChannel(ch chan<- int) {
	for {
		val := rand.Intn(100)
		ch <- val
	}
}

//Бесконечная функция чтения из канала 
func ReadFromChannel(ch <-chan int, i int) {
	for {
		select {
		//Ожидаем значения из канала
		case v := <-ch:
			//Выводим значение и индекс горутины
			fmt.Printf("Горутина %d: %v\n", i, v)
		}
	}
}