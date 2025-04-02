package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
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
			Error: fmt.Sprintf("error parsing request " + err.Error()),
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
	input.CourseID = courseID
	if err = m.services.CreateMaterial(input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error creating material " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (m *MaterialHandler) GetMaterialByID(c *gin.Context) {
	id, err := pkg.GetID(c, "material_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting material id " + err.Error()),
		})
		return
	}
	material, err := m.services.GetMaterialByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting material " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, material)
}

func (m *MaterialHandler) GetCourseMaterial(c *gin.Context) {
	courseID, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}
	materials, err := m.services.GetCourseMaterial(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting material " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, materials)
}

func (m *MaterialHandler) UpdateTitleMaterial(c *gin.Context) {
	materialID, err := pkg.GetID(c, "material_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting material id " + err.Error()),
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
	err = m.services.UpdateTitleMaterial(materialID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating material " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (m *MaterialHandler) UpdateContentMaterial(c *gin.Context) {
	materialID, err := pkg.GetID(c, "material_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting material id " + err.Error()),
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
	err = m.services.UpdateContentMaterial(materialID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating material " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (m *MaterialHandler) DeleteMaterial(c *gin.Context) {
	materialID, err := pkg.GetID(c, "material_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting material id " + err.Error()),
		})
		return
	}
	err = m.services.DeleteMaterial(materialID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error deleting material " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
