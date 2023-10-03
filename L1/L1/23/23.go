package main

import "fmt"

func main() {
	slice := []int{ 1, 2, 3, 4, 5 }
	removeIndex := 2

	fmt.Println("Исходный слайс:", slice)

	// Проверяем, что индекс находится в допустимом диапазоне
	if removeIndex >= 0 && removeIndex < len(slice) {
		slice = append(slice[:removeIndex], slice[removeIndex+1:]...)
		fmt.Println("Срез после удаления элемента:", slice)
	} else {
		fmt.Println("Индекс выходит за пределы слайса")
	}
}
