package service

import (
	"github.com/danielclasen/openclass/model"
	"testing"
)

type MockRepo struct {
}

func (r *MockRepo) FindAllBySessionId(sessionId int) []*model.Participation {
	return nil
}

func (r *MockRepo) Store(participation model.Participation) (*model.Participation, error) {
	return &participation, nil
}

type MockService struct {
}

func (s *MockService) GetSessionsForCourseId(courseId int) ([]*model.Session, error) {
	return nil, nil
}
func (s *MockService) GetSession(sessionId int) (*model.Session, error) {
	session := model.Session{}
	session.MaxStudents = 1
	session.Id = sessionId
	return &session, nil
}
func (s *MockService) CreateSession(session model.Session) (*model.Session, error) {
	return nil, nil
}

func TestParticipationService_CreateParticipationForSessionId(t *testing.T) {
	//given
	sut := ParticipationService{}
	sut.participationRepository = &MockRepo{}
	sut.sessionService = &MockService{}

	person := model.Person{}
	//when

	part, _ := sut.CreateParticipationForSessionId(1, person)

	//then
	if part.Person != person || part.SessionId != 1 {
		t.Errorf("err")
	}
}
