package main

import (
	"errors"
	"fmt"
	"slices"
)

type Book struct {
	Id     string
	Title  string
	Author string
	Year   int
}

func createBook() Book {

	var id, title, author string
	var year int
	fmt.Print("Digite o Id do livro: ")
	fmt.Scanln(&id)
	fmt.Print("Digite o título do livro: ")
	fmt.Scanln(&title)
	fmt.Print("Digite o autor do livro: ")
	fmt.Scanln(&author)
	fmt.Print("Digite o ano de publicação do livro: ")
	fmt.Scan(&year)

	newBook := Book{
		Id:     id,
		Title:  title,
		Author: author,
		Year:   year,
	}

	return newBook
}

func updateBook(book *Book, newId string, newTitle string, newAuthor string, newYear int) {
	book.Id = newId
	book.Title = newTitle
	book.Author = newAuthor
	book.Year = newYear
}

type Repository struct {
	Db []*Book
}

func (r *Repository) addBook(book *Book) {
	r.Db = append(r.Db, book)
}

func (r Repository) getAllBooks() {
	fmt.Println("Todos os livros: ")
	for i := 0; i < len(r.Db); i++ {
		fmt.Println(*r.Db[i])
	}
}

func (r Repository) findById(id string) (*Book, error) {
	for i := 0; i < len(r.Db); i++ {
		if r.Db[i].Id == id {
			return r.Db[i], nil
		}
	}
	return nil, fmt.Errorf("Não encontrado")
}

func (r *Repository) removeBook(id string) error {
	_, err := r.findById(id)
	if err != nil {
		return err
	}

	var index int = -1
	for i := 0; i < len(r.Db); i++ {
		if r.Db[i].Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("livro não encontrado")
	}

	r.Db = slices.Delete(r.Db, index, index+1)
	return nil
}

func main() {

	var repo Repository

	for {
		fmt.Println("\n**** Bem Vindos ao CRUD de livros: ****\n")
		fmt.Println("1- Criar Livro")
		fmt.Println("2- Listar Livros")
		fmt.Println("3- Buscar por ID")
		fmt.Println("4-	Atualizar Livro")
		fmt.Println("5- Remover Livro")
		fmt.Println("6- Sair\n")

		fmt.Print("Selecione uma das opções acima para continuar: ")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			book := createBook()
			repo.addBook(&book)
		case 2:
			repo.getAllBooks()
		case 3:
			fmt.Print("Digite o Id do livro: ")
			var id string
			fmt.Scanln(&id)
			book, err := repo.findById(id)
			if err == nil {
				fmt.Print("Livro encontrado: ", *book)
			}

		case 4:
			fmt.Print("Digite o Id do livro a se atualizar: ")
			var id string
			fmt.Scanln(&id)
			book, err := repo.findById(id)
			if err == nil {
				var newid, newtitle, newauthor string
				var newyear int
				fmt.Print("Digite o Id do livro: ")
				fmt.Scanln(&newid)
				fmt.Print("Digite o título do livro: ")
				fmt.Print()
				fmt.Scanln(&newtitle)
				fmt.Print("Digite o autor do livro: ")
				fmt.Scanln(&newauthor)
				fmt.Print("Digite o ano de publicação do livro: ")
				fmt.Scanln(&newyear)
				updateBook(book, newid, newtitle, newauthor, newyear)
			}
		case 5:
			fmt.Print("Digite o Id do livro a ser removido: ")
			var id string
			fmt.Scanln(&id)
			_, err := repo.findById(id)
			if err == nil {
				repo.removeBook(id)
			}
		case 6:
			return
		default:
			fmt.Println("Opção inexistente\n")
		}

	}

}
