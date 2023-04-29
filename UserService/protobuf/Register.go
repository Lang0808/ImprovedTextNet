package GrpcUserService

import (
	"UserService/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"

	UserDB "UserService/db/UserDB"
)

func ValidateRegisterUserRequest(request *RegisterUserRequest) bool {
	if len(request.GetUsername()) < 8 || len(request.GetUsername()) > 50 {
		return false
	}
	if len(request.GetPassword()) < 8 || len(request.GetPassword()) > 50 {
		return false
	}
	return true
}

func Register(ctx context.Context, request *RegisterUserRequest) (*RegisterUserResponse, error) {
	if !ValidateRegisterUserRequest(request) {
		return nil, errors.New("Request is invalid")
	}
	password, err := utils.Hash(request.Password)
	if err != nil {
		fmt.Printf("Hash password fail. Error: %v", err)
		return nil, errors.New(fmt.Sprintf("Hash Password fail. Error: %v", err))
	}
	passwordTemp := sql.NullString{
		String: password,
		Valid:  true,
	}
	Avatar := sql.NullString{
		String: "",
		Valid:  true,
	}
	in := UserDB.CreateUserParams{
		Username: request.Username,
		Password: passwordTemp,
		Avatar:   Avatar,
	}
	err = UserDB.CreateUser(ctx, in)
	if err != nil {
		fmt.Printf("UserDB.CreateUser fail. Error: %v", err)
		return nil, errors.New(fmt.Sprint("Create User fail. Error: %v", err))
	}
	return &RegisterUserResponse{}, nil
}
