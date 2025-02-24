package repository

import (
	"celeste/module/user/domain/entity"
)

// UserQueryRepositoryInterface holds the implementable method for user query repository
type UserQueryRepositoryInterface interface {
	// SelectUsers select all users
	SelectUsers(page uint) ([]entity.User, uint, error)
	// SelectUserByWalletAddress select a user by wallet address
	SelectUserByWalletAddress(walletAddress string) (entity.User, error)
	// SelectUserByEmail select a user by email
	SelectUserByEmail(email string) (entity.User, error)
	// SelectUserSSS select user sss3
	SelectUserSSS(walletAddress string) (entity.UserSSS, error)
}
