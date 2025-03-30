package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/models"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	input := CourseRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
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
			Error: fmt.Sprintf("error creating course"),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id := GetCourseIDParam(c)
	course, err := h.services.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course"),
		})
	}
	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.services.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting courses"),
		})
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) UpdateTitleCourse(c *gin.Context) {
	id := GetCourseIDParam(c)
	var input string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
		return
	}
	err := h.services.UpdateTitleCourse(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating course"),
		})
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *CourseHandler) UpdateDescriptionCourse(c *gin.Context) {
	id := GetCourseIDParam(c)
	var input string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
		return
	}
	err := h.services.UpdateDescriptionCourse(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating course"),
		})
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *CourseHandler) DeleteCourseByID(c *gin.Context) {
	id := GetCourseIDParam(c)
	err := h.services.DeleteCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error deleting course"),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func GetCourseIDParam(c *gin.Context) int {
	v, _ := c.Get("course_id")

	id, _ := strconv.Atoi(v.(string))

	return id
}
