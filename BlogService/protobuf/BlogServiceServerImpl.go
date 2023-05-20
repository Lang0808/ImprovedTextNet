package GrpcBlogService

import (
	context "context"
	"errors"
)

type BlogServiceServerImpl struct {
}

func (u BlogServiceServerImpl) CreatedBlog(ctx context.Context, req *CreateBlogRequest) (*CreateBlogResponse, error) {
	return CreatedBlog_Server(ctx, req)
}

func (u BlogServiceServerImpl) MultiGetBlogs(context.Context, *MultiGetBlogsRequest) (*MultiGetBlogsResponse, error) {
	return nil, errors.New("Unsupported")
}

func (u BlogServiceServerImpl) GetBlogsOfUser(context.Context, *GetBlogsOfUserRequest) (*GetBlogsOfUserResponse, error) {
	return nil, errors.New("Unsupported")
}

func (u BlogServiceServerImpl) mustEmbedUnimplementedBlogServiceServer() {

}
