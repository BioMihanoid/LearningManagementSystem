package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
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

	authHandler := NewAuthHandler(h.services)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}

	authGroup := auth.Group("/")
	authGroup.Use(pkg.GetAccessWithToken)
	{
		userHandler := NewUserHandler(h.services)
		authGroup.GET("/profile", userHandler.GetProfile)
		authGroup.POST("/profile/update_first_name", userHandler.UpdateFirstName)
		authGroup.POST("/profile/update_last_name", userHandler.UpdateLastName)
		authGroup.POST("/profile/update_email", userHandler.UpdateEmail)
		authGroup.POST("profile/change_password", userHandler.UpdatePassword)

		courseHandler := NewCourseHandler(h.services)
		authGroup.GET("/courses/:course_id", courseHandler.GetCourseByID)
		authGroup.GET("/courses", courseHandler.GetAllCourses)

		authGroup.POST("/subscribe/:course_id", userHandler.SubscribeOnCourse)
		authGroup.POST("/unsubscribe/:course_id", userHandler.UnsubscribeOnCourse)
		authGroup.GET("/profile/subscriptions", userHandler.GetAllCoursesCurrentUser)

		materialHandler := NewMaterialHandler(h.services)
		authGroup.GET("/courses/:course_id/material/:material_id", materialHandler.GetMaterialByID)
		authGroup.GET("/courses/:course_id/materials", materialHandler.GetCourseMaterial)

		testHandler := NewTestHandler(h.services)
		authGroup.GET("/test/:id", testHandler.GetTestByID)
		authGroup.POST("/test/:id/submit", testHandler.SubmitTest)

		teacherGroup := authGroup.Group("/teacher")
		teacherGroup.Use(pkg.GetAccessRole(1, h.services))
		{
			teacherGroup.POST("/courses", courseHandler.CreateCourse)
			teacherGroup.POST("/courses/:course_id/update_title", courseHandler.UpdateTitleCourse)
			teacherGroup.POST("/courses/:course_id/update_description", courseHandler.UpdateDescriptionCourse)
			teacherGroup.DELETE("/courses/:course_id", courseHandler.DeleteCourseByID)

			teacherGroup.POST("/courses/:course_id/material", materialHandler.CreateMaterial)
			teacherGroup.POST("/courses/:course_id/material/:material_id/update_title", materialHandler.UpdateTitleMaterial)
			teacherGroup.POST("/courses/:course_id/material/:material_id/update_content", materialHandler.UpdateContentMaterial)
			teacherGroup.DELETE("/courses/:course_id/material/:material_id", materialHandler.DeleteMaterial)
		}

		adminGroup := authGroup.Group("/admin")
		adminGroup.Use(pkg.GetAccessRole(3, h.services))
		{
			adminGroup.GET("/users", userHandler.GetAllUsers)
			adminGroup.GET("users/subscriptions/:course_id", userHandler.GetAllUserSubscribedToTheCourse)
			adminGroup.POST("/users/:user_id", userHandler.ChangeUserRole)
			adminGroup.DELETE("/users/:user_id", userHandler.DeleteUser)
		}
	}

	return router
}
