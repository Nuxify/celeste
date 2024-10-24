package service

import (
	"context"

	"github.com/segmentio/ksuid"

	"celeste/module/user/domain/entity"
	"celeste/module/user/domain/repository"
	repositoryTypes "celeste/module/user/infrastructure/repository/types"
	"celeste/module/user/infrastructure/service/types"
)

// UserCommandService handles the user command service logic
type UserCommandService struct {
	repository.UserCommandRepositoryInterface
}

// CreateUser create a user
func (service *UserCommandService) CreateUser(ctx context.Context, data types.CreateUser) (entity.User, error) {
	user := repositoryTypes.CreateUser{
		ID:   data.ID,
		Data: data.Data,
	}

	// check id if empty create new unique id
	if len(user.ID) == 0 {
		user.ID = generateID()
	}

	res, err := service.UserCommandRepositoryInterface.InsertUser(user)
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}

// generateID generates unique id
func generateID() string {
	return ksuid.New().String()
}
