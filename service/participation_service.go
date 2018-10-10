package service

import (
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/repository"
)

type ParticipationService struct {
	ParticipationRepository repository.ParticipationRepository
}

func (service *ParticipationService) GetAllParticipatingPersonsForSessionId(sessionId int) []*model.Person {
	participations := service.ParticipationRepository.FindAllBySessionId(sessionId)

	var persons []*model.Person

	for i := 0; i < len(participations); i++ {
		persons = append(persons, &participations[i].Person)
	}

	return persons
}

func (service *ParticipationService) CreateParticipationForSessionId(sessionId int, person model.Person) (*model.Participation, error) {
	return service.ParticipationRepository.Store(model.Participation{SessionId: sessionId, Person: person})
}
