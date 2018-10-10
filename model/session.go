package model

import (
	"time"
)

// Session represents the actual happening of a course on a given date on a given location.
type Session struct {
	Id          int       `json:"id"`
	CourseId    int       `json:"courseId" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	MaxStudents int       `json:"maxStudents" binding:"required"`
}
