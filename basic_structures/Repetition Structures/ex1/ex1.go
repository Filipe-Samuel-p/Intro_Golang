package main

import "fmt"

func main() {

	var total int

	fmt.Print("Quantos múltiplos de 3 quer ver: ")
	fmt.Scan(&total)

	for i := 1; i <= total; i++ {
		fmt.Println(i * 3)
	}
}
