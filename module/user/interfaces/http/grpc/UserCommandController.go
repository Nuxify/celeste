package grpc

import (
	"celeste/module/user/application"
)

// UserCommandController handles the grpc user command requests
type UserCommandController struct {
	application.UserCommandServiceInterface
}

// // CreateUser creates a new user
// func (controller *UserCommandController) CreateUser(ctx context.Context, req *grpcPB.CreateUserRequest) (*grpcPB.UserResponse, error) {
// 	user := serviceTypes.CreateUser{
// 		ID:   req.Id,
// 		Data: req.Data,
// 	}

// 	res, err := controller.UserCommandServiceInterface.CreateUser(context.TODO(), user)
// 	if err != nil { 
// 		var code codes.Code

// 		switch err.Error() {
// 		case errors.DatabaseError:
// 			code = codes.Internal
// 		case errors.MissingRecord:
// 			code = codes.NotFound
// 		default:
// 			code = codes.Unknown
// 		}

// 		st := status.New(code, fmt.Sprintf("[RECORD] %s", err.Error()))

// 		return nil, st.Err()
// 	}

// 	createProtoTime, _ := ptypes.TimestampProto(time.Now())

// 	return &grpcPB.UserResponse{
// 		Id:        res.ID,
// 		Data:      res.Data,
// 		CreatedAt: createProtoTime,
// 	}, nil
// }
