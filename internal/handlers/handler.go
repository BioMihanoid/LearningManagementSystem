package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{
		services: services,
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

		// courseHandler
		authGroup.GET("/courses")     // all courses
		authGroup.GET("/courses/:id") // courses by id

		// lessonHandler
		authGroup.GET("/courses/:id/lessons") // get lessons course
		authGroup.GET("/lessons/:id")         //get lesson by id

		// testHandler
		authGroup.GET("/test/:id")         // get test by id
		authGroup.POST("/test/:id/submit") // get result test by id

		// admin routes
		adminGroup := authGroup.Group("/admin")
		adminGroup.Use() // TODO: function check role
		{
			// courseHandler
			authGroup.POST("/courses")       // create course
			authGroup.PUT("/courses/:id")    // update course
			authGroup.DELETE("/courses/:id") // delete course

			// lessonHandler
			authGroup.POST("/lessons")       // create lesson
			authGroup.PUT("/lessons/:id")    // update lesson
			authGroup.DELETE("/lessons/:id") // delete lesson

			// userHandler
			authGroup.GET("/users")          // get all users
			authGroup.PUT("/users/:user_id") // change role user by id
		}
	}

	return router
}
