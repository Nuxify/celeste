package grpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"celeste/internal/errors"
	"celeste/module/user/application"
	grpcPB "celeste/module/user/interfaces/http/grpc/pb"
)

// UserQueryController handles the grpc user query requests
type UserQueryController struct {
	application.UserQueryServiceInterface
}

// GetUserByID retrieves the user id from the proto
func (controller *UserQueryController) GetUserByID(ctx context.Context, req *grpcPB.GetUserRequest) (*grpcPB.UserResponse, error) {
	res, err := controller.UserQueryServiceInterface.GetUserByID(context.TODO(), req.Id)
	if err != nil {
		var code codes.Code

		switch err.Error() {
		case errors.DatabaseError:
			code = codes.Internal
		case errors.MissingRecord:
			code = codes.NotFound
		default:
			code = codes.Unknown
		}

		st := status.New(code, fmt.Sprintf("[RECORD] %s", err.Error()))

		return nil, st.Err()
	}

	createProtoTime, _ := ptypes.TimestampProto(res.CreatedAt)

	return &grpcPB.UserResponse{
		Id:        res.ID,
		Data:      res.Data,
		CreatedAt: createProtoTime,
	}, nil
}
