package repository

import (
	"database/sql"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository/postgres"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email string, passwordHash string) (models.User, error)
	GetUserID(email string) (int, error)
}

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	ChangeUserRole(id int, roleId int) error
	UpdateFirstName(id int, name string) error
	UpdateLastName(id int, name string) error
	UpdateEmail(id int, email string) error
	ChangePassword(id int, password string) error
	DeleteUser(id int) error
}

type RoleRepository interface {
	GetLevelAccess(id int) (int, error)
}

type CourseRepository interface {
	CreateCourse(course models.Course) error
	UpdateTitleCourse(id int, title string) error
	UpdateDescriptionCourse(id int, description string) error
	GetCourseByID(id int) (models.Course, error)
	GetAllCourses() ([]models.Course, error)
	DeleteCourseByID(id int) error
}

type EnrollmentRepository interface {
	SubscribeOnCourse(userID int, courseID int) error
	UnsubscribeOnCourse(userID int, courseID int) error
	GetAllUserSubscribedToTheCourse(courseID int) ([]models.Enrollment, error)
	GetAllCoursesCurrentUser(userID int) ([]models.Enrollment, error)
}

type MaterialRepository interface {
	CreateMaterial(material models.Material) error
	GetMaterialByID(materialID int) (models.Material, error)
	GetCourseMaterial(courseID int) ([]models.Material, error)
	UpdateTitleMaterial(materialID int, title string) error
	UpdateContentMaterial(materialID int, content string) error
	DeleteMaterial(materialID int) error
}

type TestRepository interface {
	CreateTest(courseID int, question string, answer string) error
	GetTestByID(testID int) (models.Test, error)
	GetAllTestsByCourseID(courseID int) ([]models.Test, error)
	GetAllTests() ([]models.Test, error)
	UpdateQuestionTest(testID int, question string) error
	UpdateAnswerTest(testID int, answer string) error
	DeleteTest(testID int) error
}

type TestResultRepository interface {
	CreateTestResult(userID int, testID int, score int) error
	GetTestResult(testResultID int) (models.TestResult, error)
	UpdateTestResult(testResultID int, score int) error
	DeleteTestResult(testResultID int) error
}

type LogRepository interface {
	CreateLog(userID int, action string) error
	GetLogByID(logID int) (models.Log, error)
	GetLogsCurrentUser(userID int) ([]models.Log, error)
	DeleteLogByID(logID int) error
}

type Repository struct {
	Authorization
	UserRepository
	RoleRepository
	CourseRepository
	EnrollmentRepository
	MaterialRepository
	TestRepository
	TestResultRepository
	LogRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:        postgres.NewAuth(db),
		UserRepository:       postgres.NewUser(db),
		RoleRepository:       postgres.NewRole(db),
		CourseRepository:     postgres.NewCourse(db),
		EnrollmentRepository: postgres.NewEnrollment(db),
		MaterialRepository:   postgres.NewMaterial(db),
		TestRepository:       postgres.NewTest(db),
		TestResultRepository: postgres.NewTestResult(db),
		LogRepository:        postgres.NewLog(db),
	}
}
