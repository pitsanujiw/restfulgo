package models

// Person is a representation of a person
type Users struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}
