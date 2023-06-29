package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type GetUserBlogsModel struct {
}

type GetUserBlogsParams struct {
	UserId int32
	Count  int32
	Limit  int32
}

func (g GetUserBlogsModel) GetSrcId(r *http.Request) (int32, error) {
	return DefaultGetSrcId(r)
}

func (g GetUserBlogsModel) GetCommandId() int32 {
	return int32(GET_USER_BLOGS)
}

func GetParams(r *http.Request) (*GetUserBlogsParams, error) {
	Count, err := strconv.ParseInt(r.FormValue("Count"), 10, 32)
	if err != nil {
		return nil, errors.New("Invalid Count")
	}
	Limit, err := strconv.ParseInt(r.FormValue("Limit"), 10, 32)
	if err != nil {
		return nil, errors.New("Invalid Limit")
	}
	UserId, err := GetUserId(r.FormValue("UserId"))
	if err != nil {
		return nil, errors.New("Invalid UserId")
	}

	params := GetUserBlogsParams{
		Count:  int32(Count),
		Limit:  int32(Limit),
		UserId: UserId,
	}
	return &params, nil
}

func (g GetUserBlogsModel) Handle(srcId int32, w http.ResponseWriter, r *http.Request,
	extra *map[string]string) {
	params, err := GetParams(r)
	if err != nil {
		msg := GetApiMessageWithErrorCode(int32(INVALID_PARAMS), err)
		fmt.Fprintf(w, msg)
		return
	}

}

func (g GetUserBlogsModel) PostProcess(w *http.ResponseWriter, r *http.Request) {
	DefaultPostProcess(w, r)
}

func GetUserBlogs(w http.ResponseWriter, r *http.Request) {
	requestHandler := RequestHandler{
		inner: GetUserBlogsModel{},
	}
	requestHandler.process(w, r)
}
