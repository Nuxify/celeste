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

// SelectUserByID decorator pattern for select user repository
func (repository *UserQueryRepositoryCircuitBreaker) SelectUserByID(ID string) (entity.User, error) {
	output := make(chan entity.User, 1)
	hystrix.ConfigureCommand("select_user_by_id", config.Settings())
	errors := hystrix.Go("select_user_by_id", func() error {
		user, err := repository.UserQueryRepositoryInterface.SelectUserByID(ID)
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
