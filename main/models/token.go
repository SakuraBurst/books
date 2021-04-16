package models

import (
	"fmt"
	"time"

	"github.com/SakuraBurst/books.git/main/driver"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken() {
	now := time.Now().Unix()
	daySeconds := time.Hour.Seconds() * 24
	fmt.Println(now)
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: now + int64(daySeconds),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey, _ := KeyFunc(token)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Println(ss)

	// tokens, err := jwt.Parse(ss, KeyFunc)
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	key := driver.GetEnv("JWT_KEY")
	return []byte(key), nil
}
