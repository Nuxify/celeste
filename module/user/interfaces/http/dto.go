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

type CreateUserResponse struct {
	WalletAddress string `json:"walletAddress"`
	SSS2          string `json:"sss2"`
	SSS3          string `json:"sss3"`
}

type GetUserResponse struct {
	WalletAddress   string  `json:"walletAddress"`
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	SSS1            string  `json:"sss1"`
	Name            string  `json:"name"`
	EmailVerifiedAt *uint64 `json:"emailVerifiedAt"`
	CreatedAt       uint64  `json:"createdAt"`
	UpdatedAt       uint64  `json:"updatedAt"`
}

type GetPaginatedUserResponse struct {
	Users []GetUserResponse `json:"users"`
	Total uint              `json:"total"`
}
