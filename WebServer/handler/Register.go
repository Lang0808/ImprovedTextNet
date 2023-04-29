package handler

import (
	GrpcUserService "WebServer/clients/UserService"
	"context"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	request := GrpcUserService.RegisterUserRequest{
		Username: username,
		Password: password,
	}

	_, err := GrpcUserService.RegisterUser(context.Background(), &request)
	var resp ApiMessage
	if err != nil {
		resp = ApiMessage{
			Err:     -1,
			Message: fmt.Sprintf("UserServiceError: %v", err),
			Data:    "",
		}
	} else {
		resp = ApiMessage{
			Err:     0,
			Message: "Register successfully !",
			Data:    "",
		}
	}
	str := GetApiMessage(resp)
	fmt.Fprintf(w, str)
}
