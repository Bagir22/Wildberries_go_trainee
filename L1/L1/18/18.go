package main

import (
	"fmt"
	"sync"
)

//структура Counter с счетчиком и мьютексом
type Counter struct {
	value int
	mu    sync.Mutex
}

//Инкремент счетчика
func (c *Counter) Increment() {
	c.mu.Lock() //Блокируем

	defer c.mu.Unlock() //Разблокируем
	c.value++ //Увеличиваем счетчик
}

//Получаем значение счетчика
func (c *Counter) GetValue() int {
	c.mu.Lock() //Блокируем
	defer c.mu.Unlock() //Разблокируем
	return c.value //Возразаем значение счетчика
}

func main() {
	counter := &Counter{} //Создаем экземпляр Counter

	var wg sync.WaitGroup //Создаем wg для синхронизации горутин
	countWorkers := 10 //5 горутин

	for i := 0; i < countWorkers; i++ {
		wg.Add(1) // Увеличиваем счетчик ожидаемых горутин
		go func() {
			defer wg.Done() //Отложенно уменьшаем счетчик ожидаемых горутин
			counter.Increment() //Увеливаем счетчик
		}()
	}

	wg.Wait() //Блокируем до завершения всех горутин

	fmt.Printf("Counter value: %d\n", counter.GetValue()) //Выводим итоговое количество счетчика
}
