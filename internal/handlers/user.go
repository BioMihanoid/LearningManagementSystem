package handlers

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service service.Service
}

func NewUserHandler(service service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) GetProfile(c *gin.Context) {
	v, _ := c.Get("userId")
	c.JSON(http.StatusOK, gin.H{
		"userId": v,
	})
}

func (u *UserHandler) UpdateProfile(c *gin.Context) {}

func (u *UserHandler) GetAllUsers(c *gin.Context) {}

func (u *UserHandler) ChangeUserRole(c *gin.Context) {}
