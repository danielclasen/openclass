package model

type Course struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Lecturer    string `json:"lecturer" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
}
