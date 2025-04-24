package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/leolucena22/todo-list-API/src/model"
)

type BibliotecaRepository struct {
	connection *sql.DB
}

func NewBibliotecaRepository(connection *sql.DB) BibliotecaRepository {
	return BibliotecaRepository{
		connection: connection,
	}
}

func (r *BibliotecaRepository) Create(book *model.Book) (int, error) {
	var id int
	err := r.connection.QueryRow(`
		INSERT INTO books (title, isbn, Year, AuthorID, createdAt) VALUES ($1, $2, $3, $4, $5) RETURNING id
`, book.Title, book.ISBN, book.Year, book.AuthorID, time.Now()).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("erro ao criar livro: %w", err)
	}

	return id, nil
}
