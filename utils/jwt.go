package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const myKey = "helloWorld"

func GenerateToken(userEmail string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  userEmail,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(myKey))
}

func VerifyToken(token string) {
	jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return myKey, nil
	})

	// if err != nil {
	// 	return errors.New("Could not parse token")
	// }

}
