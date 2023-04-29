package GrpcUserService

import (
	"context"
	"errors"
	"fmt"

	"UserService/utils"

	UserDB "UserService/db/UserDB"
)

func ValidateLoginUserRequest(request *LoginUserRequest) bool {
	if len(request.GetUsername()) < 8 || len(request.GetUsername()) > 50 {
		return false
	}
	if len(request.GetPassword()) < 8 || len(request.GetPassword()) > 50 {
		return false
	}
	return true
}

func Login(ctx context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {
	if !ValidateLoginUserRequest(request) {
		return nil, errors.New("Request is invalid")
	}
	user, err := UserDB.GetUserFromUsername(ctx, request.GetUsername())
	if err != nil {
		fmt.Printf("UserDB.GetUserFromUsername fail. Error: %v", err)
		return nil, err
	}
	if !utils.CompareHash(request.GetPassword(), user.Password.String) {
		return nil, errors.New("Username or password is incorrect")
	}
	return &LoginUserResponse{
		UserId:   user.ID.Int32,
		Username: user.Username,
		Avatar:   user.Avatar.String,
	}, nil
}
