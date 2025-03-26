package pkg

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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

func GetAccessWithToken(ctx *gin.Context) {
	authHeaderValue := ctx.GetHeader("Authorization")
	parsed := strings.Split(authHeaderValue, " ")
	if len(parsed) > 1 && parsed[0] == "Bearer" {
		userId, err := GetUserIdFromJWT(parsed[1])
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Set("userId", userId)
	}
	ctx.Next()
}

func GetAccessRole(levelAccess int, service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Get("userId")
		userId, _ := strconv.Atoi(c.GetString("userId"))

		user, err := service.GetUserById(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": err.Error()})
			c.Abort()
			return
		}

		level, err := service.GetLevelAccess(user.RoleID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": err.Error()})
			c.Abort()
			return
		}

		if level < levelAccess {
			c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
