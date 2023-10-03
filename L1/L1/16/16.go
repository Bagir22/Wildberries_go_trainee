package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Исходный массив:", arr)

	quicksort(arr)
	fmt.Println("Отсортированный массив:", arr)
}

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left := 0
	right := len(arr) - 1

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quicksort(arr[:left])
	quicksort(arr[left+1:])
}
