package handler

import (
	GrpcRelationshipService "WebServer/clients/RelationshipService"
	GrpcUserService "WebServer/clients/UserService"
	utils "WebServer/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type Profile struct {
	Relationships []string
	Username      string
	Avatar        string
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
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

	request := &GrpcRelationshipService.GetRelationshipRequest{
		UserIdFrom: UserIdFrom,
		UserIdTo:   UserIdTo,
	}
	relationshipsResp, err := GrpcRelationshipService.GetRelationship(context.Background(), request)
	if err != nil {
		resp := ApiMessage{
			Err:     -1,
			Message: fmt.Sprintf("Error when get relationships. Error: %v", err),
			Data:    "",
		}
		str := GetApiMessage(resp)
		fmt.Fprintf(w, str)
		return
	}

	var Relationships []string
	for i := 0; i < len(relationshipsResp.Relationships); i = i + 1 {
		relationship := relationshipsResp.Relationships[i]
		if relationship == GrpcRelationshipService.RelationshipType_BLOCKED {
			resp := ApiMessage{
				Err:     -2,
				Message: fmt.Sprintf("User from is blocked by user to."),
				Data:    "",
			}
			str := GetApiMessage(resp)
			fmt.Fprintf(w, str)
			return
		}
		Relationships = append(Relationships, relationship.String())
	}

	req := &GrpcUserService.GetUserInfoRequest{
		UserId: UserIdTo,
	}
	GetUserInfoResp, err := GrpcUserService.GetUserInfo(context.Background(), req)
	if err != nil {
		resp := ApiMessage{
			Err:     -3,
			Message: fmt.Sprintf("Error when get user info. Error: %v", err),
			Data:    "",
		}
		str := GetApiMessage(resp)
		fmt.Fprintf(w, str)
		return
	}

	Profile := Profile{
		Relationships: Relationships,
		Username:      GetUserInfoResp.Username,
		Avatar:        GetUserInfoResp.Avatar,
	}
	Data, err := utils.Marshal(Profile)
	if err != nil {
		Data = fmt.Sprintf("Error occurs when parse data. Error %v", err)
	}
	resp := ApiMessage{
		Err:     0,
		Message: "Success",
		Data:    Data,
	}
	fmt.Fprintf(w, GetApiMessage(resp))
}
