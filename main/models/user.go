package models

import (
	"encoding/json"
	"io"

	"github.com/SakuraBurst/books.git/main/helpers/checks"
)

type User struct {
	ID        int
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

func (u User) NewInstanseFromJson(body io.ReadCloser) InstanseMaker {
	decoder := json.NewDecoder(body)
	decoder.Decode(&u)
	return u
}

func (u User) NewInstanseFromDB(row Scans) (InstanseMaker, error) {
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
