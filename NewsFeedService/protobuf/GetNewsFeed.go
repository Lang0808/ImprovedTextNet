package GrpcNewsFeedService

import (
	"context"
	"errors"
	"fmt"

	NewsFeedDB "NewsFeedService/db/NewsFeedDB"
)

func GetNewsFeed_Server(ctx context.Context, req *GetNewsFeedRequest) (*GetNewsFeedResponse, error) {
	params := NewsFeedDB.GetNewsFeedParams{
		Userid: req.UserId,
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	listNewsFeed, err := NewsFeedDB.GetNewsFeed(ctx, params)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when get newsfeed: %v", err))
	}
	var listBlogId []int32
	for i := 0; i < len(listNewsFeed); i = i + 1 {
		listBlogId = append(listBlogId, listNewsFeed[i].Blogid)
	}
	return &GetNewsFeedResponse{BlogId: listBlogId}, nil

}
