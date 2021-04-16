package models

import (
	"encoding/json"
	"fmt"
	"io"

	"errors"

	"github.com/SakuraBurst/books.git/main/helpers/checks"
	"github.com/SakuraBurst/books.git/main/helpers/crypt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

type UserResponse struct {
	Token string        `json:"token"`
	User  InstanseMaker `json:"user"`
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
	fmt.Println(u)
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

func (u *User) DeletePasswordField() {
	u.Password = ""
}

func (u User) ComparePassword(pass []byte) bool {
	fmt.Println(u)
	return bcrypt.CompareHashAndPassword([]byte(u.Password), pass) == nil
}
