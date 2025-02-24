package application

import (
	"context"

	"celeste/module/user/domain/entity"
)

// UserQueryServiceInterface holds the implementable methods for the user query service
type UserQueryServiceInterface interface {
	// GetUsers get all users
	GetUsers(ctx context.Context, page uint) ([]entity.User, uint, error)
	// GetUserByWalletAddress get the user provided by its wallet address
	GetUserByWalletAddress(ctx context.Context, walletAddress string) (entity.User, error)
	// GetUserByEmail get user by email
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	// GetUserSSS get user sss3
	GetUserSSS(ctx context.Context, walletAddress string) (entity.UserSSS, error)
}
