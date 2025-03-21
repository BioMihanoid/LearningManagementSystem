package pkg

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func AuthorizeRole(requiredRole string, service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, _ := c.Get("userId")

		id, _ := strconv.Atoi(v.(string))

		user, err := service.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "not admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
