package repository

import (
	"celeste/module/iam/domain/repository"
)

// IAMQueryRepositoryCircuitBreaker holds the implementable methods for iam query circuitbreaker
type IAMQueryRepositoryCircuitBreaker struct {
	repository.IAMQueryRepositoryInterface
}
