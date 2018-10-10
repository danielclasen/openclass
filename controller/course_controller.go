package controller

import (
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CourseController struct {
	apiRouter      *gin.RouterGroup
	courseService  *service.CourseService
	sessionService *service.SessionService
}

func NewCourseController(apiRouter *gin.RouterGroup, courseService *service.CourseService, sessionService *service.SessionService) CourseController {
	return CourseController{apiRouter: apiRouter, courseService: courseService, sessionService: sessionService}
}

func (controller *CourseController) Routes() *gin.RouterGroup {
	api := controller.apiRouter.Group("/courses")
	{
		api.GET("/", controller.getListHandler())
		api.GET("/:id", controller.getByIdHandler())
		api.GET("/:id/sessions", controller.getSessionsHandler())
		api.POST("/", controller.postHandler())
	}
	return api
}

func (controller *CourseController) getListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, controller.courseService.GetCourses())
	}
}

func (controller *CourseController) getByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		course, e := controller.courseService.GetCourse(id)

		if !handleError(c, e, http.StatusNotFound) {
			c.JSON(http.StatusOK, course)
		}
	}
}

func (controller *CourseController) postHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		newCourse := model.Course{}
		c.ShouldBind(&newCourse)
		course, e := controller.courseService.CreateCourse(newCourse)
		if !handleError(c, e, http.StatusBadRequest) {
			c.JSON(http.StatusCreated, course)
		}
	}
}

func (controller *CourseController) getSessionsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseId, _ := strconv.Atoi(c.Param("id"))

		if _, err := controller.courseService.GetCourse(courseId); err != nil {
			handleError(c, err, http.StatusNotFound)
		}

		sessions, err := controller.sessionService.GetSessionsForCourseId(courseId)

		if !handleError(c, err, http.StatusBadRequest) {
			c.JSON(http.StatusOK, sessions)
		}
	}
}
