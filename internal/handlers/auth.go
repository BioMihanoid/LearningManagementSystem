package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/models"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
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

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Token string `json:"token" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	input := RegisterRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request: %s", err.Error()),
		})
		return
	}

	user, err := h.service.GetUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get user"),
		})
		return
	}

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("user with email %s exists", input.Email),
		})
		return
	}

	// TODO: validate data: validator

	user = models.User{
		Name:     input.Username,
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error create user"),
		})
		return
	}

	jwt, err := GenerateJWT(strconv.Itoa(id), time.Now().Add(30*time.Minute))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error create jwt"),
		})
		return
	}

	output := RegisterResponse{Token: jwt}

	c.JSON(http.StatusOK, output)
}

func (h *AuthHandler) Login(c *gin.Context) {
	input := LoginRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
	}

	user, err := h.service.GetUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get user"),
		})
		return
	}

	jwt, err := GenerateJWT(strconv.Itoa(int(user.ID)), time.Now().Add(30*time.Minute))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error create jwt"),
		})
		return
	}

	output := LoginResponse{
		Token: jwt,
	}

	c.JSON(http.StatusOK, output)
}
