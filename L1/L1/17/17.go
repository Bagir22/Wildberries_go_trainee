package main

import (
	"fmt"
)

func main() {
	//Создадим отсортированный массив
	arr := []int{1, 2, 2, 5, 6, 7, 8, 9, 10, 12, 13, 15, 18, 20}

	//Будем искать индекс числа 12
	target := 12
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Число %d найдено, его индекс - %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найдено\n", target)
	}

	//Будем искать индекс числа 3
	target = 3
	index = binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Число %d найдено, его индекс - %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найдено\n", target)
	}
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
