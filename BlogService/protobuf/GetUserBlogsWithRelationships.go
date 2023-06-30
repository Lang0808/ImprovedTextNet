package GrpcBlogService

import (
	"context"
	"errors"
)

type GetUserBlogsWithRelationshipsModel struct {
}

func GetUserBlogsWithRelationships_Server(ctx context.Context,
	req *GetUserBlogsWithRelationshipsRequest, extra *map[string]string) (*GetBlogsOfUserResponse, error) {
	return nil, errors.New("Unsupported")
}
