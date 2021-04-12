package models

import (
	"encoding/json"
	"errors"
	"io"
)

type Book struct {
	Title  string `json:"title"`
	ID     int    `json:"id"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func (book Book) IsValid() bool {
	return len(book.Author) > 0 && len(book.Title) > 0 && len(book.Year) > 0
}

func (book Book) NewInstanseFromJson(body io.ReadCloser) (InstanseMaker, error) {
	decoder := json.NewDecoder(body)
	decoder.Decode(&book)
	if book.IsValid() {
		return book, nil
	} else {
		return nil, errors.New("some data is unprocessable")
	}
}

func (book Book) NewInstanseFromDB(row Scans) (InstanseMaker, error) {
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return nil, err
	}
	return book, nil
}
