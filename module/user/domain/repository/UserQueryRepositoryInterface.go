package repository

import (
	"celeste/module/user/domain/entity"
)

// UserQueryRepositoryInterface holds the implementable method for user query repository
type UserQueryRepositoryInterface interface {
	SelectUserByID(ID string) (entity.User, error)
}
