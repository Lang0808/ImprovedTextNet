package GrpcUserService

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	LibUtils "github.com/Lang0808/libutils"
)

var client UserServiceClient

func InitConnection() error {
	host := LibUtils.Get("UserService.host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Cannot initialize connection to UserService")
		return err
	}
	client = NewUserServiceClient(conn)
	return nil
}

func LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	return client.LoginUser(ctx, in)
}

func RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	return client.RegisterUser(ctx, in)
}

func GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	return client.GetUserInfo(ctx, in)
}
