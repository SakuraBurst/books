package crypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CryptPass(binaryPassword []byte) string {
	hash, err := bcrypt.GenerateFromPassword(binaryPassword, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
