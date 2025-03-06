package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.Service
}

func NewAuthHandler(service service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {}

func (h *AuthHandler) Login(c *gin.Context) {}

func (h *AuthHandler) RefreshToken(c *gin.Context) {}
