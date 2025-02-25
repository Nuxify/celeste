package repository

import (
	"errors"
	"fmt"
	"log"

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
func (repository *UserCommandRepository) InsertUser(data repositoryTypes.CreateUser) error {
	user := entity.User{
		WalletAddress: data.WalletAddress,
		Email:         data.Email,
		Password:      data.Password,
		SSS1:          data.SSS1,
		Name:          data.Name,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (wallet_address, email, password, sss_1, name) VALUES (:wallet_address, :email, :password, :sss_1, :name)", user.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, user)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.New(apiError.DuplicateRecord)
		}
		return errors.New(apiError.DatabaseError)
	}

	// insert user sss record
	err = repository.insertUserSSS(data.SSS3, data.WalletAddress)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser update user
func (repository *UserCommandRepository) UpdateUser(data repositoryTypes.UpdateUser) error {
	user := entity.User{
		WalletAddress: data.WalletAddress,
		Name:          data.Name,
	}

	// update user
	stmt := fmt.Sprintf("UPDATE %s SET name=:name "+
		"WHERE wallet_address=:wallet_address", user.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, user)
	if err != nil {
		log.Println(err)
		return errors.New(apiError.DatabaseError)
	}

	return nil
}

// UpdateUserPassword updates user password
func (repository *UserCommandRepository) UpdateUserPassword(data repositoryTypes.UpdateUserPassword) error {
	user := &entity.User{
		WalletAddress: data.WalletAddress,
		Password:      data.Password,
	}

	// update users
	stmt := fmt.Sprintf("UPDATE %s SET password=:password "+
		"WHERE wallet_address=:wallet_address", user.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, user)
	if err != nil {
		log.Println(err)
		return errors.New(apiError.DatabaseError)
	}

	return nil
}

// insertUserSSS creates a new user sss
func (repository *UserCommandRepository) insertUserSSS(sss3, userWalletAddress string) error {
	sss := entity.UserSSS{
		SSS3:              sss3,
		UserWalletAddress: userWalletAddress,
	}

	// insert user sss
	stmt := fmt.Sprintf("INSERT INTO %s (user_wallet_address,sss_3) "+
		"VALUES (:user_wallet_address,:sss_3)", sss.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, sss)
	if err != nil {
		log.Println(err)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.New(apiError.DuplicateRecord)
		}

		return errors.New(apiError.DatabaseError)
	}

	return nil
}
