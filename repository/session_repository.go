package repository

import (
	"errors"
	"github.com/danielclasen/openclass/model"
	"time"
)

// SessionRepository is a layer between the services and the actual underlying data source. In that particular case the
// data source is hardcoded and kept in memory without any persistence.
type SessionRepository struct {
	Sessions []model.Session
}

// NewSessionRepository creates a new repository instance and initializes the data source with some hardcoded test values.
func NewSessionRepository() SessionRepository {
	repo := SessionRepository{}

	time1, _ := time.Parse(time.RFC3339, "2018-10-15T14:00:00Z")
	time2, _ := time.Parse(time.RFC3339, "2018-10-16T14:00:00Z")

	repo.Sessions = []model.Session{
		{Id: 1, CourseId: 1, Date: time1, Location: "Bonn", MaxStudents: 10},
		{Id: 2, CourseId: 1, Date: time2, Location: "Köln", MaxStudents: 10},
	}
	return repo
}

// FindById queries the data source to find the one Session with the given id. It returns the found Session instance or
// nil,error if not found.
func (s *SessionRepository) FindById(id int) (*model.Session, error) {
	for _, v := range s.Sessions {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errors.New("session not found")
}

// FindAllByCourseId queries the data source for all known Session instances matching the given courseId.
// It returns a slice of pointers to the found Session instance.
func (s *SessionRepository) FindAllByCourseId(courseId int) []*model.Session {
	var filtered []*model.Session
	for i := 0; i < len(s.Sessions); i++ {
		if s.Sessions[i].CourseId == courseId {
			filtered = append(filtered, &s.Sessions[i])
		}
	}
	return filtered
}

// Store saves the given Session instance in the data source. It returns the freshly saved Session instance or nil,error
// if the given id is not unique.
func (s *SessionRepository) Store(session model.Session) (*model.Session, error) {
	if data, e := s.FindById(session.Id); data == nil && e != nil {
		s.Sessions = append(s.Sessions, session)
		return &session, nil
	} else {
		return nil, errors.New("duplicate session id")
	}
}
