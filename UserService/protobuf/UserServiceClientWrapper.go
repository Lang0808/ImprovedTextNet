package GrpcUserService

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var client UserServiceClient

func InitConnection() error {
	host := viper.GetString("UserService.host")
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

func AddOrUpdateRelationship(ctx context.Context, in *AddOrUpdateRelationshipRequest, opts ...grpc.CallOption) (*AddOrUpdateRelationshipResponse, error) {
	return client.AddOrUpdateRelationship(ctx, in)
}
