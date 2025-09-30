package main

import "fmt"

func bubbleSort(array []int) []int {

	var aux int
	var count int = 0

	for count < len(array)-1 {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				aux = array[j]
				array[j] = array[j+1]
				array[j+1] = aux
			}
		}
		count++
	}

	return array

}

func main() {

	fmt.Print("Entre com o tamanho do vetor: ")
	var size int
	fmt.Scan(&size)
	array := make([]int, size) // slice

	for i := 0; i < len(array); i++ {
		fmt.Print("Digite um número inteiro: ")
		fmt.Scan(&array[i])
	}

	fmt.Println("vetor original: ", array)

	sortArray := bubbleSort(array)

	fmt.Print("Vetor organizado: ", sortArray)

}
