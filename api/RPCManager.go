package api

import (
	"dnsServer/config"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func Setup() {
	serverPort := config.LoadConfig().Server.Port
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":"+serverPort)

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %s", serverPort)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}

func ParentClient() *rpc.Client {
	parentAddress := config.LoadConfig().Parent.Address
	parentPort := config.LoadConfig().Parent.Port

	client, _ := rpc.DialHTTP("tcp", parentAddress+":"+parentPort)
	return client
}
