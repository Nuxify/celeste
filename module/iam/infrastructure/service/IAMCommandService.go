package service

import (
	"github.com/segmentio/ksuid"

	"celeste/module/iam/domain/repository"
)

// IAMCommandService handles the iam command service logic
type IAMCommandService struct {
	repository.IAMCommandRepositoryInterface
}

// generateID generates unique id
func generateID() string {
	return ksuid.New().String()
}
