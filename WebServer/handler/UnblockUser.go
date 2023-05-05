package handler

import (
	GrpcRelationshipService "WebServer/clients/RelationshipService"
	utils "WebServer/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func UnblockUser(w http.ResponseWriter, r *http.Request) {
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
	UserIdTo64, err := strconv.ParseInt(r.FormValue("UserId"), 10, 64)
	if err != nil {
		resp := ApiMessage{
			Err:     -2,
			Message: fmt.Sprintf("Please check field UserId. Error: %v", err),
			Data:    "",
		}
		str := GetApiMessage(resp)
		fmt.Fprintf(w, str)
		return
	}
	UserIdTo, err := utils.DenoiseUserId(UserIdTo64)
	if err != nil {
		resp := ApiMessage{
			Err:     -4,
			Message: fmt.Sprintf("%v", err),
			Data:    "",
		}
		str := GetApiMessage(resp)
		fmt.Fprintf(w, str)
		return
	}

	request := &GrpcRelationshipService.UnblockUserRequest{
		UserIdFrom: UserIdFrom,
		UserIdTo:   UserIdTo,
	}
	_, err = GrpcRelationshipService.UnblockUser(context.Background(), request)
	if err != nil {
		resp := ApiMessage{
			Err:     -1,
			Message: fmt.Sprintf("Error when call to RelationshipService. Error: %v", err),
			Data:    "",
		}
		str := GetApiMessage(resp)
		fmt.Fprintf(w, str)
		return
	}
	fmt.Fprintf(w, GetApiMessage(SUCCESS_MSG))
}
