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

// @Summary Создать материал (Только для преподавателей)
// @Description Создает новый учебный материал для курса
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param course_id path int true "ID курса"
// @Param input body models.Material true "Данные материала"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/courses/{course_id}/materials [post]
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

// @Summary Получить материал по ID
// @Description Возвращает материал по его идентификатору
// @Tags materials
// @Security ApiKeyAuth
// @Produce json
// @Param material_id path int true "ID материала"
// @Success 200 {object} models.Material
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /courses/materials/{material_id} [get]
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

// @Summary Получить материалы курса
// @Description Возвращает все материалы указанного курса
// @Tags materials
// @Security ApiKeyAuth
// @Produce json
// @Param course_id path int true "ID курса"
// @Success 200 {array} models.Material
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /courses/{course_id}/materials [get]
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

// @Summary Обновить заголовок материала (Только для преподавателей)
// @Description Изменяет заголовок существующего материала
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param material_id path int true "ID материала"
// @Param title body string true "Новый заголовок"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/materials/{material_id}/update_title [post]
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

// @Summary Обновить содержимое материала (Только для преподавателей)
// @Description Изменяет содержимое существующего материала
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param material_id path int true "ID материала"
// @Param content body string true "Новое содержимое"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/materials/{material_id}/update_content [post]
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

// @Summary Удалить материал (Только для преподавателей)
// @Description Удаляет материал по его идентификатору
// @Tags teacher
// @Security ApiKeyAuth
// @Produce json
// @Param material_id path int true "ID материала"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /teacher/materials/{material_id} [delete]
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
