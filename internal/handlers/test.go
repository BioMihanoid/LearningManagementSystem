package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	service service.Service
}

func NewTestHandler(service service.Service) *TestHandler {
	return &TestHandler{service: service}
}

func (t *TestHandler) GetTestByID(c *gin.Context) {}
func (t *TestHandler) SubmitTest(c *gin.Context)  {}
