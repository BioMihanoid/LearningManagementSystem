package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AuthHandler struct {
	service service.Service
}

func NewAuthHandler(service service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

var RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	input := RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid request format")})
		return
	}

	// TODO: validate data: validator

	u := models.User{
		Name:     input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	_, err := h.service.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// TODO: check existing user in db

	// TODO: do createUser

	userId := strconv.Itoa(int(u.ID))

	jwt, _ := GenerateJWT(userId, time.Now().Add(15*time.Minute))

	c.JSON(http.StatusOK, jwt)
}

func (h *AuthHandler) Login(c *gin.Context) {

}

func (h *AuthHandler) RefreshToken(c *gin.Context) {}
