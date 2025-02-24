package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth"

	"celeste/interfaces/http/rest/viewmodels"
	"celeste/internal/errors"
	"celeste/module/user/application"
	types "celeste/module/user/interfaces/http"
)

// UserQueryController request controller for user query
type UserQueryController struct {
	application.UserQueryServiceInterface
}

// GetUsers get all users
func (controller *UserQueryController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// pagination
	var page int
	if len(r.URL.Query().Get("page")) > 0 {
		var err error

		page, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   "Invalid page value.",
				ErrorCode: errors.InvalidRequestPayload,
			}

			response.JSON(w)
			return
		}

		if page == 0 {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusBadRequest,
				Success:   false,
				Message:   "Invalid page number.",
				ErrorCode: errors.InvalidRequestPayload,
			}

			response.JSON(w)
			return
		}
	}

	res, totalCount, err := controller.UserQueryServiceInterface.GetUsers(context.TODO(), uint(page))
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Database error."
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

	users := []types.GetUserResponse{}
	for _, user := range res {
		var emailVerifiedTimestamp *uint64
		if user.EmailVerifiedAt != nil {
			timestamp := uint64(user.EmailVerifiedAt.Unix())
			emailVerifiedTimestamp = &timestamp
		}

		users = append(users, types.GetUserResponse{
			WalletAddress:   user.WalletAddress,
			Email:           user.Email,
			Name:            user.Name,
			EmailVerifiedAt: emailVerifiedTimestamp,
			CreatedAt:       uint64(user.CreatedAt.Unix()),
			UpdatedAt:       uint64(user.UpdatedAt.Unix()),
		})
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched all users.",
		Data: &types.GetPaginatedUserResponse{
			Users:      users,
			TotalCount: totalCount,
		},
	}

	response.JSON(w)
}

// GetCurrentUser get current user
func (controller *UserQueryController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	walletAddress := claims["id"].(string)

	res, err := controller.UserQueryServiceInterface.GetUserByWalletAddress(context.TODO(), walletAddress)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No records found."

		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Database error."
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

	user := &types.GetUserResponse{
		WalletAddress: res.WalletAddress,
		Email:         res.Email,
		Name:          res.Name,
		CreatedAt:     uint64(res.CreatedAt.Unix()),
		UpdatedAt:     uint64(res.UpdatedAt.Unix()),
	}

	if res.EmailVerifiedAt != nil {
		timestamp := uint64(res.EmailVerifiedAt.Unix())
		user.EmailVerifiedAt = &timestamp
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched current user.",
		Data:    user,
	}

	response.JSON(w)
}
