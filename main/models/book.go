package models

type Book struct {
	Title  string
	ID     int
	Author string
	Year   string
}

func (book Book) IsBookValid() bool {
	return len(book.Author) > 0 && len(book.Title) > 0 && len(book.Year) > 0
}
