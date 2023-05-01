package GrpcRelationshipService

import (
	context "context"
	"errors"
)

type RelationshipServiceServerImpl struct {
}

func (r RelationshipServiceServerImpl) SendFriendRequest(ctx context.Context, req *SendFriendRequestRequest) (*SendFriendRequestResponse, error) {
	return SendFriendRequest_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) AcceptFriendRequest(ctx context.Context, req *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error) {
	return AcceptFriendRequest_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) RemoveFriendRequest(context.Context, *RemoveFriendRequestRequest) (*RemoveFriendRequestResponse, error) {
	return nil, errors.New("Unsupported.")
}

func (r RelationshipServiceServerImpl) BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponse, error) {
	return nil, errors.New("Unsupported.")
}

func (r RelationshipServiceServerImpl) UnblockUser(context.Context, *UnblockUserRequest) (*UnblockUserResponse, error) {
	return nil, errors.New("Unsupported.")
}

func (r RelationshipServiceServerImpl) FollowUser(context.Context, *FollowUserRequest) (*FollowUserResponse, error) {
	return nil, errors.New("Unsupported.")
}

func (r RelationshipServiceServerImpl) UnfollowUser(context.Context, *UnfollowUserRequest) (*UnfollowUserResponse, error) {
	return nil, errors.New("Unsupported.")
}

func (r RelationshipServiceServerImpl) mustEmbedUnimplementedRelationshipServiceServer() {

}
