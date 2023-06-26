package handler

import (
	GrpcBlogService "WebServer/clients/BlogService"
	"context"

	"fmt"
	"net/http"
	"strconv"
)

type CreateBlogModel struct {
}

func (c CreateBlogModel) GetSrcId(r *http.Request) (int32, error) {
	return DefaultGetSrcId(r)
}

func (c CreateBlogModel) GetCommandId() int32 {
	return int32(CREATE_BLOG)
}

func (c CreateBlogModel) Handle(srcId int32, w http.ResponseWriter, r *http.Request,
	extra *map[string]string) {
	// Get privacy mode of post
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
	(*extra)["PrivacyMode"] = string(PrivacyMode32)
	// call blog service to store blog
	req := GrpcBlogService.CreateBlogRequest{
		Title:   r.FormValue("Title"),
		Privacy: GrpcBlogService.PrivacyMode(PrivacyMode32),
		Author:  srcId,
		Content: r.FormValue("Content"),
	}
	ok := ValidateCreateBlogRequest(&req)
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

func (c CreateBlogModel) PostProcess(w *http.ResponseWriter, r *http.Request) {
	DefaultPostProcess(w, r)
}

func ValidateCreateBlogRequest(req *GrpcBlogService.CreateBlogRequest) bool {
	if len(req.Title) == 0 {
		return false
	}
	if len(req.Content) == 0 {
		return false
	}
	return true
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	requestHandler := RequestHandler{
		inner: CreateBlogModel{},
	}
	requestHandler.process(w, r)
}
