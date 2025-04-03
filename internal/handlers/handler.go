package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/middleware"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/BioMihanoid/LearningManagementSystem/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authHandler := NewAuthHandler(h.services)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}

	authGroup := auth.Group("/")
	authGroup.Use(middleware.GetAccessWithToken)
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
		authGroup.GET("/courses/:course_id/materials/:material_id", materialHandler.GetMaterialByID)
		authGroup.GET("/courses/:course_id/materials", materialHandler.GetCourseMaterial)

		testHandler := NewTestHandler(h.services)
		authGroup.GET("/courses/:course_id/tests/:test_id", testHandler.GetTestByID)
		authGroup.POST("/courses/:course_id/tests/:test_id/submit", testHandler.SubmitTest)

		authGroup.GET("/courses/:course_id/test_results/:test_result_id", testHandler.GetTestResult)

		teacherGroup := authGroup.Group("/teacher")
		teacherGroup.Use(middleware.GetAccessRole(1, h.services))
		{
			teacherGroup.POST("/courses", courseHandler.CreateCourse)
			teacherGroup.POST("/courses/:course_id/update_title", courseHandler.UpdateTitleCourse)
			teacherGroup.POST("/courses/:course_id/update_description", courseHandler.UpdateDescriptionCourse)
			teacherGroup.DELETE("/courses/:course_id", courseHandler.DeleteCourseByID)

			teacherGroup.POST("/courses/:course_id/materials", materialHandler.CreateMaterial)
			teacherGroup.POST("/courses/:course_id/materials/:material_id/update_title", materialHandler.UpdateTitleMaterial)
			teacherGroup.POST("/courses/:course_id/materials/:material_id/update_content", materialHandler.UpdateContentMaterial)
			teacherGroup.DELETE("/courses/:course_id/materials/:material_id", materialHandler.DeleteMaterial)

			teacherGroup.POST("/courses/:course_id/tests/create_test", testHandler.CreateTest)
			teacherGroup.DELETE("/courses/:course_id/tests/:test_id/delete_test", testHandler.DeleteTest)

			teacherGroup.POST("/courses/:course_id/test_results/:test_result_id/update_result", testHandler.UpdateTestResult)
			teacherGroup.DELETE("/courses/:course_id/test_results/:test_result_id/delete_test_result", testHandler.DeleteTestResult)
		}

		adminGroup := authGroup.Group("/admin")
		adminGroup.Use(middleware.GetAccessRole(3, h.services))
		{
			adminGroup.GET("/users", userHandler.GetAllUsers)
			adminGroup.GET("users/subscriptions/:course_id", userHandler.GetAllUserSubscribedToTheCourse)
			adminGroup.POST("/users/:user_id", userHandler.ChangeUserRole)
			adminGroup.DELETE("/users/:user_id", userHandler.DeleteUser)
		}
	}

	return router
}
