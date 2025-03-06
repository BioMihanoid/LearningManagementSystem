package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
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

	// TODO: do auth with middleware and token
	authHandler := NewAuthHandler(h.services)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}

	authGroup := auth.Group("/")
	authGroup.Use() // TODO: function check auth
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
