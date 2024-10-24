package repository

import (
	"celeste/module/user/domain/entity"
	"celeste/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryInterface holds the implementable methods for user command repository
type UserCommandRepositoryInterface interface {
	InsertUser(data types.CreateUser) (entity.User, error)
}
