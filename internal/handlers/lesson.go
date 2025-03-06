package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	services service.Service
}

func NewLessonHandler(services service.Service) *LessonHandler {
	return &LessonHandler{services: services}
}

func (l *LessonHandler) GetCourseLesson(c *gin.Context) {}
func (l *LessonHandler) GetLessonByID(c *gin.Context)   {}
func (l *LessonHandler) CreateLesson(c *gin.Context)    {}
func (l *LessonHandler) UpdateLesson(c *gin.Context)    {}
func (l *LessonHandler) DeleteLesson(c *gin.Context)    {}
