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

type TestHandler struct {
	service service.Service
}

func NewTestHandler(service service.Service) *TestHandler {
	return &TestHandler{service: service}
}

type UserAnswerRequest struct {
	Answer string `json:"answer" binding:"required"`
}

func (t *TestHandler) CreateTest(c *gin.Context) {
	input := models.Test{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
		return
	}
	input.ID = GetCourseIDParam(c)
	err := t.service.CreateTest(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "ok")
}

func (t *TestHandler) GetTestByID(c *gin.Context) {
	id := GetUserID(c)
	test, err := t.service.GetTestByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{})
	}
	c.JSON(http.StatusOK, test)
}

func (t *TestHandler) GetAllTestsCourse(c *gin.Context) {
	id := GetCourseIDParam(c)
	tests, err := t.service.GetAllTestsCourse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, tests)
}

func (t *TestHandler) GetAllTests(c *gin.Context) {
	tests, err := t.service.GetAllTests()
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, tests)
}

func (t *TestHandler) UpdateQuestionTest(c *gin.Context) {
	id := GetTestIDParam(c)
	input := models.Test{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
	}
	err := t.service.UpdateQuestionTest(id, input.Question)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) UpdateAnswerTest(c *gin.Context) {
	id := GetTestIDParam(c)
	input := models.Test{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request"),
		})
	}
	err := t.service.UpdateAnswerTest(id, input.Answer)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) DeleteTest(c *gin.Context) {
	id := GetTestIDParam(c)
	err := t.service.DeleteTest(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) SubmitTest(c *gin.Context) {
	input := UserAnswerRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{})
	}
	test, err := t.service.GetTestByID(GetTestIDParam(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	result := t.service.SubmitTest(test.Answer, input.Answer)
	err = t.service.CreateTestResult(GetUserID(c), test.ID, result)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (t *TestHandler) GetTestResult(c *gin.Context) {
	testRes, err := t.service.GetTestResult(GetUserID(c), GetTestResultIDParam(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, testRes.Score)
}

func (t *TestHandler) UpdateTestResult(c *gin.Context) {
	testResultID := GetTestResultIDParam(c)
	err := t.service.UpdateTestResult(GetUserID(c), testResultID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) DeleteTestResult(c *gin.Context) {
	testResultID := GetTestResultIDParam(c)
	err := t.service.DeleteTestResult(testResultID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "ok")
}

func GetTestIDParam(c *gin.Context) int {
	v, _ := c.Get("test_id")
	id, _ := strconv.Atoi(v.(string))
	return id
}

func GetTestResultIDParam(c *gin.Context) int {
	v, _ := c.Get("test_result_id")
	id, _ := strconv.Atoi(v.(string))
	return id
}
