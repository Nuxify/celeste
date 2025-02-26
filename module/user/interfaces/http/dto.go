package http

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validate         *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
	ValidationErrors map[string]string   = map[string]string{
		"CreateUserRequest.Email":                   "Email field is required.",
		"CreateUserRequest.Password":                "Password field is required.",
		"CreateUserRequest.Name":                    "Name field is required.",
		"UpdateUserRequest.Name":                    "Name field is required.",
		"UpdateUserPasswordRequest.CurrentPassword": "Current password field is required.",
		"UpdateUserPasswordRequest.NewPassword":     "New password field is required.",
	}
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUserEmailVerifiedAtRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserPasswordRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
}

type GetUserResponse struct {
	WalletAddress   string  `json:"wallet_address"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	EmailVerifiedAt *uint64 `json:"email_verified_at"`
	CreatedAt       uint64  `json:"created_at"`
	UpdatedAt       uint64  `json:"updated_at"`
}

type GetPaginatedUserResponse struct {
	Users      []GetUserResponse `json:"users"`
	TotalCount uint              `json:"total_count"`
}

type CreateUserResponse struct {
	SSS2 string `json:"sss2"`
	SSS3 string `json:"sss3"`
}
