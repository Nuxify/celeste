package application

import (
	"context"

	"celeste/module/user/domain/entity"
	"celeste/module/user/infrastructure/service/types"
)

// UserCommandServiceInterface holds the implementable methods for the user command service
type UserCommandServiceInterface interface {
	CreateUser(ctx context.Context, data types.CreateUser) (entity.User, error)
}
