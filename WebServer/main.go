package main

import (
	"WebServer/utils"
	"fmt"
	"net/http"

	GrpcUserService "WebServer/clients/UserService"

	GrpcRelationshipService "WebServer/clients/RelationshipService"

	GrpcBlogService "WebServer/clients/BlogService"

	handler "WebServer/handler"

	"github.com/spf13/viper"
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
	fmt.Printf("Inititialize config done\n")
	err = GrpcUserService.InitConnection()
	if err != nil {
		return
	}
	fmt.Printf("Initialize connection to UserService done\n")
	err = GrpcRelationshipService.InitConnection()
	if err != nil {
		return
	}
	fmt.Printf("Initialize connection to RelationshipService done\n")
	err = GrpcBlogService.InitConnection()
	if err != nil {
		return
	}
	fmt.Printf("Initialize connection to BlogService done\n")
	NoisedUserId := utils.NoiseUserId(1)
	fmt.Printf("Noise user id 1: %v\n", NoisedUserId)
	DenoisedUserId, err := utils.DenoiseUserId(NoisedUserId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("DenoisedUserId: %v\n", DenoisedUserId)
	}
	NoisedUserId = utils.NoiseUserId(3)
	fmt.Printf("Noise user id 3: %v\n", NoisedUserId)
	DenoisedUserId, err = utils.DenoiseUserId(NoisedUserId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("DenoisedUserId: %v\n", DenoisedUserId)
	}
	fmt.Printf("Initialize connection to RelationshipService done\n")
	http.HandleFunc("/api/RegisterUser", handler.RegisterUser)
	http.HandleFunc("/api/LoginUser", handler.LoginUser)
	http.HandleFunc("/api/SendFriendRequest", handler.SendFriendRequest)
	http.HandleFunc("/api/GetProfile", handler.GetProfile)
	http.HandleFunc("/api/AcceptFriendRequest", handler.AcceptFriendRequest)
	http.HandleFunc("/api/UnfriendUser", handler.UnfriendUser)
	http.HandleFunc("/api/UnfollowUser", handler.UnfollowUser)
	http.HandleFunc("/api/FollowUser", handler.FollowUser)
	http.HandleFunc("/api/BlockUser", handler.BlockUser)
	http.HandleFunc("/api/UnblockUser", handler.UnblockUser)
	http.HandleFunc("/api/CreateBlog", handler.CreateBlog)
	host := viper.GetString("server.host")
	fmt.Printf("Server is listening at %v\n", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		fmt.Printf("Listen to request fail. Error: %v\b", err)
		return
	}
}
