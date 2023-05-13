package handler

import (
	GrpcBlogService "WebServer/clients/BlogService"
	"context"

	utils "WebServer/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

func ValidateCreateBlogRequest(req GrpcBlogService.CreateBlogRequest) bool {
	if len(req.Title) == 0 {
		return false
	}
	if len(req.Content) == 0 {
		return false
	}
	return true
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", viper.GetString("client.domain"))
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")

	AuthCookie, err := r.Cookie("jwt-token")
	if err != nil {
		fmt.Fprintf(w, GetApiMessage(UNAUTHORIZE_MSG))
		return
	}
	UserIdFrom, err := utils.GetUserIdInJWTToken(AuthCookie.Value)
	if err != nil {
		fmt.Fprintf(w, GetApiMessage(UNAUTHORIZE_MSG))
		return
	}
	PrivacyMode64, err := strconv.ParseInt(r.FormValue("PrivacyMode"), 10, 32)
	if err != nil {
		msg_resp := ApiMessage{
			Err:     -1,
			Message: "Cannot parse Privacy mode",
			Data:    "",
		}
		fmt.Fprintf(w, GetApiMessage(msg_resp))
		return
	}
	PrivacyMode32 := int32(PrivacyMode64)
	req := GrpcBlogService.CreateBlogRequest{
		Title:   r.FormValue("Title"),
		Privacy: GrpcBlogService.PrivacyMode(PrivacyMode32),
		Author:  UserIdFrom,
		Content: r.FormValue("Content"),
	}
	ok := ValidateCreateBlogRequest(req)
	if !ok {
		msg_resp := ApiMessage{
			Err:     -2,
			Message: "Input is invalid",
			Data:    "",
		}
		fmt.Fprintf(w, GetApiMessage(msg_resp))
		return
	}
	_, err = GrpcBlogService.CreatedBlog(context.Background(), &req)
	if err != nil {
		msg_resp := ApiMessage{
			Err:     -3,
			Message: fmt.Sprintf("Error when call blog service. Error: %v", err),
			Data:    "",
		}
		fmt.Fprintf(w, GetApiMessage(msg_resp))
		return
	}
	fmt.Fprintf(w, GetApiMessage(SUCCESS_MSG))
}
