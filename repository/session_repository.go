package repository

import (
	"errors"
	"openclass/model"
	"time"
)

type SessionRepository interface {
	FindById(id int) (*model.Session, error)
	FindAllByCourseId(courseId int) *[]model.Session
	Store(session model.Session) (*model.Session, error)
}

type sessionStorage struct {
	Sessions []model.Session
}

func NewSessionRepository() SessionRepository {
	repo := new(sessionStorage)

	time1, _ := time.Parse(time.RFC3339, "2018-10-15T14:00:00Z")
	time2, _ := time.Parse(time.RFC3339, "2018-10-16T14:00:00Z")

	repo.Sessions = []model.Session{
		model.Session{Id: 1, CourseId: 1, Date: time1, Location: "Bonn", MaxStudents: 10},
		model.Session{Id: 2, CourseId: 1, Date: time2, Location: "Köln", MaxStudents: 10},
	}
	return repo
}

func (s *sessionStorage) FindById(id int) (*model.Session, error) {
	for _, v := range s.Sessions {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errors.New("session not found")
}

func (s *sessionStorage) FindAllByCourseId(courseId int) *[]model.Session {

	filteredSessions := make([]model.Session, 0)
	for _, session := range s.Sessions {
		if session.CourseId == courseId {
			filteredSessions = append(filteredSessions, session)
		}
	}

	return &filteredSessions
}

func (s *sessionStorage) Store(session model.Session) (*model.Session, error) {
	if data, e := s.FindById(session.Id); data == nil && e != nil {
		s.Sessions = append(s.Sessions, session)
		return &session, nil
	} else {
		return nil, errors.New("duplicate session id")
	}
}
