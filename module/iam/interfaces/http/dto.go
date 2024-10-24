package http

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validate         *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
	ValidationErrors map[string]string   = map[string]string{
		"CreateUserRequest.ID":   "ID field is required.",
		"CreateUserRequest.Data": "Data field is required.",
	}
)

// CreateUserRequest request struct for create user
type CreateUserRequest struct {
	ID   string `json:"id" validate:"required"`
	Data string `json:"data" validate:"required"`
}

// UserResponse response struct
type UserResponse struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"createdAt"`
}
