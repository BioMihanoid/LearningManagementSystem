package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Handler struct {
	services service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: *services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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
			userId, err := pkg.GetUserIdFromJWT(parsed[1])
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}
			ctx.Set("userId", userId)
		}
		ctx.Next()
	})
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

		teacherGroup := authGroup.Group("/teacher")
		//TODO: create new check role fn
		teacherGroup.Use()
		{
			teacherGroup.POST("/courses")
			teacherGroup.PUT("/courses/:id")
			teacherGroup.DELETE("/courses/:id")

			teacherGroup.POST("/lessons", lessonHandler.CreateLesson)
			teacherGroup.PUT("/lessons/:id", lessonHandler.UpdateLesson)
			teacherGroup.DELETE("/lessons/:id", lessonHandler.DeleteLesson)
		}

		adminGroup := authGroup.Group("/admin")
		//TODO: create new check role fn
		adminGroup.Use()
		{
			adminGroup.GET("/users", userHandler.GetAllUsers)
			adminGroup.PUT("/users/:user_id", userHandler.ChangeUserRole)
		}
	}

	return router
}
