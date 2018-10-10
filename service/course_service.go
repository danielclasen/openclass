package service

import (
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/repository"
)

// CourseService is a layer between the controller facade and the repository. It holds all necessary dependencies needed
// to perform the required business logic and/or data mutation applied to the Course type.
type CourseService struct {
	courseRepository *repository.CourseRepository
}

// NewCourseService returns a new service instance
func NewCourseService(repository *repository.CourseRepository) CourseService {
	return CourseService{courseRepository: repository}
}

// GetCourses fetches all known Courses from the repository. It returns a slice of pointers to the actual Course instances.
func (service *CourseService) GetCourses() []*model.Course {
	return service.courseRepository.FindAll()
}

// GetCourse fetches the one Course with the matching id from repository. It returns a pointer to the found Course and any error occurred.
func (service *CourseService) GetCourse(id int) (*model.Course, error) {
	return service.courseRepository.FindById(id)
}

// CreateCourse pushes an instace of a Course to the repository to store it. It returns a pointer to the newly created
// Course instance and any error occurred.
func (service *CourseService) CreateCourse(course model.Course) (*model.Course, error) {
	return service.courseRepository.Store(course)
}
