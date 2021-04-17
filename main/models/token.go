package models

import (
	"fmt"
	"time"

	"github.com/SakuraBurst/books.git/main/driver"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(user User) string {
	now := time.Now().Unix()
	daySeconds := time.Hour.Seconds() * 24
	fmt.Println(now)
	// Create the Claims
	claims := UserClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: now + int64(daySeconds),
			Issuer:    "ya",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey, _ := KeyFunc(token)
	ss, _ := token.SignedString(mySigningKey)
	return ss

}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	key := driver.GetEnv("JWT_KEY")
	return []byte(key), nil
}
