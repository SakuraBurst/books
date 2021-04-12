package models

import (
	"encoding/json"
	"io"

	"errors"

	"github.com/SakuraBurst/books.git/main/helpers/checks"
	"github.com/SakuraBurst/books.git/main/helpers/crypt"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type UserResponse struct {
	User InstanseMaker `json:"user"`
}

func (u User) IsValid() bool {
	return checks.IsStringLengthMoreThanZero(u.FirstName) &&
		checks.IsStringLengthMoreThanZero(u.LastName) &&
		checks.IsStringLengthMoreThanZero(u.Email) &&
		checks.IsStringLengthMoreThanZero(u.Password)
}

func (u User) NewInstanseFromJson(body io.ReadCloser) (InstanseMaker, error) {
	decoder := json.NewDecoder(body)
	decoder.Decode(&u)
	if u.IsValid() {
		u.Password = crypt.CryptPass([]byte(u.Password))
		return u, nil
	} else {
		return nil, errors.New("some data is unprocessable")
	}

}

func (u User) NewInstanseFromDB(row Scans) (InstanseMaker, error) {
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return u, nil
}
