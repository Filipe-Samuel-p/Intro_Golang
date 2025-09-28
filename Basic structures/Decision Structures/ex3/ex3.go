package main

import "fmt"

func main() {

	var number int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Print("O número ", number, " é par")
	} else {
		fmt.Print("O número ", number, " é ímpar")
	}

}
