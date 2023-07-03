package GrpcBlogService

import (
	"context"
	"errors"
)

func GetUserBlogsWithRelationships_Server(ctx context.Context,
	req *GetUserBlogsWithRelationshipsRequest, extra *map[string]string) (*GetBlogsOfUserResponse, error) {
	for relationship := range req.Relationship {

	}

	return nil, errors.New("Unsupported")
}
