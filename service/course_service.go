package service

import (
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/repository"
)

type CourseService struct {
	CourseRepository repository.CourseRepository
}

func (service *CourseService) GetCourses() *[]model.Course {
	return service.CourseRepository.FindAll()
}

func (service *CourseService) GetCourse(id int) (*model.Course, error) {
	return service.CourseRepository.FindById(id)
}

func (service *CourseService) CreateCourse(course model.Course) (*model.Course, error) {
	return service.CourseRepository.Store(course)
}
