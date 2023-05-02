package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func SendFriendRequest_Server(ctx context.Context, req *SendFriendRequestRequest) (*SendFriendRequestResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}
	isFollowing := false

	for i := 0; i < len(relationships); i = i + 1 {
		relationship := relationships[i]
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKED.String()] {
			return nil, errors.New("User from is blocked by user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKING.String()] {
			return nil, errors.New("User from is blocking user to. Please unblock before send friend request")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FRIEND.String()] {
			return nil, errors.New("User from is being a friend with user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_SENT_REQUEST.String()] {
			return nil, errors.New("User from has already sent friend request to user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_RECEIVED_REQUEST.String()] {
			return nil, errors.New("User to sent user from a friend request. Please accept it")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FOLLOWING.String()] {
			isFollowing = true
		}
	}
	_, err = AddRelationshipBetween2User(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_SENT_REQUEST)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when add send friend request relationship. Error: %v", err))
	}
	_, err = AddRelationshipBetween2User(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_RECEIVED_REQUEST)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when add receive friend request relationship. Error: %v", err))
	}
	// Automatically follow
	if !isFollowing {
		_, err = AddRelationshipBetween2User(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWING)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when add user from auto follow user to relationship. Error: %v", err))
		}
		_, err = AddRelationshipBetween2User(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWED)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when add user to auto be followed by user from relationship. Error: %v", err))
		}
	}
	return &SendFriendRequestResponse{}, nil
}
