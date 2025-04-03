package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseHandler struct {
	services service.Service
}

func NewCourseHandler(services service.Service) *CourseHandler {
	return &CourseHandler{services}
}

type CourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// @Summary Создать новый курс (Только для преподавателей)
// @Description Создает новый учебный курс
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body CourseRequest true "Данные курса"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/courses [post]
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	input := CourseRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}
	course := models.Course{
		Title:       input.Title,
		Description: input.Description,
	}
	err := h.services.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error creating course " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// @Summary Получить курс по ID
// @Description Возвращает курс по id
// @Tags courses
// @Security ApiKeyAuth
// @Produce json
// @Param course_id path int true "ID курса"
// @Success 200 {object} models.Course
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /courses/{course_id} [get]
func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}
	course, err := h.services.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, course)
}

// @Summary Получить все курсы
// @Description Возвращает список всех доступных курсов
// @Tags courses
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Course
// @Failure 401 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /courses [get]
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.services.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting courses " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, courses)
}

// @Summary Обновить название курса (Только для преподавателей)
// @Description Изменяет название существующего курса
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param course_id path int true "ID курса"
// @Param title body string true "Новое название"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/courses/{course_id}/update_title [post]
func (h *CourseHandler) UpdateTitleCourse(c *gin.Context) {
	id, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}
	var input string
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}
	err = h.services.UpdateTitleCourse(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating course " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// @Summary Обновить описание курса (Только для преподавателей)
// @Description Изменяет описание существующего курса
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param course_id path int true "ID курса"
// @Param description body string true "Новое описание"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/courses/{course_id}/update_description [post]
func (h *CourseHandler) UpdateDescriptionCourse(c *gin.Context) {
	id, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}
	var input string
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}
	err = h.services.UpdateDescriptionCourse(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating course " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// @Summary Удалить курс (Только для преподавателей)
// @Description Удаляет курс по его идентификатору
// @Tags teacher
// @Security ApiKeyAuth
// @Produce json
// @Param course_id path int true "ID курса"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/courses/{course_id} [delete]
func (h *CourseHandler) DeleteCourseByID(c *gin.Context) {
	id, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}
	err = h.services.DeleteCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error deleting course " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
