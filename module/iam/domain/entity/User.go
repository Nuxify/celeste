package entity

import (
	"time"
)

// User holds the user entity fields
type User struct {
	ID        string
	Data      string
	CreatedAt time.Time `db:"created_at"`
}

// GetModelName returns the model name of user entity that can be used for naming schemas
func (entity *User) GetModelName() string {
	return "users"
}
