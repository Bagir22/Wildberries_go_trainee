package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Создаем пустые мапы
	hm1, hm2, intersection :=  make(map[int]bool), make(map[int]bool), make(map[int]bool)

	//Вставляем по 10 случайных чисел в обе мапы
	insertRandomToMap(hm1, 10)
	insertRandomToMap(hm2, 10)

	//Ищем пересечения
	findIntersection(hm1, hm2, intersection)

	fmt.Println("Пересечение множеств:")
	for k := range intersection {
		fmt.Print(k, " ")
	}
	fmt.Println()
}

func insertRandomToMap(hm map[int]bool, count int) {
	//Добавление случайного числа в мапу 
	for i := 0; i < count; i++ {
		hm[rand.Intn(10)] = true
	}
}

func findIntersection(hm1, hm2, intersection map[int]bool) {
	//Поиск пересечений
	//Если значения нет в мапе пересечения и есть в другой мапе, то
	//Добавляем в мапу пересечения
	for k := range hm1 {
		if !intersection[k] && hm2[k] {
			intersection[k] = true
		}
	}

	for k := range hm2 {
		if !intersection[k] && hm1[k] {
			intersection[k] = true
		}
	}
}
