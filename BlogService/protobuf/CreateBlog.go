package GrpcBlogService

import (
	"context"
	"errors"
	"fmt"
	"time"

	BlogDB "BlogService/db/BlogDB"
)

func validRequest(req *CreateBlogRequest) bool {
	if len(req.Title) == 0 {
		return false
	}
	if len(req.Content) == 0 {
		return false
	}
	if req.Author < 0 {
		return false
	}
	return true
}

func CreatedBlog_Server(ctx context.Context, req *CreateBlogRequest) (*CreateBlogResponse, error) {
	if !validRequest(req) {
		return nil, errors.New("Invalid request")
	}
	currentTime := time.Now().Unix()
	params := BlogDB.CreateBlogParams{
		Title:       req.Title,
		Privacy:     PrivacyMode_value[req.Privacy.String()],
		Createdtime: currentTime,
		Updatedtime: currentTime,
		Author:      req.Author,
		Content:     req.Content,
	}
	_, err := BlogDB.CreateBlog(ctx, params)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when save blog to db: %v", err))
	}
	return &CreateBlogResponse{}, nil
}
