package model

// Participation represents the participation of a Person in a Session of a Course.
type Participation struct {
	SessionId int
	Person    Person
}
