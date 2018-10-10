package service

import (
	"errors"
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/repository"
)

// ParticipationService is a layer between the controller facade and the repository. It holds all necessary dependencies needed
// to perform the required business logic and/or data mutation applied to the Person/Participation type.
type ParticipationService struct {
	participationRepository repository.ParticipationStorage
	sessionService          SessionServiceProvider
}

// NewParticipationService returns a new service instance
func NewParticipationService(participationRepository *repository.ParticipationRepository, sessionService *SessionService) ParticipationService {
	return ParticipationService{participationRepository: participationRepository, sessionService: sessionService}
}

// GetAllParticipatingPersonsForSessionId fetches all Participations for the given sessionId from the repository and maps
// them to return a slice of pointers to the actual Person instances.
func (service *ParticipationService) GetAllParticipatingPersonsForSessionId(sessionId int) []*model.Person {
	participations := service.participationRepository.FindAllBySessionId(sessionId)

	var persons []*model.Person

	for i := 0; i < len(participations); i++ {
		persons = append(persons, &participations[i].Person)
	}

	return persons
}

// CreateParticipationForSessionId creates and pushes a new Participation to the repository. It checks that the maxStudents
// count is not exceeded and returns nil,error otherwise. On success it returns the freshly created Participation instance and any error occurred.
func (service *ParticipationService) CreateParticipationForSessionId(sessionId int, person model.Person) (*model.Participation, error) {

	session, err := service.sessionService.GetSession(sessionId)
	if err != nil {
		return nil, errors.New("session not found")
	}

	participations := service.participationRepository.FindAllBySessionId(sessionId)

	if len(participations) >= session.MaxStudents {
		return nil, errors.New("maxStudents exceeded")
	}

	return service.participationRepository.Store(model.Participation{SessionId: sessionId, Person: person})
}
