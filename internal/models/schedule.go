package models

import (
	uuid "github.com/google/uuid"
)

type Schedule struct {
	Id          uuid.UUID   `json:"id" db:"id"`
	Description string      `json:"description" db:"description"`
	Users       []uuid.UUID `json:"users" db:"users" pg:"array"`
}
