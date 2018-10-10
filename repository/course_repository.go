package repository

import (
	"errors"
	"github.com/danielclasen/openclass/model"
)

// CourseRepository is a layer between the services and the actual underlying data source. In that particular case the
// data source is hardcoded and kept in memory without any persistence.
type CourseRepository struct {
	Courses []model.Course
}

// NewCourseRepository creates a new repository instance and initializes the data source with some hardcoded test values.
func NewCourseRepository() CourseRepository {
	repo := CourseRepository{}
	repo.Courses = []model.Course{
		{Id: 1, Title: "Java for beginner", Description: "A Java beginner guide.", Lecturer: "Mr. Foo", Price: 9000},
		{Id: 2, Title: "Golang for experts", Description: "A deep dive into the world of go.", Lecturer: "Mr. Bar", Price: 18000},
	}
	return repo
}

// FindAll queries the data source for all known Course instances. It returns a slice of pointers to the found Course instance.
func (s *CourseRepository) FindAll() []*model.Course {
	var result []*model.Course

	for i := 0; i < len(s.Courses); i++ {
		result = append(result, &s.Courses[i])
	}

	return result
}

// FindById queries the data source to find the one Course with the given id. It returns the found Course instance or
// nil,error if not found.
func (s *CourseRepository) FindById(id int) (*model.Course, error) {
	for _, v := range s.Courses {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errors.New("course not found")
}

// Store saves the given Course instance in the data source. It returns the freshly saved Course instance or nil,error if the
// given id is not unique.
func (s *CourseRepository) Store(course model.Course) (*model.Course, error) {
	if data, e := s.FindById(course.Id); data == nil && e != nil {
		s.Courses = append(s.Courses, course)
		return &course, nil
	} else {
		return nil, errors.New("duplicate course id")
	}
}
