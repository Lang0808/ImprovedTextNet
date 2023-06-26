package handler

import (
	GrpcRelationshipService "WebServer/clients/RelationshipService"
	utils "WebServer/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type UnfriendUserModel struct {
}

func (u UnfriendUserModel) GetSrcId(r *http.Request) (int32, error) {
	return DefaultGetSrcId(r)
}

func (u UnfriendUserModel) GetCommandId() int32 {
	return int32(UNFRIEND_USER)
}

func (u UnfriendUserModel) Handle(srcId int32, w http.ResponseWriter, r *http.Request,
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
	request := &GrpcRelationshipService.UnfriendUserRequest{
		UserIdFrom: srcId,
		UserIdTo:   UserIdTo,
	}
	_, err = GrpcRelationshipService.UnfriendUser(context.Background(), request)
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

func (u UnfriendUserModel) PostProcess(w *http.ResponseWriter, r *http.Request) {
	DefaultPostProcess(w, r)
}

func UnfriendUser(w http.ResponseWriter, r *http.Request) {
	requestHandler := RequestHandler{
		inner: UnfriendUserModel{},
	}
	requestHandler.process(w, r)
}
