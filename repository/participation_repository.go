package repository

import (
	"errors"
	"openclass/model"
)

type ParticipationRepository interface {
	FindAllBySessionId(sessionId int) *[]model.Participation
	Store(participation model.Participation) (*model.Participation, error)
}

type participationStorage struct {
	Participations []model.Participation
}

func NewParticipationRepository() ParticipationRepository {
	repo := new(participationStorage)

	repo.Participations = []model.Participation{}

	return repo
}

func (s *participationStorage) FindAllBySessionId(sessionId int) *[]model.Participation {

	filteredParticipations := make([]model.Participation, 0)
	for _, participation := range s.Participations {
		if participation.SessionId == sessionId {
			filteredParticipations = append(filteredParticipations, participation)
		}
	}

	return &filteredParticipations
}

func (s *participationStorage) Store(participation model.Participation) (*model.Participation, error) {

	for _, v := range s.Participations {
		if v.SessionId == participation.SessionId && v.Person.Email == participation.Person.Email {
			return nil, errors.New("already participating")
		}
	}

	s.Participations = append(s.Participations, participation)

	return &participation, nil

}
