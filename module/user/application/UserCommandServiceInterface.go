package application

import (
	"context"

	"celeste/module/user/infrastructure/service/types"
)

// UserCommandServiceInterface holds the implementable methods for the user command service
type UserCommandServiceInterface interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, data types.CreateUser) (types.CreateUserResult, error)
	// DeactivateUser deactivates user
	DeactivateUser(ctx context.Context, walletAddress string) error
	// UpdateUser updates user
	UpdateUser(ctx context.Context, data types.UpdateUser) error
	// UpdateUserEmailVerifiedAt updates user email verified at
	UpdateUserEmailVerifiedAt(ctx context.Context, email string) error
	// UpdateUserPassword updates user password
	UpdateUserPassword(ctx context.Context, data types.UpdateUserPassword) error
}
