package main

import "fmt"

func main() {
	//Список строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	//Множество
	set := make(map[string]bool)

	//Добавление уникальных слов в множество
	for _, v := range arr {
		set[v] = true
	}

	fmt.Println("Множество:", set)
}
