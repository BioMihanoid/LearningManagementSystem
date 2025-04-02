package service

import (
	models2 "github.com/BioMihanoid/LearningManagementSystem/internal/models"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
)

type Authorization interface {
	CreateUser(user models2.User) (int, error)
	GetUser(name string, password string) (models2.User, error)
}

type UserService interface {
	GetAllUsers() ([]models2.User, error)
	GetUserById(id int) (models2.User, error)
	ChangeUserRole(id int, roleID int) error
	UpdateFirstName(id int, name string) error
	UpdateLastName(id int, name string) error
	UpdateEmail(id int, email string) error
	UpdatePassword(id int, password string, replyPassword string) error
	DeleteUser(id int) error
}

type RoleService interface {
	GetLevelAccess(roleID int) (int, error)
}

type CourseService interface {
	CreateCourse(course models2.Course) error
	UpdateTitleCourse(id int, title string) error
	UpdateDescriptionCourse(id int, description string) error
	GetCourseByID(id int) (models2.Course, error)
	GetAllCourses() ([]models2.Course, error)
	DeleteCourseByID(id int) error
}

type MaterialService interface {
	CreateMaterial(material models2.Material) error
	GetMaterialByID(materialID int) (models2.Material, error)
	GetCourseMaterial(courseID int) ([]models2.Material, error)
	UpdateTitleMaterial(materialID int, title string) error
	UpdateContentMaterial(materialID int, content string) error
	DeleteMaterial(materialID int) error
}

type EnrollmentService interface {
	SubscribeOnCourse(userID int, courseID int) error
	UnsubscribeOnCourse(userID int, courseID int) error
	GetAllUserSubscribedToTheCourse(courseID int) ([]models2.Enrollment, error)
	GetAllCoursesCurrentUser(userID int) ([]models2.Enrollment, error)
}

type TestService interface {
	CreateTest(test models2.Test) error
	GetTestByID(testID int) (models2.Test, error)
	GetAllTestsCourse(courseID int) ([]models2.Test, error)
	GetAllTests() ([]models2.Test, error)
	UpdateQuestionTest(testID int, question string) error
	UpdateAnswerTest(testID int, answer string) error
	DeleteTest(testID int) error
	SubmitTest(testAnswer string, userAnswer string) int
}

type TestResultService interface {
	CreateTestResult(userID int, testID int, score int) error
	GetTestResult(userID int, testResultID int) (models2.TestResult, error)
	UpdateTestResult(testResultID int, score int) error
	DeleteTestResult(testResultID int) error
}

type Service struct {
	Authorization
	UserService
	RoleService
	CourseService
	MaterialService
	EnrollmentService
	TestService
	TestResultService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:     NewAuth(repos),
		UserService:       NewUser(repos),
		RoleService:       NewRole(repos),
		CourseService:     NewCourse(repos),
		MaterialService:   NewMaterial(repos),
		EnrollmentService: NewEnrollment(repos),
		TestService:       NewTest(repos),
		TestResultService: NewTestResult(repos),
	}
}
