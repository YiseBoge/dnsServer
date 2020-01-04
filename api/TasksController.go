package api

import (
	"dnsServer/config"
	"dnsServer/models"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

func Serve() {
	//serverPort := config.LoadConfig().Server.Port
	serverPort := "4040"
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
	log.Printf("serving rpc on port %d", 4040)
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

func (a *API) FindAll(empty string, result *[]models.DomainName) error {
	fmt.Println("called find all")
	domains := models.DomainName{}.FindAll()
	*result = domains
	fmt.Println("returning find all")
	return nil
}

//func (api *API) Lookup(name string, address *string) error {
//	print("abebe")
//	*address = name
//
//	client := ParentClient()
//	var result string
//	err := client.Call("API.Lookup", "abebe", &result)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (a *API) FindByName(name string, result *[]models.DomainName) error {
	fmt.Println("called find by name")
	fmt.Println(name)
	domains := models.DomainName{}.FindByName(name)
	fmt.Println(domains)
	*result = domains
	fmt.Println("returning find by name")
	return nil
}
