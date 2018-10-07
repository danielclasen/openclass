package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openclass/model"
	"openclass/service"
	"strconv"
)

type SessionController struct {
	ApiRouter            gin.IRouter
	SessionService       *service.SessionService
	ParticipationService *service.ParticipationService
}

func (controller *SessionController) Routes() *gin.RouterGroup {
	api := controller.ApiRouter.Group("/sessions")
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
		session, e := controller.SessionService.GetSession(id)

		if !handleError(c, e, http.StatusNotFound) {
			c.JSON(http.StatusOK, session)
		}
	}
}

func (controller *SessionController) postHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		newSession := model.Session{}
		c.ShouldBind(&newSession)

		session, e := controller.SessionService.CreateSession(newSession)
		if !handleError(c, e, http.StatusBadRequest) {
			c.JSON(http.StatusCreated, session)
		}
	}
}

func (controller *SessionController) getParticipationsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, e := controller.SessionService.GetSession(id)
		if !handleError(c, e, http.StatusNotFound) {

			persons := controller.ParticipationService.GetAllParticipatingPersonsForSessionId(id)
			c.JSON(http.StatusOK, persons)

		}
	}
}

func (controller *SessionController) postParticipationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, e := controller.SessionService.GetSession(id)
		if !handleError(c, e, http.StatusNotFound) {

			person := model.Person{}
			c.ShouldBind(&person)

			_, err := controller.ParticipationService.CreateParticipationForSessionId(id, person)

			if !handleError(c, err, http.StatusBadRequest) {
				c.JSON(http.StatusNoContent, nil)
			}

		}
	}
}
