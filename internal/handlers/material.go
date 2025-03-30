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

type MaterialHandler struct {
	services service.Service
}

func NewMaterialHandler(services service.Service) *MaterialHandler {
	return &MaterialHandler{services: services}
}

func (m *MaterialHandler) CreateMaterial(c *gin.Context) {
	input := models.Material{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
	}
	input.CourseID = GetCourseIDParam(c)
	if err := m.services.CreateMaterial(input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func (m *MaterialHandler) GetMaterialByID(c *gin.Context) {
	id := GetMaterialIDParam(c)
	material, err := m.services.GetMaterialByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, material)
}

func (m *MaterialHandler) GetCourseMaterial(c *gin.Context) {
	id := GetCourseIDParam(c)
	materials, err := m.services.GetCourseMaterial(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, materials)
}

func (m *MaterialHandler) UpdateTitleMaterial(c *gin.Context) {
	id := GetMaterialIDParam(c)
	var input string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
	}
	err := m.services.UpdateTitleMaterial(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func (m *MaterialHandler) UpdateContentMaterial(c *gin.Context) {
	id := GetMaterialIDParam(c)
	var input string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
	}
	err := m.services.UpdateTitleMaterial(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func (m *MaterialHandler) DeleteMaterial(c *gin.Context) {
	id := GetMaterialIDParam(c)
	err := m.services.DeleteMaterial(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func GetMaterialIDParam(c *gin.Context) int {
	v, _ := c.Get("material_id")

	id, _ := strconv.Atoi(v.(string))

	return id
}
