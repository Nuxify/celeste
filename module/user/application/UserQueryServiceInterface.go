package application

import (
	"context"

	"celeste/module/user/domain/entity"
)

// UserQueryServiceInterface holds the implementable methods for the user query service
type UserQueryServiceInterface interface {
	// GetUsers get all users
	GetUsers(ctx context.Context, page uint, search *string) ([]entity.User, uint, error)
	// GetUserByEmail get user by email
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	// GetUserByWalletAddress get the user provided by its wallet address
	GetUserByWalletAddress(ctx context.Context, walletAddress string) (entity.User, error)
}
