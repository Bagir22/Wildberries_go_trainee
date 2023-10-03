package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создание мьютекса для управления доступом к данным
	var rwMutex sync.RWMutex

	data := make(map[int]int)

	//Создание WaitGroup для синхронизации горутин
	var wg sync.WaitGroup
	numWorkers := 5 

	wg.Add(numWorkers) //Количество горутин записываюших в мапу
	for i := 1; i <= numWorkers; i++ {
		//Запуск конкурентной записи в map
		go func(i int) {
			//Уменьшение горутин при завершении
			defer wg.Done()

			value := i * 10

			//Блокировка перед записью в мапу
			rwMutex.Lock()
			data[i] = value
			fmt.Printf("Worker %d записал %d в map\n", i, value)
			//Разблокировка 
			rwMutex.Unlock()
			
		}(i)
	}

	wg.Wait() //Ожидание завершения всех горутин

	//Выводим содержимое map после всех записей
	fmt.Println("Содержимое map:")
	rwMutex.RLock()
	for key, value := range data {
		fmt.Printf("%d: %d\n", key, value)
	}
	rwMutex.RUnlock()
}
