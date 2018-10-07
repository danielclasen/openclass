package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openclass/model"
	"openclass/service"
	"strconv"
)

type CourseController struct {
	ApiRouter      gin.IRouter
	CourseService  *service.CourseService
	SessionService *service.SessionService
}

func (controller *CourseController) Routes() *gin.RouterGroup {
	api := controller.ApiRouter.Group("/courses")
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
		c.JSON(http.StatusOK, controller.CourseService.GetCourses())
	}
}

func (controller *CourseController) getByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		course, e := controller.CourseService.GetCourse(id)

		if !handleError(c, e, http.StatusNotFound) {
			c.JSON(http.StatusOK, course)
		}
	}
}

func (controller *CourseController) postHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		newCourse := model.Course{}
		c.ShouldBind(&newCourse)
		course, e := controller.CourseService.CreateCourse(newCourse)
		if !handleError(c, e, http.StatusBadRequest) {
			c.JSON(http.StatusCreated, course)
		}
	}
}

func (controller *CourseController) getSessionsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseId, _ := strconv.Atoi(c.Param("id"))

		if _, err := controller.CourseService.GetCourse(courseId); err != nil {
			handleError(c, err, http.StatusNotFound)
		}

		sessions, err := controller.SessionService.GetSessionsForCourseId(courseId)

		if !handleError(c, err, http.StatusBadRequest) {
			c.JSON(http.StatusOK, sessions)
		}
	}
}
