package handler

import (
	GrpcRelationshipService "WebServer/clients/RelationshipService"
	utils "WebServer/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type SendFriendRequestModel struct {
}

func (s SendFriendRequestModel) GetSrcId(r *http.Request) (int32, error) {
	return DefaultGetSrcId(r)
}

func (s SendFriendRequestModel) GetCommandId() int32 {
	return int32(SEND_FRIEND_REQUEST)
}

func (s SendFriendRequestModel) Handle(srcId int32, w http.ResponseWriter, r *http.Request,
	extra *map[string]string) {
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
	request := &GrpcRelationshipService.SendFriendRequestRequest{
		UserIdFrom: srcId,
		UserIdTo:   UserIdTo,
	}
	_, err = GrpcRelationshipService.SendFriendRequest(context.Background(), request)
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

func (s SendFriendRequestModel) PostProcess(w *http.ResponseWriter, r *http.Request) {
	DefaultPostProcess(w, r)
}

func SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	requestHandler := RequestHandler{
		inner: SendFriendRequestModel{},
	}
	requestHandler.process(w, r)
}
