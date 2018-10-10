package repository

import (
	"errors"
	"github.com/danielclasen/openclass/model"
)

// ParticipationRepository is a layer between the services and the actual underlying data source. In that particular case the
// data source is hardcoded and kept in memory without any persistence.
type ParticipationRepository struct {
	Participations []model.Participation
}

// NewParticipationRepository creates a new repository instance and initializes the data source with some hardcoded test values.
func NewParticipationRepository() ParticipationRepository {
	repo := ParticipationRepository{}

	repo.Participations = []model.Participation{}

	return repo
}

// FindAllBySessionId queries the data source for all known Participation instances matching the given sessionId.
// It returns a slice of pointers to the found Participation instance.
func (s *ParticipationRepository) FindAllBySessionId(sessionId int) []*model.Participation {

	var filtered []*model.Participation

	for i := 0; i < len(s.Participations); i++ {
		filtered = append(filtered, &s.Participations[i])
	}

	return filtered
}

// Store saves the given Participation instance in the data source. It returns the freshly saved Participation instance
// or nil,error if the given combination of session and Person.Email is not unique.
func (s *ParticipationRepository) Store(participation model.Participation) (*model.Participation, error) {

	for _, v := range s.Participations {
		if v.SessionId == participation.SessionId && v.Person.Email == participation.Person.Email {
			return nil, errors.New("already participating")
		}
	}

	s.Participations = append(s.Participations, participation)

	return &participation, nil

}
