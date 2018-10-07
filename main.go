package main

import (
	"github.com/danielclasen/openclass/controller"
	"github.com/danielclasen/openclass/repository"
	"github.com/danielclasen/openclass/service"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	engine, api := getMainEngine()
	setupRouting(api)

	// Start and run the server
	engine.Run(":3000")
}

func getMainEngine() (engine *gin.Engine, api *gin.RouterGroup) {
	engine = gin.Default()
	engine.Use(static.Serve("/", static.LocalFile("./web/dist", true)))
	api = engine.Group("/api/v1")
	return
}

func setupRouting(api *gin.RouterGroup) {

	//repo initialization
	coursesRepository := repository.NewCourseRepository()

	sessionRepository := repository.NewSessionRepository()

	participationRepository := repository.NewParticipationRepository()

	//service initialization
	courseService := new(service.CourseService)
	courseService.CourseRepository = coursesRepository

	sessionService := new(service.SessionService)
	sessionService.SessionRepository = sessionRepository
	sessionService.CourseService = courseService

	participationService := new(service.ParticipationService)
	participationService.ParticipationRepository = participationRepository

	//controller initialization
	courseController := new(controller.CourseController)
	courseController.ApiRouter = api
	courseController.CourseService = courseService
	courseController.SessionService = sessionService
	courseController.Routes()

	sessionController := new(controller.SessionController)
	sessionController.ApiRouter = api
	sessionController.SessionService = sessionService
	sessionController.ParticipationService = participationService
	sessionController.Routes()

}
