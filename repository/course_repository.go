package repository

import (
	"errors"
	"github.com/danielclasen/openclass/model"
)

type CourseRepository interface {
	FindAll() *[]model.Course
	FindById(id int) (*model.Course, error)
	Store(course model.Course) (*model.Course, error)
}

type courseStorage struct {
	Courses []model.Course
}

func NewCourseRepository() CourseRepository {
	repo := new(courseStorage)
	repo.Courses = []model.Course{
		model.Course{Id: 1, Title: "Java for beginner", Description: "A Java beginner guide.", Lecturer: "Mr. Foo", Price: 9000},
		model.Course{Id: 2, Title: "Golang for experts", Description: "A deep dive into the world of go.", Lecturer: "Mr. Bar", Price: 18000},
	}
	return repo
}

func (s *courseStorage) FindAll() *[]model.Course {
	return &s.Courses
}

func (s *courseStorage) FindById(id int) (*model.Course, error) {
	for _, v := range s.Courses {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errors.New("course not found")
}

func (s *courseStorage) Store(course model.Course) (*model.Course, error) {
	if data, e := s.FindById(course.Id); data == nil && e != nil {
		s.Courses = append(s.Courses, course)
		return &course, nil
	} else {
		return nil, errors.New("duplicate course id")
	}
}
