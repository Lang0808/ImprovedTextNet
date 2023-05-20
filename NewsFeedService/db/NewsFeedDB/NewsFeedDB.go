package BlogDB

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

func AddNewsFeed(ctx context.Context, arg AddNewsFeedParams) (sql.Result, error) {
	queries := New(db)
	return queries.AddNewsFeed(ctx, arg)
}

func GetNewsFeed(ctx context.Context, arg GetNewsFeedParams) (sql.Result, error) {
	queries := New(db)
	return queries.GetNewsFeed(ctx, arg)
}
