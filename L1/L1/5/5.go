package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const seconds = 10 //время

func main() {
	//Создаем контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*seconds)
	defer cancel()

	//Создаем канал для передачи целых чисел
	ch := make(chan int)

	//Запуск горутины отправлящей данные в канал
	go WriteToChannel(ctx, ch)

	//Запуск горутины читаюзей данные с канала
	ReadFromChannel(ctx, ch)
}

func WriteToChannel(ctx context.Context, out chan<- int) {
	defer close(out)

	for {
		select {
		//При завершении контекста выходим из функции
		case <-ctx.Done():
			return
		//Отправляем случайное число в канал с задержкой в 0.5 секунды
		case out <- rand.Intn(100):
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func ReadFromChannel(ctx context.Context, ch chan int) {
	for {
		select {
		//При завершении контекста выходим из функции
		case <-ctx.Done():
			fmt.Println("Завершение чтения из канала")
			return
		//Читаем из канала
		case val := <-ch:
			fmt.Printf("Прочитано: %d\n", val)
		}
	}
}
