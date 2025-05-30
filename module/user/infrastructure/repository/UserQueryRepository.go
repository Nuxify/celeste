package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"celeste/infrastructures/database/mysql/types"
	apiError "celeste/internal/errors"
	"celeste/module/user/domain/entity"
)

// UserQueryRepository handles the user query repository logic
type UserQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectUsers select all users
func (repository *UserQueryRepository) SelectUsers(page uint, search *string) ([]entity.User, uint, error) {
	var user entity.User
	var users []entity.User

	stmt := fmt.Sprintf("SELECT * FROM %s", user.GetModelName())

	// if search is set
	if search != nil {
		stmt = fmt.Sprintf("%s WHERE (name LIKE %s OR email LIKE %s)", stmt, "'%"+*search+"%'", "'%"+*search+"%'")
	}

	stmt = fmt.Sprintf("%s ORDER BY updated_at DESC", stmt)

	// get total count
	var counter struct {
		Total uint `json:"total"`
	}
	totalCountStmt := strings.ReplaceAll(stmt, "SELECT *", "SELECT COUNT(*) as total")

	err := repository.QueryRow(totalCountStmt, map[string]interface{}{}, &counter)
	if err != nil {
		log.Println(err)
		return []entity.User{}, 0, errors.New(apiError.DatabaseError)
	}

	// apply pagination
	if page > 0 {
		var limit uint = 10
		offset := limit * (page - 1)

		stmt = fmt.Sprintf("%s LIMIT %d OFFSET %d", stmt, limit, offset)
	}

	err = repository.Query(stmt, map[string]interface{}{}, &users)
	if err != nil {
		log.Println(err)
		return []entity.User{}, 0, errors.New(apiError.DatabaseError)
	} else if len(users) == 0 {
		return []entity.User{}, 0, errors.New(apiError.MissingRecord)
	}

	return users, counter.Total, nil
}

// SelectUserByWalletAddress select a user by wallet address
func (repository *UserQueryRepository) SelectUserByWalletAddress(walletAddress string) (entity.User, error) {
	var user entity.User

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE wallet_address=:wallet_address", user.GetModelName())
	err := repository.QueryRow(stmt, map[string]interface{}{
		"wallet_address": walletAddress,
	}, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New(apiError.MissingRecord)
		}

		log.Println(err)
		return user, errors.New(apiError.DatabaseError)
	}

	return user, nil
}

// SelectUserByEmail select a user by email
func (repository *UserQueryRepository) SelectUserByEmail(email string) (entity.User, error) {
	var user entity.User

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE email=:email", user.GetModelName())
	err := repository.QueryRow(stmt, map[string]interface{}{
		"email": email,
	}, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New(apiError.MissingRecord)
		}

		log.Println(err)
		return user, errors.New(apiError.DatabaseError)
	}

	return user, nil
}
