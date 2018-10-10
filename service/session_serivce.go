package service

import (
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/repository"
)

// SessionService is a layer between the controller facade and the repository. It holds all necessary dependencies needed
// to perform the required business logic and/or data mutation applied to the Session type.
type SessionService struct {
	sessionRepository *repository.SessionRepository
	courseService     *CourseService
}

// NewSessionService returns a new service instance
func NewSessionService(sessionRepository *repository.SessionRepository, courseService *CourseService) SessionService {
	return SessionService{sessionRepository: sessionRepository, courseService: courseService}
}

// GetSessionsForCourseId fetches all Sessions for the given courseId. It returns a slice of pointers to the actual Session
// instances and any error occurred.
func (service *SessionService) GetSessionsForCourseId(courseId int) (sessions []*model.Session, err error) {
	return service.sessionRepository.FindAllByCourseId(courseId), nil
}

// GetSession fetches the one Session with the given sessionId from the repository. It returns a pointer to the found Session
// and any error occurred.
func (service *SessionService) GetSession(sessionId int) (sessions *model.Session, err error) {
	return service.sessionRepository.FindById(sessionId)
}

// CreateSession pushes a new Session instance to the repository. It returns the freshly created Session instance and
// any error occurred.
func (service *SessionService) CreateSession(session model.Session) (*model.Session, error) {
	if _, err := service.courseService.GetCourse(session.CourseId); err != nil {
		return nil, err
	}
	return service.sessionRepository.Store(session)
}
