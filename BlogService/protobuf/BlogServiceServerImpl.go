package GrpcBlogService

import (
	context "context"
	"errors"
	"time"

	LibUtils "github.com/Lang0808/libutils"
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

func (u BlogServiceServerImpl) GetUserBlogsWithRelationships(ctx context.Context, req *GetUserBlogsWithRelationshipsRequest) (*GetBlogsOfUserResponse, error) {
	logEntry := GetLogEntry(int32(GET_USER_BLOGS_WITH_RELATIONSHIPS))
	start := time.Now().UnixMilli()

	resp, err := GetUserBlogsWithRelationships_Server(ctx, req, &logEntry.Extra)

	end := time.Now().UnixMilli()
	logEntry.ExecTime = int32(end - start)
	LibUtils.Info_entry("BlogServiceServerImpl", "GetUserBlogsWithRelationships", logEntry)
	return resp, err
}

func (u BlogServiceServerImpl) mustEmbedUnimplementedBlogServiceServer() {

}

func GetLogEntry(commandId int32) LibUtils.LogEntry {
	logEntry := LibUtils.LogEntry{}
	logEntry.CommandId = commandId
	logEntry.SrcId = -1
	logEntry.Extra = make(map[string]string)
	return logEntry
}
