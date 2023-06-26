package handler

import (
	GrpcUserService "WebServer/clients/UserService"
	utils "WebServer/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	LibUtils "github.com/Lang0808/libutils"
)

type LoginResponse struct {
	Username string
	Avatar   string
}

type LoginUserModel struct {
}

func (l LoginUserModel) GetSrcId(r *http.Request) (int32, error) {
	return -1, nil
}

func (l LoginUserModel) GetCommandId() int32 {
	return int32(LOGIN)
}

func (l LoginUserModel) Handle(srcId int32, w http.ResponseWriter, r *http.Request,
	extra *map[string]string) {
	// Get username and password
	Username := r.FormValue("Username")
	Password := r.FormValue("Password")
	in := &GrpcUserService.LoginUserRequest{
		Username: Username,
		Password: Password,
	}
	// call UserService to verify username and password
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
	// create cookie JWT and set cookie
	jwtToken, err := utils.CreateJWTToken(res.UserId)
	cookie := http.Cookie{
		Path:    "/",
		Name:    "jwt-token",
		Domain:  LibUtils.Get("client.domainCookie"),
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

func (l LoginUserModel) PostProcess(w *http.ResponseWriter, r *http.Request) {
	DefaultPostProcess(w, r)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	requestHandler := RequestHandler{
		inner: LoginUserModel{},
	}
	requestHandler.process(w, r)
}
