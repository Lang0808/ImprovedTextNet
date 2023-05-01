package RelationshipDB

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitConnection() (err error) {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.dbName")
	protocol := viper.GetString("mysql.protocol")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	connection := fmt.Sprintf("%v:%v@%v(%v:%v)/%v", username, password, protocol, host, port, dbName)
	fmt.Printf("Connecting to database %v\n", connection)
	db, err = sql.Open("mysql", connection)
	if err != nil {
		fmt.Printf("Connect to db fail: %v\n", err)
	}
	return err
}

func AddRelationship(ctx context.Context, in AddRelationshipParams) (sql.Result, error) {
	queries := New(db)
	return queries.AddRelationship(ctx, in)
}

func GetAllRelationshipsBetween2Users(ctx context.Context, in GetAllRelationshipsBetween2UsersParams) ([]Relationship, error) {
	queries := New(db)
	return queries.GetAllRelationshipsBetween2Users(ctx, in)
}

func GetAllUserFrom(ctx context.Context, in GetAllUserFromParams) ([]Relationship, error) {
	queries := New(db)
	return queries.GetAllUserFrom(ctx, in)
}

func GetAllUserTo(ctx context.Context, in GetAllUserToParams) ([]Relationship, error) {
	queries := New(db)
	return queries.GetAllUserTo(ctx, in)
}

func RemoveRelationship(ctx context.Context, in RemoveRelationshipParams) (sql.Result, error) {
	queries := New(db)
	return queries.RemoveRelationship(ctx, in)
}
