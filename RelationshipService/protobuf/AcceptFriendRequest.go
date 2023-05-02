package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func AcceptFriendRequest_Server(ctx context.Context, req *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}
	isReceivedRequest := false
	isFollowing := false
	for i := 0; i < len(relationships); i = i + 1 {
		relationship := relationships[i]
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKED.String()] {
			return nil, errors.New("User from is blocked by user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKING.String()] {
			return nil, errors.New("User from is blocking user to. Please unblock before accept friend request")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FRIEND.String()] {
			return nil, errors.New("User from is being a friend with user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_SENT_REQUEST.String()] {
			return nil, errors.New("User from has already sent friend request to user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_RECEIVED_REQUEST.String()] {
			isReceivedRequest = true
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FOLLOWING.String()] {
			isFollowing = true
		}
	}
	if isReceivedRequest == false {
		return nil, errors.New("User to didnt send friend request to user from")
	}
	_, err = AddRelationshipBetween2User(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FRIEND)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot add relationship friend between user from and user to. Error: %v", err))
	}
	_, err = AddRelationshipBetween2User(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FRIEND)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot add relationship friend between user to and user from. Error: %v", err))
	}
	_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_RECEIVED_REQUEST)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot remove relationship user from reveived friend request from user to. Error: %v", err))
	}
	_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_SENT_REQUEST)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot remove relationship user to sent friend request to user from. Error: %v", err))
	}
	if isFollowing == false {
		// user from follow user to automatically
		_, err = AddRelationshipBetween2User(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWING)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Cannot add relationship user from follows user to automatically. Error: %v", err))
		}
		_, err = AddRelationshipBetween2User(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWED)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Cannot add relationship user to is followed by user from automatically. Error: %v", err))
		}
	}
	return &AcceptFriendRequestResponse{}, nil
}
