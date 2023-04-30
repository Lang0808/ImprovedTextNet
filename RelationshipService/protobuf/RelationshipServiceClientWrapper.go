package GrpcUserService

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var client RelationshipServiceClient

func InitConnection() error {
	host := viper.GetString("RelationshipService.host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Cannot initialize connection to RelationshipService")
		return err
	}
	client = NewRelationshipServiceClient(conn)
	return nil
}

func SendFriendRequest(ctx context.Context, in *SendFriendRequestRequest, opts ...grpc.CallOption) (*SendFriendRequestResponse, error) {
	return client.SendFriendRequest(ctx, in, opts)
}

func AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*AcceptFriendRequestResponse, error) {
	return client.AcceptFriendRequest(ctx, in, opts)
}

func RemoveFriendRequest(ctx context.Context, in *RemoveFriendRequestRequest, opts ...grpc.CallOption) (*RemoveFriendRequestResponse, error) {
	return client.RemoveFriendRequest(ctx, in, opts)
}

func BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error) {
	return client.BlockUser(ctx, in, opts)
}

func UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*UnblockUserResponse, error) {
	return client.UnblockUser(ctx, in, opts)
}

func FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponse, error) {
	return client.FollowUser(ctx, in, opts)
}

func UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*UnfollowUserResponse, error) {
	return client.UnfollowUser(ctx, in, opts)
}
