package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func FollowUser_Server(ctx context.Context, req *FollowUserRequest) (*FollowUserResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}

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
	_, err = AddRelationshipBetween2User(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWING)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when add relationship user from is following user to. Error: %v", err))
	}
	_, err = AddRelationshipBetween2User(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWED)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when add relationship user to is followed by user from. Error: %v", err))
	}
	return &FollowUserResponse{}, nil
}
