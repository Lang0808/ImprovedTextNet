package GrpcUserService

import (
	context "context"
)

type UserServiceServerImpl struct {
}

func (u UserServiceServerImpl) LoginUser(ctx context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {
	return Login(ctx, request)
}

func (u UserServiceServerImpl) RegisterUser(ctx context.Context, request *RegisterUserRequest) (*RegisterUserResponse, error) {
	return Register(ctx, request)
}

func (u UserServiceServerImpl) GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return GetUserInfo_Server(ctx, req)
}

func (u UserServiceServerImpl) mustEmbedUnimplementedUserServiceServer() {

}
