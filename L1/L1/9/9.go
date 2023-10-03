package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем каналы для ввода и x2 чисел
	inputCh := make(chan int)
	outputCh := make(chan int)

	arr := []int{1, 2, 3, 4, 5}
	//Пишем числа в 1ый канал
	go inputNums(inputCh, arr)

	var wg sync.WaitGroup
	wg.Add(1)

	go x2Nums(inputCh, outputCh, &wg)


	// Горутина для чтения результата из 2ой канала и вывода в stdout
	go printNums(outputCh)

	wg.Wait()
	close(outputCh) // Закрываем 2ой канал после завершения всех горутин
}

func inputNums(ch chan<- int, arr []int) {
	for _, num := range arr {
		ch <- num
	}
	close(ch)
}

func x2Nums(inputCh <-chan int, outputCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range inputCh {
		outputCh <- num * 2
	}
}

func printNums(ch <-chan int) {
	for result := range ch {
		fmt.Println("Num x2:", result)
	}
}

