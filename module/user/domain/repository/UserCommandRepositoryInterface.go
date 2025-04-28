package repository

import (
	"celeste/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryInterface holds the implementable methods for user command repository
type UserCommandRepositoryInterface interface {
	// DeactivateUser deactivates user
	DeactivateUser(data types.DeactivateUser) error
	// InsertUser inserts a new user
	InsertUser(data types.CreateUser) error
	// UpdateUser updates user
	UpdateUser(data types.UpdateUser) error
	// UpdateUserEmailVerifiedAt updates user email verified at
	UpdateUserEmailVerifiedAt(email string) error
	// UpdateUserPassword updates user password
	UpdateUserPassword(data types.UpdateUserPassword) error
}
