package handler

import (
	GrpcRelationshipService "WebServer/clients/RelationshipService"
	utils "WebServer/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type BlockUserModel struct {
}

func (b BlockUserModel) GetSrcId(r *http.Request) (int32, error) {
	return DefaultGetSrcId(r)
}

func (b BlockUserModel) GetCommandId() int32 {
	return int32(BLOCK_USER)
}

func (b BlockUserModel) Handle(srcId int32, w http.ResponseWriter, r *http.Request,
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
	request := &GrpcRelationshipService.BlockUserRequest{
		UserIdFrom: srcId,
		UserIdTo:   UserIdTo,
	}
	_, err = GrpcRelationshipService.BlockUser(context.Background(), request)
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

func (b BlockUserModel) PostProcess(w *http.ResponseWriter, r *http.Request) {

}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	requestHandler := RequestHandler{
		inner: BlockUserModel{},
	}
	requestHandler.process(w, r)
}
