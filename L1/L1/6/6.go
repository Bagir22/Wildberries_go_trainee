package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//Использование канала для сигнала завершения
	doneChannel := make(chan bool)
	go func() {
		defer fmt.Println("Завершение по каналу")
		time.Sleep(2 * time.Second)
		doneChannel <- true
	}()

	//Использование sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Завершение по WaitGroup")
		defer wg.Done()
		time.Sleep(2 * time.Second)
	}()

	//Использование контекста
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		defer fmt.Println("Завершение по контексту")
		<-ctx.Done()	
	}()

	//Использование runtime.Goexit()
	go func() {
		defer fmt.Println("Заверешение по runtime.Goexit()")
		runtime.Goexit()
	}()

	//Ожидание завершения горутин с помощью канала
	<-doneChannel

	//Ожидание завершения горутин с помощью sync.WaitGroup
	wg.Wait()

	//Остановка горутины с контекстом
	cancel()

	time.Sleep(1 * time.Second)
}
