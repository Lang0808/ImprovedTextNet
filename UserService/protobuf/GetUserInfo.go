package GrpcUserService

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	UserDB "UserService/db/UserDB"
)

func GetUserInfo_Server(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	id := sql.NullInt32{
		Int32: req.UserId,
		Valid: true,
	}
	User, err := UserDB.GetUserFromUserId(ctx, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when get user info from DB. Error: %v", err))
	}
	resp := GetUserInfoResponse{
		UserId:   req.UserId,
		Username: User.Username,
		Avatar:   User.Avatar.String,
	}
	return &resp, nil
}
