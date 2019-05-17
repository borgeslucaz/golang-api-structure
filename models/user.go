package models

import (
	"github.com/go-pg/pg/types"
	"github.com/borgeslucaz/golang-api-structure/utils"
	"time"

)

// User struct 
type User struct {
	Base                  // Import default fields ID, Created, Updated e Deleted
	Email       string         `json:"email"        db:"email"`
	Password    string         `json:"-"     db:"password,omitempty"`
	ValidatedAt types.NullTime `json:"validated_at" db:"validated_at"`
	LastLogin   types.NullTime `json:"last_login"   db:"last_login"`
	IsActive    bool           `json:"is_active"    db:"is_active"`
}

// NewUser creates a instance of user with hashed password
func NewUser(email string, password string) (*User, error) {
	var err error
	u := new(User)
	var hash utils.Hash
	u.Password, err = hash.Generate(password)
	if err != nil {
		return u, err
	}

	now := time.Now()
	u.Email = email
	u.CreatedAt = now
	u.UpdatedAt = now
	u.IsActive = false
	return u, nil
}