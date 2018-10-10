package controller

import (
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SessionController holds all the required dependencies to act as a controller facade or delivery layer for the Session
// model and its immediate children/siblings.
type SessionController struct {
	apiRouter            *gin.RouterGroup
	sessionService       *service.SessionService
	participationService *service.ParticipationService
}

// NewSessionController creates a new controller instance with the given dependencies.
func NewSessionController(apiRouter *gin.RouterGroup, sessionService *service.SessionService, participationService *service.ParticipationService) SessionController {
	return SessionController{apiRouter: apiRouter, sessionService: sessionService, participationService: participationService}
}

// Routes registers this controllers sub-routing in the main apiRouter. It returns a RouterGroup containing only the
// routes for the operations on the Session model.
func (controller *SessionController) Routes() *gin.RouterGroup {
	api := controller.apiRouter.Group("/sessions")
	{
		api.GET("/:id", controller.getByIdHandler())
		api.GET("/:id/participations", controller.getParticipationsHandler())
		api.POST("/:id/participations", controller.postParticipationHandler())
		api.POST("/", controller.postHandler())
	}
	return api
}

func (controller *SessionController) getByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		session, e := controller.sessionService.GetSession(id)

		if !handleError(c, e, http.StatusNotFound) {
			c.JSON(http.StatusOK, session)
		}
	}
}

func (controller *SessionController) postHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		newSession := model.Session{}
		c.ShouldBind(&newSession)

		session, e := controller.sessionService.CreateSession(newSession)
		if !handleError(c, e, http.StatusBadRequest) {
			c.JSON(http.StatusCreated, session)
		}
	}
}

func (controller *SessionController) getParticipationsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, e := controller.sessionService.GetSession(id)
		if !handleError(c, e, http.StatusNotFound) {

			persons := controller.participationService.GetAllParticipatingPersonsForSessionId(id)

			if len(persons) == 0 {
				c.JSON(http.StatusOK, []model.Person{})
			} else {
				c.JSON(http.StatusOK, persons)
			}

		}
	}
}

func (controller *SessionController) postParticipationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, e := controller.sessionService.GetSession(id)
		if !handleError(c, e, http.StatusNotFound) {

			person := model.Person{}
			c.ShouldBind(&person)

			_, err := controller.participationService.CreateParticipationForSessionId(id, person)

			if !handleError(c, err, http.StatusBadRequest) {
				c.JSON(http.StatusNoContent, nil)
			}

		}
	}
}
