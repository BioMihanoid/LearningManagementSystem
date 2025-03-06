package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	services service.Service
}

func NewCourseHandler(services service.Service) *CourseHandler {
	return &CourseHandler{services}
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {}
func (h *CourseHandler) GetCourseByID(c *gin.Context) {}
func (h *CourseHandler) CreateCourse(c *gin.Context)  {}
func (h *CourseHandler) UpdateCourse(c *gin.Context)  {}
func (h *CourseHandler) DeleteCourse(c *gin.Context)  {}
