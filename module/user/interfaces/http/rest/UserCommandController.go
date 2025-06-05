package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"celeste/interfaces/http/rest/viewmodels"
	"celeste/internal/errors"
	apiError "celeste/internal/errors"
	"celeste/module/user/application"
	serviceTypes "celeste/module/user/infrastructure/service/types"
	types "celeste/module/user/interfaces/http"
)

// UserCommandController request controller for user command
type UserCommandController struct {
	application.UserCommandServiceInterface
}

// CreateUser request handler to create user
func (controller *UserCommandController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request types.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// validate request
	err := types.Validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   types.ValidationErrors[errors[0].StructNamespace()],
				ErrorCode: apiError.InvalidPayload,
			}

			response.JSON(w)
			return
		}

		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	res, err := controller.UserCommandServiceInterface.CreateUser(context.TODO(), serviceTypes.CreateUser{
		Email:    strings.ToLower(request.Email),
		Password: request.Password,
		Name:     request.Name,
	})
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving user."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "User ID already exist."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully created user.",
		Data: &types.CreateUserResponse{
			WalletAddress: res.WalletAddress,
			SSS2:          res.SSS2,
			SSS3:          res.SSS3,
		},
	}

	response.JSON(w)
}

// DeactivateUser request handler to deactivate user
func (controller *UserCommandController) DeactivateUser(w http.ResponseWriter, r *http.Request) {
	walletAddress := chi.URLParam(r, "walletAddress")
	if len(walletAddress) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Wallet address is required.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	err := controller.UserCommandServiceInterface.DeactivateUser(context.TODO(), walletAddress)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating user."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully deactivated user.",
	}

	response.JSON(w)
}

// SignEIP191 request handler to sign EIP-191 message
func (controller *UserCommandController) SignEIP191(w http.ResponseWriter, r *http.Request) {
	var request types.SignEIP191Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// validate request
	err := types.Validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   types.ValidationErrors[errors[0].StructNamespace()],
				ErrorCode: apiError.InvalidPayload,
			}

			response.JSON(w)
			return
		}

		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	signature, err := controller.UserCommandServiceInterface.SignEIP191(context.TODO(), serviceTypes.SignEIP191{
		WalletAddress: request.WalletAddress,
		Message:       request.Message,
		ShareKey:      request.ShareKey,
	})
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.EthInvalidUserPrivateKey,
			errors.EthInvalidUserPublicKey,
			errors.UnauthorizedAccess:
			httpCode = http.StatusInternalServerError
			errorMsg = "Signer wallet error."
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating user email verified at."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully updated user email verified at.",
		Data: &types.SignEIP191RequestResponse{
			Signature: signature,
		},
	}

	response.JSON(w)
}

// SignEIP712 request handler to sign EIP-712 message
func (controller *UserCommandController) SignEIP712(w http.ResponseWriter, r *http.Request) {
	var request types.SignEIP712Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// validate request
	err := types.Validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   types.ValidationErrors[errors[0].StructNamespace()],
				ErrorCode: apiError.InvalidPayload,
			}

			response.JSON(w)
			return
		}

		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	signature, err := controller.UserCommandServiceInterface.SignEIP712(context.TODO(), serviceTypes.SignEIP712{
		WalletAddress: request.WalletAddress,
		ShareKey:      request.ShareKey,
		SignerData: apitypes.TypedData{
			Types:       request.SignerData.Types,
			PrimaryType: request.SignerData.PrimaryType,
			Domain: apitypes.TypedDataDomain{
				Name:              request.SignerData.Domain.Name,
				Version:           request.SignerData.Domain.Version,
				ChainId:           request.SignerData.Domain.ChainId,
				VerifyingContract: request.SignerData.Domain.VerifyingContract,
			},
			Message: request.SignerData.Message,
		},
	})
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.EthInvalidUserPrivateKey,
			errors.EthInvalidUserPublicKey,
			errors.UnauthorizedAccess:
			httpCode = http.StatusInternalServerError
			errorMsg = "Signer wallet error."
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while signing EIP-712 message."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully signed EIP-712 message.",
		Data: &types.SignEIP712RequestResponse{
			Signature: signature,
		},
	}

	response.JSON(w)
}

// UpdateUserEmailVerifiedAt request handler to update user email verified at
func (controller *UserCommandController) UpdateUserEmailVerifiedAt(w http.ResponseWriter, r *http.Request) {
	var request types.UpdateUserEmailVerifiedAtRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// validate request
	err := types.Validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   types.ValidationErrors[errors[0].StructNamespace()],
				ErrorCode: apiError.InvalidPayload,
			}

			response.JSON(w)
			return
		}

		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	err = controller.UserCommandServiceInterface.UpdateUserEmailVerifiedAt(context.TODO(), request.Email)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating user email verified at."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully updated user email verified at.",
	}

	response.JSON(w)
}

// UpdateUserByWalletAddress request handler to update user
func (controller *UserCommandController) UpdateUserByWalletAddress(w http.ResponseWriter, r *http.Request) {
	walletAddress := chi.URLParam(r, "walletAddress")
	if len(walletAddress) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Wallet address is required.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	var request types.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// validate request
	err := types.Validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   types.ValidationErrors[errors[0].StructNamespace()],
				ErrorCode: apiError.InvalidPayload,
			}

			response.JSON(w)
			return
		}

		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	err = controller.UserCommandServiceInterface.UpdateUser(context.TODO(), serviceTypes.UpdateUser{
		WalletAddress: walletAddress,
		Name:          request.Name,
	})
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating user."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully updated user.",
	}

	response.JSON(w)
}

// UpdateUserPassword request handler to update user password
func (controller *UserCommandController) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	walletAddress := chi.URLParam(r, "walletAddress")
	if len(walletAddress) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "ID is required.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	var request types.UpdateUserPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// validate request
	err := types.Validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   types.ValidationErrors[errors[0].StructNamespace()],
				ErrorCode: apiError.InvalidPayload,
			}

			response.JSON(w)
			return
		}

		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid payload request.",
			ErrorCode: apiError.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	err = controller.UserCommandServiceInterface.UpdateUserPassword(context.TODO(), serviceTypes.UpdateUserPassword{
		WalletAddress: walletAddress,
		Password:      request.Password,
	})
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating password."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully updated password.",
	}

	response.JSON(w)
}
