package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.Service
}

func NewUserHandler(service service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) GetProfile(c *gin.Context) {}

func (u *UserHandler) UpdateProfile(c *gin.Context) {}
