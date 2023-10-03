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

	for _, value := range arr {
		//Добавление новой горутины в счетчик wg 
		wg.Add(1)

		go func(value int) { //Запуск анонимной функции расчитывающей квадрат
			fmt.Printf("Square of %d is %d\n", value, int(math.Pow(float64(value), 2)))
			wg.Done() //Уменьшение счетика горутин после завершения 1ой горутины 
		}(value)
	}

	wg.Wait() //Ожидание завершения всех горутин
}
