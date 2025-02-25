package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "celeste/configs/hystrix"
	"celeste/module/user/domain/repository"
	repositoryTypes "celeste/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryCircuitBreaker circuit breaker for user command repository
type UserCommandRepositoryCircuitBreaker struct {
	repository.UserCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// InsertUser decorator pattern to insert user
func (repository *UserCommandRepositoryCircuitBreaker) InsertUser(data repositoryTypes.CreateUser) error {
	output := make(chan error, 1)
	errChan := make(chan error, 1)

	hystrix.ConfigureCommand("insert_user", config.Settings())
	errors := hystrix.Go("insert_user", func() error {
		err := repository.UserCommandRepositoryInterface.InsertUser(data)
		if err != nil {
			errChan <- err
			return nil
		}

		output <- nil
		return nil
	}, nil)

	select {
	case out := <-output:
		return out
	case err := <-errChan:
		return err
	case err := <-errors:
		return err
	}
}

// UpdateUser decorator pattern to update user
func (repository *UserCommandRepositoryCircuitBreaker) UpdateUser(data repositoryTypes.UpdateUser) error {
	output := make(chan error, 1)
	errChan := make(chan error, 1)

	hystrix.ConfigureCommand("update_user", config.Settings())
	errors := hystrix.Go("update_user", func() error {
		err := repository.UserCommandRepositoryInterface.UpdateUser(data)
		if err != nil {
			errChan <- err
			return nil
		}

		output <- nil
		return nil
	}, nil)

	select {
	case out := <-output:
		return out
	case err := <-errChan:
		return err
	case err := <-errors:
		return err
	}
}

// UpdateUserEmailVerifiedAt decorator pattern to update user email verified at
func (repository *UserCommandRepositoryCircuitBreaker) UpdateUserEmailVerifiedAt(email string) error {
	output := make(chan error, 1)
	errChan := make(chan error, 1)

	hystrix.ConfigureCommand("update_user_email_verified_at", config.Settings())
	errors := hystrix.Go("update_user_email_verified_at", func() error {
		err := repository.UserCommandRepositoryInterface.UpdateUserEmailVerifiedAt(email)
		if err != nil {
			errChan <- err
			return nil
		}

		output <- nil
		return nil
	}, nil)

	select {
	case out := <-output:
		return out
	case err := <-errChan:
		return err
	case err := <-errors:
		return err
	}
}

// UpdateUserPassword decorator pattern to update user password
func (repository *UserCommandRepositoryCircuitBreaker) UpdateUserPassword(data repositoryTypes.UpdateUserPassword) error {
	output := make(chan error, 1)
	errChan := make(chan error, 1)

	hystrix.ConfigureCommand("update_user_password", config.Settings())
	errors := hystrix.Go("update_user_password", func() error {
		err := repository.UserCommandRepositoryInterface.UpdateUserPassword(data)
		if err != nil {
			errChan <- err
			return nil
		}

		output <- nil
		return nil
	}, nil)

	select {
	case out := <-output:
		return out
	case err := <-errChan:
		return err
	case err := <-errors:
		return err
	}
}
