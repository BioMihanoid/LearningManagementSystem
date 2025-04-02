package repository

import (
	"database/sql"
	models2 "github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository/postgres"
)

type Authorization interface {
	CreateUser(user models2.User) (int, error)
	GetUser(email string, passwordHash string) (models2.User, error)
	GetUserID(email string) (int, error)
}

type UserRepository interface {
	GetAllUsers() ([]models2.User, error)
	GetUserByID(id int) (models2.User, error)
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
	CreateCourse(course models2.Course) error
	UpdateTitleCourse(id int, title string) error
	UpdateDescriptionCourse(id int, description string) error
	GetCourseByID(id int) (models2.Course, error)
	GetAllCourses() ([]models2.Course, error)
	DeleteCourseByID(id int) error
}

type EnrollmentRepository interface {
	SubscribeOnCourse(userID int, courseID int) error
	UnsubscribeOnCourse(userID int, courseID int) error
	GetAllUserSubscribedToTheCourse(courseID int) ([]models2.Enrollment, error)
	GetAllCoursesCurrentUser(userID int) ([]models2.Enrollment, error)
}

type MaterialRepository interface {
	CreateMaterial(material models2.Material) error
	GetMaterialByID(materialID int) (models2.Material, error)
	GetCourseMaterial(courseID int) ([]models2.Material, error)
	UpdateTitleMaterial(materialID int, title string) error
	UpdateContentMaterial(materialID int, content string) error
	DeleteMaterial(materialID int) error
}

type TestRepository interface {
	CreateTest(courseID int, question string, answer string) error
	GetTestByID(testID int) (models2.Test, error)
	GetAllTestsCourse(courseID int) ([]models2.Test, error)
	GetAllTests() ([]models2.Test, error)
	UpdateQuestionTest(testID int, question string) error
	UpdateAnswerTest(testID int, answer string) error
	DeleteTest(testID int) error
}

type TestResultRepository interface {
	CreateTestResult(userID int, testID int, score int) error
	GetTestResult(testResultID int) (models2.TestResult, error)
	UpdateTestResult(testResultID int, score int) error
	DeleteTestResult(testResultID int) error
}

type LogRepository interface {
	CreateLog(userID int, action string) error
	GetLogByID(logID int) (models2.Log, error)
	GetLogsCurrentUser(userID int) ([]models2.Log, error)
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
