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

// GetUserByID retrieves the user provided by its id
func (service *UserQueryService) GetUserByID(ctx context.Context, ID string) (entity.User, error) {
	res, err := service.UserQueryRepositoryInterface.SelectUserByID(ID)
	if err != nil {
		return res, err
	}

	return res, nil
}
