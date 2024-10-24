package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"celeste/infrastructures/database/mysql/types"
	apiError "celeste/internal/errors"
	"celeste/module/user/domain/entity"
)

// UserQueryRepository handles the user query repository logic
type UserQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectUserByID select a user by id
func (repository *UserQueryRepository) SelectUserByID(ID string) (entity.User, error) {
	var user entity.User

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", user.GetModelName())
	err := repository.QueryRow(stmt, map[string]interface{}{
		"id": ID,
	}, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New(apiError.MissingRecord)
		}

		return user, errors.New(apiError.DatabaseError)
	}

	return user, nil
}
