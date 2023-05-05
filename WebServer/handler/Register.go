package handler

import (
	GrpcUserService "WebServer/clients/UserService"
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("CORS server: %v\n", viper.GetString("client.domain"))
	fmt.Printf("CORS client: %v\n", r.Host)

	w.Header().Set("Access-Control-Allow-Origin", viper.GetString("client.domain"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

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
