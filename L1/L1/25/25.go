package main

import (
	"fmt"
	"time"
)

//OwnSleep - собственная функция sleep
func OwnSleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало")
	OwnSleep(5 * time.Second)
	fmt.Println("Прошло 5 секунд")
}
