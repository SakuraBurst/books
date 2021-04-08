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

func (book *Book) NewInstanseFromJson(body io.ReadCloser) {
	decoder := json.NewDecoder(body)
	decoder.Decode(book)
}
