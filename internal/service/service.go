package service

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name string, password string) (models.User, error)
}

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
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
	CreateCourse(course models.Course) error
	UpdateTitleCourse(id int, title string) error
	UpdateDescriptionCourse(id int, description string) error
	GetCourseByID(id int) (models.Course, error)
	GetAllCourses() ([]models.Course, error)
	DeleteCourseByID(id int) error
}

type MaterialService interface {
	CreateMaterial(material models.Material) error
	GetMaterialByID(materialID int) (models.Material, error)
	GetCourseMaterial(courseID int) ([]models.Material, error)
	UpdateTitleMaterial(materialID int, title string) error
	UpdateContentMaterial(materialID int, content string) error
	DeleteMaterial(materialID int) error
}

type EnrollmentService interface {
	SubscribeOnCourse(userID int, courseID int) error
	UnsubscribeOnCourse(userID int, courseID int) error
	GetAllUserSubscribedToTheCourse(courseID int) ([]models.Enrollment, error)
	GetAllCoursesCurrentUser(userID int) ([]models.Enrollment, error)
}

type Service struct {
	Authorization
	UserService
	RoleService
	CourseService
	MaterialService
	EnrollmentService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:     NewAuth(repos),
		UserService:       NewUser(repos),
		RoleService:       NewRole(repos),
		CourseService:     NewCourse(repos),
		MaterialService:   NewMaterial(repos),
		EnrollmentService: NewEnrollment(repos),
	}
}
