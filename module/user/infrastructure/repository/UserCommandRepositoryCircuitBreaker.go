package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "celeste/configs/hystrix"
	"celeste/module/user/domain/entity"
	"celeste/module/user/domain/repository"
	repositoryTypes "celeste/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryCircuitBreaker circuit breaker for user command repository
type UserCommandRepositoryCircuitBreaker struct {
	repository.UserCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// InsertUser decorator pattern to insert user
func (repository *UserCommandRepositoryCircuitBreaker) InsertUser(data repositoryTypes.CreateUser) (entity.User, error) {
	output := make(chan entity.User, 1)
	hystrix.ConfigureCommand("insert_user", config.Settings())
	errors := hystrix.Go("insert_user", func() error {
		user, err := repository.UserCommandRepositoryInterface.InsertUser(data)
		if err != nil {
			return err
		}

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.User{}, err
	}
}
