package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// Создание массива arr с пятью элементами
	arr := [5]int{2, 4, 6, 8, 10}

	//Создание WaitGroup для синхронизации горутин
	wg := sync.WaitGroup{}

	//Создание переменной sum для суммы квадратов
	var sum int 

	for _, value := range arr {
		//Добавление новой горутины в счетчик wg 
		wg.Add(1)
		go func(value int) {
			sum += int(math.Pow(float64(value), 2)) //Добавление квадрата числа к сумме
			wg.Done() //Уменьшение счетика горутин после завершения 1ой горутины 
		}(value)
	}

	wg.Wait() //Ожидание завершения всех горутин

	fmt.Printf("Sum of squares is: %d\n", sum) //Печать суммы квадратов
}
