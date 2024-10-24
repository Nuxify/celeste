package grpc

import (
	"celeste/module/iam/application"
)

// IAMQueryController handles the grpc iam query requests
type IAMQueryController struct {
	application.IAMQueryServiceInterface
}
