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

// @Summary Получить профиль пользователя
// @Description Возвращает информацию о текущем аутентифицированном пользователе
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} profileResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /auth/profile [get]
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

// @Summary Получить всех пользователей (Только для администратора)
// @Description Возвращает список всех зарегистрированных пользователей
// @Tags admin
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 403 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /auth/admin/users [get]
func (u *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error get users " + err.Error()),
		})
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Изменить роль пользователя (Только для администратора)
// @Description Обновляет роль указанного пользователя
// @Tags admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "ID пользователя"
// @Param input body roleRequest true "Новая роль"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /auth/admin/users/{user_id} [post]
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

// @Summary Обновить имя пользователя
// @Description Изменяет имя текущего пользователя
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body profileRequest true "Новое имя"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/profile/update_first_name [post]
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

// @Summary Обновить фамилию пользователя
// @Description Изменяет фамилию текущего пользователя
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body profileRequest true "Новая фамилия"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/profile/update_last_name [post]
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

// @Summary Обновить email пользователя
// @Description Изменяет email текущего пользователя
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body profileRequest true "Новый email"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/profile/update_email [post]
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

// @Summary Обновить пароль пользователя
// @Description Изменяет пароль текущего пользователя
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body passwordRequest true "Новый пароль"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/profile/change_password [post]
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

// @Summary Удалить пользователя (Только для администратора)
// @Description Удаляет указанного пользователя из системы
// @Tags admin
// @Security ApiKeyAuth
// @Produce json
// @Param user_id path int true "ID пользователя"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Router /auth/admin/users/{user_id} [delete]
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

// @Summary Подписаться на курс
// @Description Добавляет текущего пользователя в подписчики курса
// @Tags subscriptions
// @Security ApiKeyAuth
// @Produce json
// @Param course_id path int true "ID курса"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/subscribe/{course_id} [post]
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

// @Summary Отписаться от курса
// @Description Удаляет текущего пользователя из подписчиков курса
// @Tags subscriptions
// @Security ApiKeyAuth
// @Produce json
// @Param course_id path int true "ID курса"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/unsubscribe/{course_id} [post]
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

// @Summary Получить подписчиков курса
// @Description Возвращает список пользователей, подписанных на указанный курс
// @Tags courses
// @Security ApiKeyAuth
// @Produce json
// @Param course_id path int true "ID курса"
// @Success 200 {array} models.User
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/users/subscriptions/{course_id} [get]
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

// @Summary Получить курсы пользователя
// @Description Возвращает список всех курсов, на которые подписан текущий пользователь
// @Tags subscriptions
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Course
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /auth/profile/subscriptions [get]
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
