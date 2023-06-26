package handler

import (
	GrpcUserService "WebServer/clients/UserService"
	"context"
	"fmt"
	"net/http"
)

type RegisterUserModel struct {
}

func (u RegisterUserModel) GetSrcId(r *http.Request) (int32, error) {
	return -1, nil
}

func (u RegisterUserModel) GetCommandId() int32 {
	return int32(REGISTER)
}

func (u RegisterUserModel) Handle(srcId int32, w http.ResponseWriter, r *http.Request,
	extra *map[string]string) {
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

func (u RegisterUserModel) PostProcess(w *http.ResponseWriter, r *http.Request) {
	DefaultPostProcess(w, r)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	requestHandler := RequestHandler{
		inner: RegisterUserModel{},
	}
	requestHandler.process(w, r)
}
