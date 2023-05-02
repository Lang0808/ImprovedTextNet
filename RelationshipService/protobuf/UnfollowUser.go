package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func UnfollowUser_Server(ctx context.Context, req *UnfollowUserRequest) (*UnfollowUserResponse, error) {
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
			return nil, errors.New("User from is blocking user to. Please unblock before follow")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FOLLOWING.String()] {
			return nil, errors.New("User from has already followed user to")
		}
	}
	if isFollowing == false {
		return nil, errors.New("User from didnt follow user to")
	}
	_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWING)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from is following user to. Error: %v", err))
	}
	_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWED)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to is followed by user from. Error: %v", err))
	}
	return &UnfollowUserResponse{}, nil
}
