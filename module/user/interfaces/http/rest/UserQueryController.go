package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"celeste/interfaces/http/rest/viewmodels"
	"celeste/internal/errors"
	"celeste/module/user/application"
	types "celeste/module/user/interfaces/http"
)

// UserQueryController request controller for user query
type UserQueryController struct {
	application.UserQueryServiceInterface
}

// GetUserByID retrieves the tenant id from the rest request
func (controller *UserQueryController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	if len(userID) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid user ID",
			ErrorCode: errors.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	res, err := controller.UserQueryServiceInterface.GetUserByID(context.TODO(), userID)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error while fetching user."
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No user found."
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
		Message: "User successfully fetched.",
		Data: &types.UserResponse{
			ID:        res.ID,
			Data:      res.Data,
			CreatedAt: res.CreatedAt.Unix(),
		},
	}

	response.JSON(w)
}
