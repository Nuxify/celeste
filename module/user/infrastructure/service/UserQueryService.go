package service

import (
	"context"

	"celeste/module/user/domain/entity"
	"celeste/module/user/domain/repository"
)

// UserQueryService handles the user query service logic
type UserQueryService struct {
	repository.UserQueryRepositoryInterface
}

// GetUsers get all users
func (service *UserQueryService) GetUsers(ctx context.Context, page uint) ([]entity.User, uint, error) {
	res, totalCount, err := service.UserQueryRepositoryInterface.SelectUsers(page)
	if err != nil {
		return []entity.User{}, 0, err
	}

	return res, totalCount, nil
}

// GetUserByEmail get user by email
func (service *UserQueryService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user, err := service.UserQueryRepositoryInterface.SelectUserByEmail(email)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// GetUserByWalletAddress get the user provided by its wallet address
func (service *UserQueryService) GetUserByWalletAddress(ctx context.Context, walletAddress string) (entity.User, error) {
	res, err := service.UserQueryRepositoryInterface.SelectUserByWalletAddress(walletAddress)
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}
