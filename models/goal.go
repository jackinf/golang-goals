package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

// Goal represents an goal record.
type Goal struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Due         time.Time `json:"due" db:"due"`
	Motivation  string    `json:"motivation" db:"motivation"`
}

// Validate validates the Goal fields.
func (m Goal) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 120)),
	)
}
