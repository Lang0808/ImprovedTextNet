package GrpcNewsFeedService

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var client NewsFeedServiceClient

func InitConnection() error {
	host := viper.GetString("NewsFeedService.host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Cannot initialize connection to NewsFeedService")
		return err
	}
	client = NewNewsFeedServiceClient(conn)
	return nil
}

func SendFriendRequest(ctx context.Context, in *SendFriendRequestRequest, opts ...grpc.CallOption) (*SendFriendRequestResponse, error) {
	return client.SendFriendRequest(ctx, in)
}

func AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*AcceptFriendRequestResponse, error) {
	return client.AcceptFriendRequest(ctx, in)
}

func RemoveFriendRequest(ctx context.Context, in *RemoveFriendRequestRequest, opts ...grpc.CallOption) (*RemoveFriendRequestResponse, error) {
	return client.RemoveFriendRequest(ctx, in)
}

func BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error) {
	return client.BlockUser(ctx, in)
}

func UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*UnblockUserResponse, error) {
	return client.UnblockUser(ctx, in)
}

func FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponse, error) {
	return client.FollowUser(ctx, in)
}

func UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*UnfollowUserResponse, error) {
	return client.UnfollowUser(ctx, in)
}

func UnfriendUser(ctx context.Context, in *UnfriendUserRequest, opts ...grpc.CallOption) (*UnfriendUserResponse, error) {
	return client.UnfriendUser(ctx, in)
}

func GetRelationship(ctx context.Context, in *GetRelationshipRequest, opts ...grpc.CallOption) (*GetRelationshipResponse, error) {
	return client.GetRelationship(ctx, in)
}

func GetAllUserTo(ctx context.Context, in *GetAllUserToRequest, opts ...grpc.CallOption) (*GetAllUserToResponse, error) {
	return client.GetAllUserTo(ctx, in)
}
