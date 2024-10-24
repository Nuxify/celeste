package rest

import (
	"celeste/module/iam/application"
)

// IAMQueryController request controller for iam query
type IAMQueryController struct {
	application.IAMQueryServiceInterface
}
