package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
)

type BookRepository struct {
	Database *sql.DB
}

func (r BookRepository) GetBooksFromDatabase(sl *[]models.Book, rw http.ResponseWriter) error {
	var book models.Book
	rows, err := r.Database.Query("SELECT * FROM books")
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return err
		} else {
			*sl = append(*sl, book)
		}

	}
	defer rows.Close()
	return nil
}

func (r BookRepository) GetBookFromDatabase(rw http.ResponseWriter, id string) (models.Book, error) {
	var book models.Book
	query := `SELECT * FROM books WHERE id = $1`
	row := r.Database.QueryRow(query, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		fmt.Println("error")
		return models.Book{}, err
	}

	return book, nil
}

func (r BookRepository) WriteBookToTheDatabase(rw http.ResponseWriter, book models.Book) error {
	insertString := `INSERT INTO books(title, author, year) VALUES($1, $2, $3)`
	_, err := r.Database.Exec(insertString, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil
}

func (r BookRepository) UpdateBookFromDatabase(rw http.ResponseWriter, book models.Book, id string) error {
	if r.checkDatabaseForBookIdExisting(id) {
		fmt.Println("oooo")
		insertString := `UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4`
		r.Database.Exec(insertString, book.Title, book.Author, book.Year, id)
		return nil
	} else {
		return errors.New("id does not exist")
	}

}

func (r BookRepository) DeleteBookFromDatabase(rw http.ResponseWriter, id string) error {
	if r.checkDatabaseForBookIdExisting(id) {
		fmt.Println("oooo")
		insertString := `DELETE FROM books WHERE id = $1`
		r.Database.Exec(insertString, id)
		return nil
	} else {
		return errors.New("id does not exist")
	}

}

func (r BookRepository) checkDatabaseForBookIdExisting(id string) bool {
	sqlStmt := `SELECT * FROM books WHERE id = $1`
	err := r.Database.QueryRow(sqlStmt, id).Scan()
	fmt.Println(err)
	return err != sql.ErrNoRows
}
