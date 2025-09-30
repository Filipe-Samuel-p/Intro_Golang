package main

import "fmt"

func sum(number1, number2 int) int {
	return number1 + number2
}

func diference(number1, number2 int) int {
	if number1 > number2 {
		return number1 - number2
	} else {
		return number2 - number1
	}
}

func product(number1 int, number2 int) int {
	return number1 * number2
}

func division(number1 int, number2 int) (int, error) {
	if number1 == 0 || number2 == 0 {
		return 0, fmt.Errorf("Divisão por zero")
	} else {
		return number1 / number2, nil
	}

}

func main() {

	var number1 int
	var number2 int

	fmt.Print("Forneca um número inteiro: ")
	fmt.Scan(&number1)

	fmt.Print("Forneca outro número inteiro: ")
	fmt.Scan(&number2)

	fmt.Println("1- Somar os dois números")
	fmt.Println("2- Diferença entre os dois números")
	fmt.Println("3- Produto entre os dois números")
	fmt.Println("4- Divisão entre os dois números")

	var option int

	fmt.Print("Escolha uma das opções: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		result := sum(number1, number2)
		fmt.Printf("O resultado da soma entre %d e %d é: %d", number1, number2, result)
	case 2:
		result := diference(number1, number2)
		fmt.Printf("O resultado da difença entre %d e %d é: %d", number1, number2, result)
	case 3:
		result := product(number1, number2)
		fmt.Printf("O resultado do produto entre %d e %d é: %d", number1, number2, result)
	case 4:
		result, err := division(number1, number2)
		if err != nil {
			fmt.Print("Erro: Divisão por zero")
			return
		}
		fmt.Printf("O resultado da dovosão entre %d e %d é: %d", number1, number2, result)

	default:
		fmt.Print("Opção inexistente")
	}

}
