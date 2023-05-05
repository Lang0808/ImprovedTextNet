package handler

import (
	GrpcUserService "WebServer/clients/UserService"
	utils "WebServer/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type LoginResponse struct {
	Username string
	Avatar   string
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", viper.GetString("client.domain"))
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")

	Username := r.FormValue("Username")
	Password := r.FormValue("Password")
	in := &GrpcUserService.LoginUserRequest{
		Username: Username,
		Password: Password,
	}
	res, err := GrpcUserService.LoginUser(context.Background(), in)
	var resp ApiMessage
	if err != nil {
		resp := ApiMessage{
			Err:     -1,
			Message: fmt.Sprintf("GrpcUserService.LoginUser error. Error: %v", err),
			Data:    "",
		}

		fmt.Fprintf(w, GetApiMessage(resp))
		return
	}
	data := LoginResponse{
		Username: res.Username,
		Avatar:   res.Avatar,
	}
	jwtToken, err := utils.CreateJWTToken(res.UserId)
	cookie := http.Cookie{
		Path:    "/",
		Domain:  viper.GetString("server.host"),
		Name:    "jwt-token",
		Value:   jwtToken,
		Expires: time.Now().Add(2 * time.Hour),
	}

	http.SetCookie(w, &cookie)
	Data, err := utils.Marshal(data)
	if err != nil {
		Data = fmt.Sprintf("Error occurs when parse data. Error %v", err)
	}
	resp = ApiMessage{
		Err:     0,
		Message: "Login successfully !",
		Data:    Data,
	}
	fmt.Fprintf(w, GetApiMessage(resp))
}
