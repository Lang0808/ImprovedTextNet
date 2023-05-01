package GrpcRelationshipService

import (
	RelationshipDB "RelationshipService/db/RelationshipDB"
	"context"
	"database/sql"
	"fmt"
)

func GetAllRelationshipsBetween2User(ctx context.Context, UserIdFrom int32, UserIdTo int32) ([]RelationshipDB.Relationship, error) {
	params := RelationshipDB.GetAllRelationshipsBetween2UsersParams{
		Useridfrom: UserIdFrom,
		Useridto:   UserIdTo,
	}
	relationships, err := RelationshipDB.GetAllRelationshipsBetween2Users(ctx, params)
	if err != nil {
		fmt.Printf("RelationshipDB.GetAllRelationshipsBetween2Users fail. Error: %v\n", err)
		return nil, err
	}
	return relationships, err
}

func AddRelationshipBetween2User(ctx context.Context, UserIdFrom int32, UserIdTo int32, Relationship RelationshipType) (sql.Result, error) {
	params := RelationshipDB.AddRelationshipParams{
		Useridfrom:   UserIdFrom,
		Useridto:     UserIdTo,
		Relationship: RelationshipType_value[Relationship.String()],
	}
	res, err := RelationshipDB.AddRelationship(ctx, params)
	if err != nil {
		fmt.Printf("RelationshipDB.AddRelationship fail. Error: %v\n", err)
		return nil, err
	}
	return res, err
}

func RemoveRelationship(ctx context.Context, UserIdFrom int32, UserIdTo int32, Relationship RelationshipType) (sql.Result, error) {
	params := RelationshipDB.RemoveRelationshipParams{
		Useridfrom:   UserIdFrom,
		Useridto:     UserIdTo,
		Relationship: RelationshipType_value[Relationship.String()],
	}
	res, err := RelationshipDB.RemoveRelationship(ctx, params)
	if err != nil {
		fmt.Printf("RelationshipDB.RemoveRelationship fail. Error: %v\n", err)
		return nil, err
	}
	return res, err
}
