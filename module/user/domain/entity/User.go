package entity

import (
	"time"
)

// User holds the user entity fields
type User struct {
	WalletAddress   string `db:"wallet_address"`
	Email           string
	Password        string
	SSS1            string `db:"sss_1"`
	Name            string
	EmailVerifiedAt *time.Time `db:"email_verified_at"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
}

// GetModelName returns the model name of user entity that can be used for naming schemas
func (entity *User) GetModelName() string {
	return "users"
}
