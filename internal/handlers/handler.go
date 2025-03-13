package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	services service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: *services,
	}
}

var key = []byte("niggerspidors")

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

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// TODO: do auth with middleware and token
	authHandler := NewAuthHandler(h.services)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}

	authGroup := auth.Group("/")
	authGroup.Use(func(ctx *gin.Context) {
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
	}) // TODO: function check auth
	{
		userHandler := NewUserHandler(h.services)
		authGroup.GET("/profile", userHandler.GetProfile)
		authGroup.PUT("/profile", userHandler.UpdateProfile)

		courseHandler := NewCourseHandler(h.services)
		authGroup.GET("/courses", courseHandler.GetAllCourses)
		authGroup.GET("/courses/:id", courseHandler.GetCourseByID)

		lessonHandler := NewLessonHandler(h.services)
		authGroup.GET("/courses/:id/lessons", lessonHandler.GetCourseLesson)
		authGroup.GET("/lessons/:id", lessonHandler.GetLessonByID)

		testHandler := NewTestHandler(h.services)
		authGroup.GET("/test/:id", testHandler.GetTestByID)
		authGroup.POST("/test/:id/submit", testHandler.SubmitTest)

		adminGroup := authGroup.Group("/admin")
		adminGroup.Use() // TODO: function check role
		{
			authGroup.POST("/courses")
			authGroup.PUT("/courses/:id")
			authGroup.DELETE("/courses/:id")

			authGroup.POST("/lessons", lessonHandler.CreateLesson)
			authGroup.PUT("/lessons/:id", lessonHandler.UpdateLesson)
			authGroup.DELETE("/lessons/:id", lessonHandler.DeleteLesson)

			authGroup.GET("/users", userHandler.GetAllUsers)
			authGroup.PUT("/users/:user_id", userHandler.ChangeUserRole)
		}
	}

	return router
}
