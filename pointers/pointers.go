package main

import "fmt"

// ============================================
// EXERCÍCIO 1: Básico - Trocar valores
// ============================================
// Complete a função para trocar os valores de duas variáveis usando ponteiros
func trocar(a, b *int) {

	change := *a
	*a = *b
	*b = change

}

// ============================================
// EXERCÍCIO 2: Dobrar um valor
// ============================================
// Complete a função para dobrar o valor de um inteiro usando ponteiros
func dobrar(n *int) {
	*n = *n + *n

}

// ============================================
// EXERCÍCIO 3: Modificar struct
// ============================================
type Pessoa struct {
	Nome  string
	Idade int
}

// Complete a função para incrementar a idade de uma pessoa
func aniversario(p *Pessoa) {
	p.Idade++
	// OU
	(*p).Idade++
}

// ============================================
// EXERCÍCIO 4: Verificar nil
// ============================================
// Complete a função que retorna o valor se o ponteiro não for nil,
// caso contrário retorna um valor padrão
func valorOuPadrao(ptr *int, padrao int) int {
	if ptr != nil {
		return *ptr
	}
	return padrao

}

// ============================================
// EXERCÍCIO 5: Criar e retornar ponteiro
// ============================================
// Complete a função que cria uma nova Pessoa e retorna seu ponteiro
func novaPessoa(nome string, idade int) *Pessoa {
	newPerson := &Pessoa{Nome: nome, Idade: idade}
	return newPerson
}

// ============================================
// EXERCÍCIO 6: Array com ponteiros
// ============================================
// Complete a função que incrementa todos os valores de um slice usando ponteiros
func incrementarTodos(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i]++
	}
}

// ============================================
// EXERCÍCIO 7: Ponteiro para ponteiro
// ============================================
// Complete a função que modifica um valor através de um ponteiro para ponteiro
func modificarPonteiroParaPonteiro(pp **int, novoValor int) {
	**pp = novoValor
}

// ============================================
// EXERCÍCIO 8: Lista encadeada simples
// ============================================
type No struct {
	Valor   int
	Proximo *No
}

// Complete a função que adiciona um novo nó no final da lista
func adicionarNoFinal(cabeca *No, valor int) {

	if cabeca.Proximo == nil {
		newNode := new(No)
		cabeca.Proximo = newNode
		newNode.Valor = valor
		newNode.Proximo = nil
		return
	}
	adicionarNoFinal(cabeca.Proximo, valor)

}

// ============================================
// TESTES - Execute para verificar suas soluções
// ============================================
func main() {
	fmt.Println("=== Testando Exercícios de Ponteiros ===\n")

	// Teste 1
	fmt.Println("Exercício 1: Trocar valores")
	x, y := 5, 10
	fmt.Printf("Antes: x=%d, y=%d\n", x, y)
	trocar(&x, &y)
	fmt.Printf("Depois: x=%d, y=%d\n", x, y)
	fmt.Printf("Esperado: x=10, y=5\n\n")

	// Teste 2
	fmt.Println("Exercício 2: Dobrar valor")
	num := 7
	fmt.Printf("Antes: %d\n", num)
	dobrar(&num)
	fmt.Printf("Depois: %d\n", num)
	fmt.Printf("Esperado: 14\n\n")

	// Teste 3
	fmt.Println("Exercício 3: Aniversário")
	pessoa := Pessoa{Nome: "Ana", Idade: 25}
	fmt.Printf("Antes: %+v\n", pessoa)
	aniversario(&pessoa)
	fmt.Printf("Depois: %+v\n", pessoa)
	fmt.Printf("Esperado: Idade=26\n\n")

	// Teste 4
	fmt.Println("Exercício 4: Valor ou padrão")
	var ptr *int = nil
	resultado := valorOuPadrao(ptr, 100)
	fmt.Printf("Com ponteiro nil: %d (esperado: 100)\n", resultado)
	valor := 42
	resultado = valorOuPadrao(&valor, 100)
	fmt.Printf("Com ponteiro válido: %d (esperado: 42)\n\n", resultado)

	// Teste 5
	fmt.Println("Exercício 5: Nova pessoa")
	novaPessoa1 := novaPessoa("Carlos", 30)
	fmt.Printf("Nova pessoa: %+v\n", novaPessoa1)
	fmt.Printf("Esperado: &{Nome:Carlos Idade:30}\n\n")

	// Teste 6
	fmt.Println("Exercício 6: Incrementar todos")
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Printf("Antes: %v\n", numeros)
	incrementarTodos(&numeros)
	fmt.Printf("Depois: %v\n", numeros)
	fmt.Printf("Esperado: [2 3 4 5 6]\n\n")

	// Teste 7
	fmt.Println("Exercício 7: Ponteiro para ponteiro")
	val := 50
	pval := &val
	fmt.Printf("Antes: %d\n", val)
	modificarPonteiroParaPonteiro(&pval, 99)
	fmt.Printf("Depois: %d\n", val)
	fmt.Printf("Esperado: 99\n\n")

	// Teste 8
	fmt.Println("Exercício 8: Lista encadeada")
	lista := &No{Valor: 1, Proximo: nil}
	adicionarNoFinal(lista, 2)
	adicionarNoFinal(lista, 3)
	fmt.Print("Lista: ")
	for n := lista; n != nil; n = n.Proximo {
		fmt.Printf("%d ", n.Valor)
	}
	fmt.Println("\nEsperado: 1 2 3")
}
