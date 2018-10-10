package controller

import (
	"github.com/danielclasen/openclass/model"
	"github.com/danielclasen/openclass/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CourseController holds all the required dependencies to act as a controller facade or delivery layer for the Course
// model and its immediate children/siblings.
type CourseController struct {
	apiRouter      *gin.RouterGroup
	courseService  *service.CourseService
	sessionService *service.SessionService
}

// NewCourseController creates a new controller instance with the given dependencies.
func NewCourseController(apiRouter *gin.RouterGroup, courseService *service.CourseService, sessionService *service.SessionService) CourseController {
	return CourseController{apiRouter: apiRouter, courseService: courseService, sessionService: sessionService}
}

// Routes registers this controllers sub-routing in the main apiRouter. It returns a RouterGroup containing only the
// routes for the operations on the Course model.
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
			if len(sessions) == 0 {
				c.JSON(http.StatusOK, []model.Session{})
			} else {
				c.JSON(http.StatusOK, sessions)
			}
		}
	}
}
