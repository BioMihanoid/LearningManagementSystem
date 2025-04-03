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

// @Summary Создать тест (Только для преподавателей)
// @Description Создает новый тест для курса
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param course_id path int true "ID курса"
// @Param input body models.Test true "Данные теста"
// @Success 201
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Router /teacher/courses/{course_id}/tests [post]
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

// @Summary Получить тест по ID
// @Description Возвращает тест по его идентификатору
// @Tags tests
// @Security ApiKeyAuth
// @Produce json
// @Param test_id path int true "ID теста"
// @Success 200 {object} models.Test
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /tests/{test_id} [get]
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

// @Summary Получить тесты курса
// @Description Возвращает все тесты для указанного курса
// @Tags tests
// @Security ApiKeyAuth
// @Produce json
// @Param course_id path int true "ID курса"
// @Success 200 {array} models.Test
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /courses/{course_id}/tests [get]
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

// @Summary Получить все тесты
// @Description Возвращает список всех доступных тестов
// @Tags tests
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Test
// @Failure 500 {object} pkg.ErrorResponse
// @Router /tests [get]
func (t *TestHandler) GetAllTests(c *gin.Context) {
	tests, err := t.service.GetAllTests()
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Error: fmt.Sprintf("error getting tests " + err.Error()),
		})
	}
	c.JSON(http.StatusOK, tests)
}

// @Summary Обновить вопрос теста (Только для преподавателей)
// @Description Изменяет вопрос существующего теста
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param test_id path int true "ID теста"
// @Param input body models.Test true "Новый вопрос"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Router /teacher/tests/{test_id}/question [put]
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

// @Summary Обновить ответ теста (Только для преподавателей)
// @Description Изменяет правильный ответ теста
// @Tags teacher
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param test_id path int true "ID теста"
// @Param input body models.Test true "Новый ответ"
// @Success 200
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Router /teacher/tests/{test_id}/answer [put]
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

// @Summary Удалить тест (Только для преподавателей)
// @Description Удаляет тест по его идентификатору
// @Tags teacher
// @Security ApiKeyAuth
// @Produce json
// @Param test_id path int true "ID теста"
// @Success 200 {string} string "ok"
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Router /teacher/tests/{test_id} [delete]
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

// @Summary Отправить ответ на тест
// @Description Проверяет ответ пользователя и сохраняет результат
// @Tags tests
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param test_id path int true "ID теста"
// @Param input body UserAnswerRequest true "Ответ пользователя"
// @Success 200 {object} models.TestResult
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /tests/{test_id}/submit [post]
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

// @Summary Получить результат теста
// @Description Возвращает результат конкретного теста для текущего пользователя
// @Tags results
// @Security ApiKeyAuth
// @Produce json
// @Param test_result_id path int true "ID результата теста"
// @Success 200 {integer} int "Баллы"
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 401 {object} pkg.ErrorResponse
// @Router /results/{test_result_id} [get]
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

// @Summary Обновить результат теста (Только для преподавателей)
// @Description Изменяет результат теста
// @Tags teacher
// @Security ApiKeyAuth
// @Produce json
// @Param test_result_id path int true "ID результата теста"
// @Success 200 {string} string "ok"
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Router /teacher/results/{test_result_id} [put]
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

// @Summary Удалить результат теста (Только для преподавателей)
// @Description Удаляет результат теста
// @Tags teacher
// @Security ApiKeyAuth
// @Produce json
// @Param test_result_id path int true "ID результата теста"
// @Success 200 {string} string "ok"
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 403 {object} pkg.ErrorResponse
// @Router /teacher/results/{test_result_id} [delete]
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
