package main

import "fmt"

func main() {

	var number1 int
	var number2 int

	fmt.Print("Entre com o número 1: ")
	fmt.Scanln(&number1)

	fmt.Print("Entre com o número 2: ")
	fmt.Scanln(&number2)

	if number1 > number2 {
		fmt.Println("O maior número é o: ", number1)
	} else { // O Else tem que vir logo após o fechamento das chaves do if

		fmt.Println("O maior número é o: ", number2)

	}
}
