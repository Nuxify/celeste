package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"celeste/module/user/domain/entity"
	"celeste/module/user/domain/repository"
)

// UserQueryRepositoryCircuitBreaker holds the implementable methods for user query circuitbreaker
type UserQueryRepositoryCircuitBreaker struct {
	repository.UserQueryRepositoryInterface
}

// SelectUsers is a decorator for the select users repository
func (repository *UserQueryRepositoryCircuitBreaker) SelectUsers(page uint) ([]entity.User, uint, error) {
	type outputData struct {
		Users      []entity.User
		TotalCount uint
	}
	output := make(chan outputData, 1)
	errChan := make(chan error, 1)
	hystrix.ConfigureCommand("select_users", config.Settings())
	errors := hystrix.Go("select_users", func() error {
		users, totalCount, err := repository.UserQueryRepositoryInterface.SelectUsers(page)
		if err != nil {
			errChan <- err
			return nil
		}

		result := outputData{
			Users:      users,
			TotalCount: totalCount,
		}

		output <- result
		return nil
	}, nil)

	select {
	case out := <-output:
		return out.Users, out.TotalCount, nil
	case err := <-errChan:
		return []entity.User{}, 0, err
	case err := <-errors:
		return []entity.User{}, 0, err
	}
}

// SelectUserByWalletAddress decorator pattern for select user repository
func (repository *UserQueryRepositoryCircuitBreaker) SelectUserByWalletAddress(walletAddress string) (entity.User, error) {
	output := make(chan entity.User, 1)
	errChan := make(chan error, 1)

	hystrix.ConfigureCommand("select_user_by_wallet_address", config.Settings())
	errors := hystrix.Go("select_user_by_wallet_address", func() error {
		user, err := repository.UserQueryRepositoryInterface.SelectUserByWalletAddress(walletAddress)
		if err != nil {
			errChan <- err
			return nil
		}

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errChan:
		return entity.User{}, err
	case err := <-errors:
		return entity.User{}, err
	}
}

// SelectUserByEmail decorator pattern for select user repository
func (repository *UserQueryRepositoryCircuitBreaker) SelectUserByEmail(email string) (entity.User, error) {
	output := make(chan entity.User, 1)
	errChan := make(chan error, 1)

	hystrix.ConfigureCommand("select_user_by_email", config.Settings())
	errors := hystrix.Go("select_user_by_email", func() error {
		user, err := repository.UserQueryRepositoryInterface.SelectUserByEmail(email)
		if err != nil {
			errChan <- err
			return nil
		}

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errChan:
		return entity.User{}, err
	case err := <-errors:
		return entity.User{}, err
	}
}