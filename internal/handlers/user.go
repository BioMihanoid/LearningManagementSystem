package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/middleware"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
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

type passwordRequest struct {
	Password      string `json:"password"`
	ReplyPassword string `json:"reply_password"`
}

func (u *UserHandler) GetProfile(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
	}

	user, err := u.service.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get user " + err.Error()),
		})
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
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get users " + err.Error()),
		})
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserHandler) ChangeUserRole(c *gin.Context) {
	input := roleRequest{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
		return
	}

	user, err := u.service.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get user " + err.Error()),
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("user with id %d does not exist", userID),
		})
		return
	}

	err = u.service.ChangeUserRole(userID, input.RoleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error change role " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) UpdateFirstName(c *gin.Context) {
	input := profileRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
	}

	err = u.service.UpdateFirstName(userID, input.FirstName)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error update first name " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) UpdateLastName(c *gin.Context) {
	input := profileRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
	}

	err = u.service.UpdateLastName(userID, input.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error update last name " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) UpdateEmail(c *gin.Context) {
	input := profileRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
		return
	}

	err = u.service.UpdateEmail(userID, input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error update email " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) UpdatePassword(c *gin.Context) {
	input := passwordRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request  " + err.Error()),
		})
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
		return
	}

	err = u.service.UpdatePassword(userID, input.Password, input.ReplyPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error update password ") + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
		return
	}

	err = u.service.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error delete user " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) SubscribeOnCourse(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
		return
	}

	courseID, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}

	err = u.service.SubscribeOnCourse(userID, courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error subscribe to course"),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) UnsubscribeOnCourse(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
		return
	}

	courseID, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}

	err = u.service.UnsubscribeOnCourse(userID, courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error unsubscribe to course " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) GetAllUserSubscribedToTheCourse(c *gin.Context) {
	courseID, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}

	allUsers, err := u.service.GetAllUserSubscribedToTheCourse(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get all subscribed to the course " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, allUsers)
}

func (u *UserHandler) GetAllCoursesCurrentUser(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting user id " + err.Error()),
		})
		return
	}

	allCourses, err := u.service.GetAllCoursesCurrentUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get all subscribed to the course"),
		})
		return
	}

	c.JSON(http.StatusOK, allCourses)
}
