package repository

import (
	"celeste/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryInterface holds the implementable methods for user command repository
type UserCommandRepositoryInterface interface {
	// InsertUser inserts a new user
	InsertUser(data types.CreateUser) error
	// UpdateUserPassword updates user password
	UpdateUserPassword(data types.UpdateUserPassword) error
	// UpdateUser updates user
	UpdateUser(data types.UpdateUser) error
}
