package GrpcRelationshipService

import (
	"context"
	"errors"
	"fmt"
)

func GetRelationship_Server(ctx context.Context, req *GetRelationshipRequest) (*GetRelationshipResponse, error) {
	relationships, err := GetAllRelationshipsBetween2User(ctx, req.GetUserIdFrom(), req.GetUserIdTo())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot get all relationships between user from and user to. Error: %v", err))
	}
	var Relationships []RelationshipType
	for i := 0; i < len(relationships); i = i + 1 {
		relationship := relationships[i]
		Relationships = append(Relationships, RelationshipType(relationship.Relationship))
	}
	return &GetRelationshipResponse{Relationships: Relationships}, nil
}
