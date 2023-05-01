package main

import (
	"context"
	"fmt"

	RelationshipDB "RelationshipService/db/RelationshipDB"
	GrpcRelationshipService "RelationshipService/protobuf"
	utils "RelationshipService/utils"
	// "google.golang.org/grpc"
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
	err = RelationshipDB.InitConnection()
	if err != nil {
		return
	}
	fmt.Printf("Initialization done\n")
	in2 := &GrpcRelationshipService.SendFriendRequestRequest{
		UserIdFrom: 1,
		UserIdTo:   2,
	}
	resp2, err := GrpcRelationshipService.SendFriendRequest_Server(context.Background(), in2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Response: %v\n", resp2)
	}
	in := &GrpcRelationshipService.AcceptFriendRequestRequest{
		UserIdFrom: 2,
		UserIdTo:   1,
	}
	resp, err := GrpcRelationshipService.AcceptFriendRequest_Server(context.Background(), in)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Response: %v\n", resp)
	}
	// host := viper.GetString("server.host")
	// lis, err := net.Listen("tcp", host)
	// if err != nil {
	// 	fmt.Printf("Cannot listen to request. Error: %v\n", err)
	// }

	// s := grpc.NewServer()
	// UserServiceServer := GrpcRelationshipService.UserServiceServerImpl{}
	// GrpcRelationshipService.RegisterUserServiceServer(s, UserServiceServer)

	// fmt.Printf("RelationshipService is listening at %v\n", host)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to start server %v\n", err)
	// }

}
