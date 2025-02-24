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

// SelectUsers select all users
func (repository *UserQueryRepository) SelectUsers(page uint) ([]entity.User, uint, error) {
	var user entity.User
	var users []entity.User

	stmt := fmt.Sprintf("SELECT * FROM %s ORDER BY updated_at DESC", user.GetModelName())

	// apply pagination
	if page > 0 {
		var limit uint = 10
		offset := limit * (page - 1)

		stmt = fmt.Sprintf("%s LIMIT %d OFFSET %d", stmt, limit, offset)
	}

	err := repository.Query(stmt, map[string]interface{}{}, &users)
	if err != nil {
		return []entity.User{}, 0, errors.New(apiError.DatabaseError)
	} else if len(users) == 0 {
		return []entity.User{}, 0, errors.New(apiError.MissingRecord)
	}

	// get total count
	var counter struct {
		Total uint `json:"total"`
	}
	stmt = fmt.Sprintf("SELECT COUNT(*) as total FROM %s", user.GetModelName())
	err = repository.QueryRow(stmt, map[string]interface{}{}, &counter)
	if err != nil {
		return []entity.User{}, 0, errors.New(apiError.DatabaseError)
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

		return user, errors.New(apiError.DatabaseError)
	}

	return user, nil
}

// SelectUserSSS select user sss3
func (repository *UserQueryRepository) SelectUserSSS(walletAddress string) (entity.UserSSS, error) {
	var sss entity.UserSSS

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE user_wallet_address=:user_wallet_address", sss.GetModelName())
	err := repository.QueryRow(stmt, map[string]interface{}{
		"user_wallet_address": walletAddress,
	}, &sss)
	if err != nil {
		if err == sql.ErrNoRows {
			return sss, errors.New(apiError.MissingRecord)
		}

		return sss, errors.New(apiError.DatabaseError)
	}

	return sss, nil
}
