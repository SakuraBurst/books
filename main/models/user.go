package models

import "github.com/SakuraBurst/books.git/main/helpers/checks"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u User) IsValid() bool {
	return checks.IsStringLengthMoreThanZero(u.FirstName) &&
		checks.IsStringLengthMoreThanZero(u.LastName) &&
		checks.IsStringLengthMoreThanZero(u.Email) &&
		checks.IsStringLengthMoreThanZero(u.Password)
}
