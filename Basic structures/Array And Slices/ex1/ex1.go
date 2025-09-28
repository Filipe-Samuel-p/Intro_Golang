package main

import "fmt"

func main() {

	var array [6]int

	for i := 0; i < len(array); i++ {
		fmt.Print("Escreva um número: ")
		fmt.Scan(&array[i])
	}

	fmt.Print(array)
}
