package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var key = []byte("niggerspidors") //TODO :fix

func GenerateJWT(userId string, timeEnd time.Time) (string, error) {
	claims := &jwt.StandardClaims{ExpiresAt: timeEnd.Unix(), Subject: userId}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func GetUserIdFromJWT(tokenString string) (string, error) {
	draft, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)

	if err != nil {
		return "", err
	}

	if draft.Valid {
		id := draft.Claims.(*jwt.StandardClaims).Subject
		return id, nil
	}

	return "", err
}
