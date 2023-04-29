package main

import (
	GrpcUserService "WebServer/clients/UserService"
	handler "WebServer/handler"
	"WebServer/utils"
	"fmt"
	"net/http"

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
	http.HandleFunc("/api/RegisterUser", handler.RegisterUser)
	host := viper.GetString("server.host")
	fmt.Printf("Server is listening at %v\n", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		fmt.Printf("Listen to request fail. Error: %v\b", err)
		return
	}

}
