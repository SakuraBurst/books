package models

import (
	"encoding/json"
	"io"
)

type Book struct {
	Title  string
	ID     int
	Author string
	Year   string
}

func (book Book) IsValid() bool {
	return len(book.Author) > 0 && len(book.Title) > 0 && len(book.Year) > 0
}

func (book Book) NewInstanseFromJson(body io.ReadCloser) InstanseMaker {
	decoder := json.NewDecoder(body)
	decoder.Decode(&book)
	return book
}

func (book Book) NewInstanseFromDB(row Scans) (InstanseMaker, error) {
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return Book{}, err
	}
	newBook := book
	return newBook, nil
}
