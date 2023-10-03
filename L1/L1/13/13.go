package main

import "fmt"

func main() {
	num1, num2 := 5, 10

	//Через присвоение
	fmt.Println("До: ", num1, num2)
	num1, num2 = num2, num1 
	fmt.Println("После: ", num1, num2)

	//Через xor
	fmt.Println("До: ", num1, num2)
	num1 = num1 ^ num2
	num2 = num2 ^ num1
	num1 = num1 ^ num2
	
	fmt.Println("После: ", num1, num2)	
}
