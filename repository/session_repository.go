package repository

import (
	"errors"
	"github.com/danielclasen/openclass/model"
	"time"
)

type SessionRepository struct {
	Sessions []model.Session
}

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

func (s *SessionRepository) FindById(id int) (*model.Session, error) {
	for _, v := range s.Sessions {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errors.New("session not found")
}

func (s *SessionRepository) FindAllByCourseId(courseId int) []*model.Session {
	var filtered []*model.Session
	for i := 0; i < len(s.Sessions); i++ {
		filtered = append(filtered, &s.Sessions[i])
	}
	return filtered
}

func (s *SessionRepository) Store(session model.Session) (*model.Session, error) {
	if data, e := s.FindById(session.Id); data == nil && e != nil {
		s.Sessions = append(s.Sessions, session)
		return &session, nil
	} else {
		return nil, errors.New("duplicate session id")
	}
}
