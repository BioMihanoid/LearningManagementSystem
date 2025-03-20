package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service service.Service
}

func NewUserHandler(service service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type profileResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// for updateUser
type profileRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type roleRequest struct {
	Role string `json:"role"`
}

func (u *UserHandler) GetProfile(c *gin.Context) {
	v, _ := c.Get("userId")

	id, _ := strconv.Atoi(v.(string))

	user, err := u.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profileResponse{
		Username: user.Name,
		Email:    user.Email,
		Role:     user.Role,
	})
}

func (u *UserHandler) UpdateProfile(c *gin.Context) {

}

func (u *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (u *UserHandler) ChangeUserRole(c *gin.Context) {
	input := roleRequest{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
		return
	}

	userID := c.Param("user_id")
	id, _ := strconv.Atoi(userID)

	user, err := u.service.GetUserById(id)

	if err != nil || user.ID == 0 {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("user with id %d does not exist", id),
		})
		return
	}

	err = u.service.ChangeUserRole(id, input.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
