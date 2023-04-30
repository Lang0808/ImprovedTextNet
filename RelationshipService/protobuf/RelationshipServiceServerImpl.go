package GrpcUserService

import (
	context "context"
	"errors"
)

type RelationshipServiceServerImpl struct {
}

func (r RelationshipServiceServerImpl) SendFriendRequest(context.Context, *SendFriendRequestRequest) (*SendFriendRequestResponse, error) {
	return nil, errors.New("Unsupported.")
}

func (r RelationshipServiceServerImpl) AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error) {
	return nil, errors.New("Unsupported.")
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
	return nil, errors.New("Unsupported.")
}
