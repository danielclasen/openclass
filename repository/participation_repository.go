package repository

import (
	"errors"
	"github.com/danielclasen/openclass/model"
)

type ParticipationRepository struct {
	Participations []model.Participation
}

func NewParticipationRepository() ParticipationRepository {
	repo := ParticipationRepository{}

	repo.Participations = []model.Participation{}

	return repo
}

func (s *ParticipationRepository) FindAllBySessionId(sessionId int) []*model.Participation {

	var filtered []*model.Participation

	for i := 0; i < len(s.Participations); i++ {
		filtered = append(filtered, &s.Participations[i])
	}

	return filtered
}

func (s *ParticipationRepository) Store(participation model.Participation) (*model.Participation, error) {

	for _, v := range s.Participations {
		if v.SessionId == participation.SessionId && v.Person.Email == participation.Person.Email {
			return nil, errors.New("already participating")
		}
	}

	s.Participations = append(s.Participations, participation)

	return &participation, nil

}
