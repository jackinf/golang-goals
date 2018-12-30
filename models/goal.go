package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

// Goal represents an goal record.
type Goal struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"name" db:"title"`
	Description string    `json:"name" db:"description"`
	Due         time.Time `json:"name" db:"due"`
	Motivation  string    `json:"name" db:"motivation"`
}

// Validate validates the Goal fields.
func (m Goal) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 120)),
	)
}
