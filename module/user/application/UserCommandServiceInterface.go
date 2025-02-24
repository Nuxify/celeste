package application

import (
	"context"

	"celeste/module/user/infrastructure/service/types"
)

// UserCommandServiceInterface holds the implementable methods for the user command service
type UserCommandServiceInterface interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, data types.CreateUser) (string, error)
	// UpdateUser updates user
	UpdateUser(ctx context.Context, data types.UpdateUser) error
	// UpdateUserPassword updates user password
	UpdateUserPassword(ctx context.Context, data types.UpdateUserPassword) error
}
