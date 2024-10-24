package application

import (
	"context"

	"celeste/module/user/domain/entity"
)

// UserQueryServiceInterface holds the implementable methods for the user query service
type UserQueryServiceInterface interface {
	GetUserByID(ctx context.Context, ID string) (entity.User, error)
}
