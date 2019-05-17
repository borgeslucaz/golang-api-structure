package models

import (
	"time"
)

type Base struct {
	ID        int       `json:"id" sql:"id,pk"`
	CreatedAt time.Time  `json:"created_at" sql:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" sql:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"deleted_at"`
}
