package GrpcBlogService

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	LibUtils "github.com/Lang0808/libutils"
)

var client BlogServiceClient

func InitConnection() error {
	host := LibUtils.Get("BlogService.host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Cannot initialize connection to UserService")
		return err
	}
	client = NewBlogServiceClient(conn)
	return nil
}

func CreatedBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	return client.CreatedBlog(ctx, in)
}

func MultiGetBlogs(ctx context.Context, in *MultiGetBlogsRequest, opts ...grpc.CallOption) (*MultiGetBlogsResponse, error) {
	return client.MultiGetBlogs(ctx, in)
}

func GetBlogsOfUser(ctx context.Context, in *GetBlogsOfUserRequest, opts ...grpc.CallOption) (*GetBlogsOfUserResponse, error) {
	return client.GetBlogsOfUser(ctx, in)
}
