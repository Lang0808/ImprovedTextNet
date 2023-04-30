package main

import (
	"fmt"
	"log"
	"net"

	UserDB "UserService/db/UserDB"
	GrpcUserService "UserService/protobuf"
	utils "UserService/utils"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func initProject() (err error) {
	err = utils.InitConfig()
	if err != nil {
		fmt.Printf("Error when read config: %v", err)
		return err
	}
	return nil
}

func main() {

	err := initProject()
	if err != nil {
		return
	}
	err = UserDB.InitConnection()
	if err != nil {
		return
	}
	fmt.Printf("Initialization done\n")
	host := viper.GetString("server.host")
	lis, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Printf("Cannot listen to request. Error: %v\n", err)
	}

	s := grpc.NewServer()
	UserServiceServer := GrpcUserService.UserServiceServerImpl{}
	GrpcUserService.RegisterUserServiceServer(s, UserServiceServer)

	fmt.Printf("UserService is listening at %v\n", host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v\n", err)
	}

}
