package service

import (
	"openclass/model"
	"openclass/repository"
)

type SessionService struct {
	SessionRepository repository.SessionRepository
	CourseService     *CourseService
}

func (service *SessionService) GetSessionsForCourseId(courseId int) (sessions *[]model.Session, err error) {
	return service.SessionRepository.FindAllByCourseId(courseId), nil
}

func (service *SessionService) GetSession(sessionId int) (sessions *model.Session, err error) {
	return service.SessionRepository.FindById(sessionId)
}

func (service *SessionService) CreateSession(session model.Session) (*model.Session, error) {
	if _, err := service.CourseService.GetCourse(session.CourseId); err != nil {
		return nil, err
	}
	return service.SessionRepository.Store(session)
}
