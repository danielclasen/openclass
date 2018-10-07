package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"openclass/controller"
	"openclass/repository"
	"openclass/service"
)

func main() {

	engine, api := GetMainEngine()
	SetupRouting(api)

	// Start and run the server
	engine.Run(":3000")
}

func GetMainEngine() (engine *gin.Engine, api *gin.RouterGroup) {
	engine = gin.Default()
	engine.Use(static.Serve("/", static.LocalFile("./web", true)))
	api = engine.Group("/api/v1")
	return
}

func SetupRouting(api *gin.RouterGroup) {

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
