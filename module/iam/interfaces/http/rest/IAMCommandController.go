package rest

import (
	"celeste/module/iam/application"
)

// IAMCommandController request controller for iam command
type IAMCommandController struct {
	application.IAMCommandServiceInterface
}
