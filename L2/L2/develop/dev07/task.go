package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/


func or(channels ...<-chan any) <-chan any {
	var wg sync.WaitGroup
	ch := make(chan any)

	defer close(ch)

	//Перебираем каналы
	for _, channel := range channels { 
		//Инкрементируем счетчик WaitGroup для отслеживания запущенных горутин
		wg.Add(1)
		go func(c <-chan any) {
			defer wg.Done() //По завершении горутины уменьшаем счетчик WaitGroup
			for v := range c {
				ch <- v  //Записываем значения в выодной канал
			}
		}(channel)
	}

	go func() { 
		wg.Wait()   //Ждем, когда все горутины завершатся
		close(ch)   // По завершении всех горутин, закрываем канал 
	}()
	
	return ch
}

func main() {
	//var or func(channels ...<-chan interface{}) <-chan interface{}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
