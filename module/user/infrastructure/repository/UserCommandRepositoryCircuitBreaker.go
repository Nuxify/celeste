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
