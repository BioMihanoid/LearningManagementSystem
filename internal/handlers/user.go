package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type profileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type roleRequest struct {
	RoleID int `json:"role"`
}

type authHeader struct {
	Token string `header:"Authorization"`
}

func (u *UserHandler) GetProfile(c *gin.Context) {
	id := GetUserID(c)

	user, err := u.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
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

	err = u.service.ChangeUserRole(id, input.RoleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetUserIDParam(c *gin.Context) int {
	v, _ := c.Get("userId")

	id, _ := strconv.Atoi(v.(string))

	return id
}

func GetUserID(c *gin.Context) int {
	h := authHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	token := strings.Split(h.Token, "Bearer ")[1]

	strID, err := pkg.GetUserIdFromJWT(token)
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}

	return id
}

func (u *UserHandler) UpdateFirstName(c *gin.Context) {
	input := profileRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
		return
	}

	id := GetUserID(c)
	err := u.service.UpdateFirstName(id, input.FirstName)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error update first name"),
		})
		return
	}
}

func (u *UserHandler) UpdateLastName(c *gin.Context) {
	input := profileRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
	}

	id := GetUserID(c)
	err := u.service.UpdateLastName(id, input.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error update last name"),
		})
		return
	}
}

func (u *UserHandler) UpdateEmail(c *gin.Context) {
	input := profileRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
		return
	}
	id := GetUserID(c)
	err := u.service.UpdateEmail(id, input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error update email"),
		})
		return
	}
}
