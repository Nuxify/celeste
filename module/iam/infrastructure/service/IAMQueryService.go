package service

import (
	"celeste/module/iam/domain/repository"
)

// IAMQueryService handles the iam query service logic
type IAMQueryService struct {
	repository.IAMQueryRepositoryInterface
}
