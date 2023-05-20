package GrpcNewsFeedService

import (
	context "context"
	"errors"
)

type NewsFeedServiceServerImpl struct {
}

func (r NewsFeedServiceServerImpl) GetNewsFeed(context.Context, *GetNewsFeedRequest) (*GetNewsFeedResponse, error) {
	return nil, errors.New("Unsupported")
}

func (r NewsFeedServiceServerImpl) mustEmbedUnimplementedNewsFeedServiceServer() {

}
