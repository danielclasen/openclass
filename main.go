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
	courseService := service.NewCourseService(&coursesRepository)
	sessionService := service.NewSessionService(&sessionRepository, &courseService)
	participationService := service.NewParticipationService(&participationRepository, &sessionService)

	//controller initialization
	courseController := controller.NewCourseController(api, &courseService, &sessionService)
	courseController.Routes()

	sessionController := controller.NewSessionController(api, &sessionService, &participationService)
	sessionController.Routes()

}
