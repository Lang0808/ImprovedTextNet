package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func UnfriendUser_Server(ctx context.Context, req *UnfriendUserRequest) (*UnfriendUserResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}
	isFriend := false
	isFollowing := false
	isFollowed := false
	for i := 0; i < len(relationships); i = i + 1 {
		relationship := relationships[i]
		if relationship.Relationship == RelationshipType_value[RelationshipType_FRIEND.String()] {
			isFriend = true
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FOLLOWING.String()] {
			isFollowing = true
		}
		if relationship.Relationship == RelationshipType_value[RelationshipType_FOLLOWED.String()] {
			isFollowed = true
		}
	}
	if isFriend == false {
		return nil, errors.New("User from and user to is not friend")
	}
	_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FRIEND)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from is friend with user to. Error: %v", err))
	}
	_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FRIEND)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to is friend with user from. Error: %v", err))
	}
	if isFollowing == false {
		_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWING)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from is following user to. Error: %v", err))
		}
		_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWED)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to is followed by user from. Error: %v", err))
		}
	}
	if isFollowed == false {
		_, err = RemoveRelationship(ctx, req.UserIdFrom, req.UserIdTo, RelationshipType_FOLLOWED)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user from is followed by user to. Error: %v", err))
		}
		_, err = RemoveRelationship(ctx, req.UserIdTo, req.UserIdFrom, RelationshipType_FOLLOWING)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occurs when remove relationship user to is following user from. Error: %v", err))
		}
	}
	return &UnfriendUserResponse{}, nil
}
