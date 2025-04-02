package handlers

import (
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/middleware"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
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
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}
	courseID, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}
	input.CourseID = courseID
	err = t.service.CreateTest(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error creating test " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusCreated, "ok")
}

func (t *TestHandler) GetTestByID(c *gin.Context) {
	testID, err := pkg.GetID(c, "test_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test id " + err.Error()),
		})
		return
	}
	test, err := t.service.GetTestByID(testID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, test)
}

func (t *TestHandler) GetAllTestsCourse(c *gin.Context) {
	courseID, err := pkg.GetID(c, "course_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting course id " + err.Error()),
		})
		return
	}
	tests, err := t.service.GetAllTestsCourse(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting tests " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, tests)
}

func (t *TestHandler) GetAllTests(c *gin.Context) {
	tests, err := t.service.GetAllTests()
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting tests " + err.Error()),
		})
	}
	c.JSON(http.StatusOK, tests)
}

func (t *TestHandler) UpdateQuestionTest(c *gin.Context) {
	testID, err := pkg.GetID(c, "test_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test id " + err.Error()),
		})
		return
	}
	input := models.Test{}
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}
	err = t.service.UpdateQuestionTest(testID, input.Question)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating question " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) UpdateAnswerTest(c *gin.Context) {
	testID, err := pkg.GetID(c, "test_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test id " + err.Error()),
		})
		return
	}
	input := models.Test{}
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}
	err = t.service.UpdateAnswerTest(testID, input.Question)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating question " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) DeleteTest(c *gin.Context) {
	testID, err := pkg.GetID(c, "test_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test id " + err.Error()),
		})
		return
	}
	err = t.service.DeleteTest(testID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error deleting test " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) SubmitTest(c *gin.Context) {
	input := UserAnswerRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error parsing request " + err.Error()),
		})
		return
	}

	testID, err := pkg.GetID(c, "test_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test id " + err.Error()),
		})
		return
	}

	test, err := t.service.GetTestByID(testID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test " + err.Error()),
		})
		return
	}
	result := t.service.SubmitTest(test.Answer, input.Answer)

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting userID " + err.Error()),
		})
		return
	}

	err = t.service.CreateTestResult(userID, test.ID, result)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error creating test " + err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (t *TestHandler) GetTestResult(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting userID " + err.Error()),
		})
		return
	}

	testResultID, err := pkg.GetID(c, "test_result_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting testResultID " + err.Error()),
		})
		return
	}

	testRes, err := t.service.GetTestResult(userID, testResultID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting test " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, testRes.Score)
}

func (t *TestHandler) UpdateTestResult(c *gin.Context) {
	testResultID, err := pkg.GetID(c, "test_result_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting testResultID " + err.Error()),
		})
		return
	}

	userId, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting userID " + err.Error()),
		})
		return
	}

	err = t.service.UpdateTestResult(userId, testResultID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error updating testResult " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (t *TestHandler) DeleteTestResult(c *gin.Context) {
	testResultID, err := pkg.GetID(c, "test_result_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting testResultID " + err.Error()),
		})
		return
	}

	err = t.service.DeleteTestResult(testResultID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error deleting testResult " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}
