package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func RemoveFriendRequest_Server(ctx context.Context, req *RemoveFriendRequestRequest) (*RemoveFriendRequestResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}
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
			return nil, errors.New("User from is being a friend with user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_SENT_REQUEST.String()] {
			return nil, errors.New("User from sent friend request to user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_RECEIVED_REQUEST.String()] {
			isReceivedRequest = true
		}
	}
	if isReceivedRequest == false {
		return nil, errors.New("User to didnt send any friend request to user from")
	}
	_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_RECEIVED_REQUEST)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from received friend request from user to. Error: %v", err))
	}
	_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_SENT_REQUEST)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to sent friend request to user from. Error: %v", err))
	}
	return &RemoveFriendRequestResponse{}, nil
}
