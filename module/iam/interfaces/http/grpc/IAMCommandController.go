package grpc

import (
	"celeste/module/iam/application"
)

// IAMCommandController handles the grpc iam command requests
type IAMCommandController struct {
	application.IAMCommandServiceInterface
}
