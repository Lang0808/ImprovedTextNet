package GrpcRelationshipService

import (
	RelationshipDB "RelationshipService/db/RelationshipDB"
	"context"
	"errors"
	"fmt"
)

func GetAllUserTo_Server(ctx context.Context, req *GetAllUserToRequest) (*GetAllUserToResponse, error) {
	// step 1: call db get all user to
	params := RelationshipDB.GetAllUserToParams{
		UseridFrom:   req.UserFromId,
		Relationship: RelationshipType_value[req.Relationship.String()],
	}
	listRelationships, err := RelationshipDB.GetAllUserTo(ctx, params)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when call db: %v", err))
	}
	// step 2: convert answer from db and return
	var listUserTo []int32
	for i := 0; i < len(listRelationships); i = i + 1 {
		listUserTo = append(listUserTo, listRelationships[i].Useridto)
	}
	return &GetAllUserToResponse{ListUserTo: listUserTo}, nil
}
