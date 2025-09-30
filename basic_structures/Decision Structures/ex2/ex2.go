package main

import (
	"fmt"
	"math"
)

func main() {

	var number float32

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&number)

	sqrt := math.Sqrt(float64(number))

	if math.IsNaN(sqrt) {
		fmt.Println("O número digitado é negativo")
		return
	}

	fmt.Print("Raiz quadrada do número digitado: ", sqrt)
}
