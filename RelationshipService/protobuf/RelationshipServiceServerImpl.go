package GrpcRelationshipService

import (
	context "context"
)

type RelationshipServiceServerImpl struct {
}

func (r RelationshipServiceServerImpl) SendFriendRequest(ctx context.Context, req *SendFriendRequestRequest) (*SendFriendRequestResponse, error) {
	return SendFriendRequest_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) AcceptFriendRequest(ctx context.Context, req *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error) {
	return AcceptFriendRequest_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) RemoveFriendRequest(ctx context.Context, req *RemoveFriendRequestRequest) (*RemoveFriendRequestResponse, error) {
	return RemoveFriendRequest_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) BlockUser(ctx context.Context, req *BlockUserRequest) (*BlockUserResponse, error) {
	return BlockUser_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) UnblockUser(ctx context.Context, req *UnblockUserRequest) (*UnblockUserResponse, error) {
	return UnblockUser_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) FollowUser(ctx context.Context, req *FollowUserRequest) (*FollowUserResponse, error) {
	return FollowUser_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) UnfollowUser(ctx context.Context, req *UnfollowUserRequest) (*UnfollowUserResponse, error) {
	return UnfollowUser_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) UnfriendUser(ctx context.Context, req *UnfriendUserRequest) (*UnfriendUserResponse, error) {
	return UnfriendUser_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) GetRelationship(ctx context.Context, req *GetRelationshipRequest) (*GetRelationshipResponse, error) {
	return GetRelationship_Server(ctx, req)
}

func (r RelationshipServiceServerImpl) mustEmbedUnimplementedRelationshipServiceServer() {

}
