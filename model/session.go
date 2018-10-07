package model

import (
	"time"
)

type Session struct {
	Id          int       `json:"id"`
	CourseId    int       `json:"courseId" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	MaxStudents int       `json:"maxStudents" binding:"required"`
}
