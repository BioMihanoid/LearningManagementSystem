package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/middleware"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
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

// @Summary Регистрация нового пользователя
// @Description Создает новую учетную запись пользователя
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body RegisterRequest true "Данные для регистрации"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	input := RegisterRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request: %s " + err.Error()),
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error incorect data " + err.Error()),
		})
		return
	}

	user, err := h.service.GetUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get user " + err.Error()),
		})
		return
	}

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("user with email %s exists", input.Email),
		})
		return
	}

	validate = validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error incorect data " + err.Error()),
		})
		return
	}

	user = models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		RoleID:    1,
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error create user " + err.Error()),
		})
		return
	}

	jwt, err := middleware.GenerateJWT(strconv.Itoa(id), time.Now().Add(30*time.Minute))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error create jwt " + err.Error()),
		})
		return
	}

	output := RegisterResponse{Token: jwt}

	c.JSON(http.StatusOK, output)
}

// @Summary Аутентификация пользователя
// @Description Вход в систему с использованием email и пароля
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body LoginRequest true "Данные для входа"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	input := LoginRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error incorect data " + err.Error()),
		})
		return
	}

	user, err := h.service.GetUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get user  " + err.Error()),
		})
		return
	}

	jwt, err := middleware.GenerateJWT(strconv.Itoa(user.ID), time.Now().Add(30*time.Minute))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error create jwt " + err.Error()),
		})
		return
	}

	output := LoginResponse{
		Token: jwt,
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Обновление токена доступа
// @Description Генерирует новый JWT токен для авторизованного пользователя
// @Tags auth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} LoginResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get userID " + err.Error()),
		})
	}

	jwt, err := middleware.GenerateJWT(strconv.Itoa(userID), time.Now().Add(30*time.Minute))

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error create jwt  " + err.Error()),
		})
		return
	}

	output := LoginResponse{
		Token: jwt,
	}

	c.JSON(http.StatusOK, output)
}
