package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"io"

	"github.com/SakuraBurst/books.git/main/helpers"
	"github.com/SakuraBurst/books.git/main/models"
)

type BookRepository struct {
	Database *sql.DB
}

func (r BookRepository) GetBooksFromDatabase(sl *[]models.Book) error {
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

func (r BookRepository) GetBookFromDatabase(id string) (models.Book, error) {
	var book models.Book
	query := `SELECT * FROM books WHERE id = $1`
	row := r.Database.QueryRow(query, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r BookRepository) WriteToTheDatabase(newInstanse models.InstanseMaker, body io.ReadCloser) error {
	book := helpers.MakeNewInstanse(newInstanse, body)
	if book.IsValid() {
		fields := models.GetFields(newInstanse)
		insertString := fmt.Sprintf(`INSERT INTO %v(%v) VALUES(%v)`, getInstanseTable(newInstanse), fields.BdFields, fields.BdValues)
		_, err := r.Database.Exec(insertString)
		if err != nil {
			return err
		}
		return nil

	} else {
		fmt.Println("error")
		return errors.New("some data is unprocessable")
	}

}

func (r BookRepository) WriteBookToTheDatabase(book models.Book) error {
	insertString := `INSERT INTO books(title, author, year) VALUES($1, $2, $3)`
	_, err := r.Database.Exec(insertString, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil
}

func (r BookRepository) UpdateBookFromDatabase(book models.Book, id string) error {
	if r.checkDatabaseForBookIdExisting(id) {
		insertString := `UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4`
		r.Database.Exec(insertString, book.Title, book.Author, book.Year, id)
		return nil
	} else {
		return errors.New("id does not exist")
	}

}

func (r BookRepository) DeleteBookFromDatabase(id string) error {
	if r.checkDatabaseForBookIdExisting(id) {
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
	return err != sql.ErrNoRows
}

func getInstanseTable(instanse models.InstanseMaker) string {
	instType := fmt.Sprintf("%T", instanse)
	return models.Tables[instType]
}
