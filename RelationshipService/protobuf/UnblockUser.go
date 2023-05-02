package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func UnblockUser_Server(ctx context.Context, req *UnblockUserRequest) (*UnblockUserResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}
	isBlocking := false

	for i := 0; i < len(relationships); i = i + 1 {
		relationship := relationships[i]
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKED.String()] {
			return nil, errors.New("User from is blocked by user to")
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_BLOCKING.String()] {
			isBlocking = true
		}
	}
	if isBlocking == false {
		return nil, errors.New("User from didnt block user to")
	}
	_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_BLOCKING)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from is blocking user to. Error: %v", err))
	}
	_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_BLOCKED)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to is blocked by user from. Error: %v", err))
	}
	return &UnblockUserResponse{}, nil
}
