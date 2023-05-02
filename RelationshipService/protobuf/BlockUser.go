package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func BlockUser_Server(ctx context.Context, req *BlockUserRequest) (*BlockUserResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}
	isFollowing := false
	isFollowed := false
	isFriend := false
	isSentRequest := false
	isReceivedRequest := false
	for i := 0; i < len(relationships); i = i + 1 {
		relationship := relationships[i]
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKED.String()] {
			return nil, errors.New("User from is blocked by user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKING.String()] {
			return nil, errors.New("User from is blocking user to. Please unblock before send friend request")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FRIEND.String()] {
			isFriend = true
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_SENT_REQUEST.String()] {
			isSentRequest = true
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_RECEIVED_REQUEST.String()] {
			isReceivedRequest = true
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FOLLOWING.String()] {
			isFollowing = true
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FOLLOWED.String()] {
			isFollowed = true
		}
	}
	_, err = AddRelationshipBetween2User(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_BLOCKING)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when add relationship user from blocks user to. Error: %v", err))
	}
	_, err = AddRelationshipBetween2User(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_BLOCKED)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when add relationship user to is blocked by user from. Error: %v", err))
	}
	if isFollowing == true {
		_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWING)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from follows user to. Error: %v", err))
		}
		_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWED)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to is followed by user from. Error: %v", err))
		}
	}
	if isFollowed == true {
		_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWED)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from is followed by user to. Error: %v", err))
		}
		_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWING)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to follows user from. Error: %v", err))
		}
	}
	if isFriend == true {
		_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FRIEND)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship friend between user from and user to. Error: %v", err))
		}
		_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FRIEND)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship friend between user to and user from. Error: %v", err))
		}
	}
	if isSentRequest == true {
		_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_SENT_REQUEST)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from sent friend request to user to. Error: %v", err))
		}
		_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_RECEIVED_REQUEST)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to received friend request from user from. Error: %v", err))
		}
	}
	if isReceivedRequest == true {
		_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_RECEIVED_REQUEST)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from received friend request from user to. Error: %v", err))
		}
		_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_SENT_REQUEST)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to sent friend request to user from. Error: %v", err))
		}
	}
	return &BlockUserResponse{}, nil
}
