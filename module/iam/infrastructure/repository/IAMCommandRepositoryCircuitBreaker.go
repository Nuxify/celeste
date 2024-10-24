package repository

import (
	hystrix_config "celeste/configs/hystrix"
	"celeste/module/iam/domain/repository"
)

// IAMCommandRepositoryCircuitBreaker circuit breaker for iam command repository
type IAMCommandRepositoryCircuitBreaker struct {
	repository.IAMCommandRepositoryInterface
}

var config = hystrix_config.Config{}
