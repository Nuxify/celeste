package http

import (
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/go-playground/validator/v10"
)

var (
	Validate         *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
	ValidationErrors map[string]string   = map[string]string{
		"CreateUserRequest.Email":                    "Email field is required.",
		"CreateUserRequest.Password":                 "Password field is required.",
		"CreateUserRequest.Name":                     "Name field is required.",
		"ReconstructPrivateKeyRequest.ShareKey":      "ShareKey field is required.",
		"ReconstructPrivateKeyRequest.WalletAddress": "WalletAddress field is required.",
		"SignEIP191Request.ShareKey":                 "ShareKey field is required.",
		"SignEIP191Request.WalletAddress":            "WalletAddress field is required.",
		"SignEIP191Request.Message":                  "Message field is required.",
		"SignEIP712Request.ShareKey":                 "ShareKey field is required.",
		"SignEIP712Request.WalletAddress":            "WalletAddress field is required.",
		"SignEIP712Request.SignerData":               "Signer data is required.",
		"UpdateUserRequest.Name":                     "Name field is required.",
		"UpdateUserPasswordRequest.CurrentPassword":  "Current password field is required.",
		"UpdateUserPasswordRequest.NewPassword":      "New password field is required.",
	}
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type ReconstructPrivateKeyRequest struct {
	ShareKey      string `json:"shareKey" validate:"required"`
	WalletAddress string `json:"walletAddress" validate:"required"`
}

type SignEIP191Request struct {
	ShareKey      string `json:"shareKey" validate:"required"`
	WalletAddress string `json:"walletAddress" validate:"required"`
	Message       string `json:"message" validate:"required"`
}

type SignEIP712Request struct {
	ShareKey      string           `json:"shareKey" validate:"required"`
	WalletAddress string           `json:"walletAddress" validate:"required"`
	SignerData    EIP712SignerData `json:"signerData" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUserEmailVerifiedAtRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password" validate:"required"`
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
	Name            string  `json:"name"`
	EmailVerifiedAt *uint64 `json:"emailVerifiedAt"`
	CreatedAt       uint64  `json:"createdAt"`
	UpdatedAt       uint64  `json:"updatedAt"`
}

type GetPaginatedUserResponse struct {
	Users []GetUserResponse `json:"users"`
	Total uint              `json:"total"`
}

type ReconstructPrivateKeyResponse struct {
	PrivateKeyHex string `json:"privateKeyHex"`
	PublicKeyHex  string `json:"publicKeyHex"`
}

type SignEIP191Response struct {
	Signature string `json:"signature"`
}

type SignEIP712Response struct {
	Signature string `json:"signature"`
}

type EIP712SignerData struct {
	Types       map[string][]apitypes.Type `json:"types"`
	PrimaryType string                     `json:"primaryType"`
	Domain      apitypes.TypedDataDomain   `json:"domain"`
	Message     apitypes.TypedDataMessage  `json:"message"`
}
