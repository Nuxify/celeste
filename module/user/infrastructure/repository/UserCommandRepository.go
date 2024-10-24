package repository

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"

	"celeste/infrastructures/database/mysql/types"
	apiError "celeste/internal/errors"
	"celeste/module/user/domain/entity"
	repositoryTypes "celeste/module/user/infrastructure/repository/types"
)

// UserCommandRepository handles the user command repository logic
type UserCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// InsertUser creates a new user
func (repository *UserCommandRepository) InsertUser(data repositoryTypes.CreateUser) (entity.User, error) {
	user := entity.User{
		ID:   data.ID,
		Data: data.Data,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (id, data) VALUES (:id, :data)", user.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, user)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return entity.User{}, errors.New(apiError.DuplicateRecord)
		}
		return entity.User{}, errors.New(apiError.DatabaseError)
	}

	return user, nil
}
