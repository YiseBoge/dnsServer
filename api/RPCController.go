package api

import (
	"dnsServer/config"
	"dnsServer/db"
	"dnsServer/models"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

type API int

var ChildrenNodes []models.ServerNode

func Serve() {
	serverPort := config.LoadConfig().Server.Port

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error Registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":"+serverPort)

	if err != nil {
		log.Fatal("Listener Error", err)
	}
	log.Printf("Serving RPC on port \"%s\"", serverPort)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error Serving: ", err)
	}
}

func (a *API) GetDescriptor(_ string, result *string) error {
	log.Println("Get Descriptor Called______")

	val := config.LoadConfig().Server.Descriptor
	print(val)
	if val == "." {
		*result = ""
	} else {
		client := ParentClient()
		var r string
		err := client.Call("API.GetDescriptor", "", r)
		if err != nil {
			log.Println("Could not get parent descriptor")
			r = ""
		}
		*result = "." + val + r
	}

	log.Println("______Get Descriptor Returning")
	return nil
}

func (a *API) FindAll(_ string, result *[]models.DomainName) error {
	log.Println("Find All Returning______")

	database := db.GetOpenDatabase()
	domains := models.DomainName{}.FindAll(database)
	*result = domains
	defer database.Close()

	log.Println("______Find All Returning")
	return nil
}

func (a *API) Lookup(name string, result *[]models.DomainName) error {
	log.Println("Lookup Called______")

	descriptor := GetMyDescriptor()
	log.Println("Server descriptor:", descriptor)

	cacheDatabase := db.GetOpenCacheDatabase()
	cacheResults := models.DomainName{}.FindByName(cacheDatabase, name)

	if len(cacheResults) > 0 {
		*result = cacheResults
		log.Println("Using Cache...")
		defer cacheDatabase.Close()

	} else {
		if strings.HasSuffix(name, descriptor) {
			foundChild := false
			for _, node := range ChildrenNodes {
				if strings.HasSuffix(name, node.Descriptor) {
					client := GetClient(node)
					var r []models.DomainName
					err := client.Call("API.Lookup", name, &r)
					if err != nil {
						return err
					}
					log.Println("Using Child Server...")
					*result = r
					models.DomainName{}.SaveAll(cacheDatabase, r)
					foundChild = true
					break
				}
			}

			if !foundChild {
				localDatabase := db.GetOpenDatabase()
				localResults := models.DomainName{}.FindByName(localDatabase, name)

				*result = localResults
				log.Println("Using Local Database...")
				defer localDatabase.Close()
			}
		} else {
			client := ParentClient()
			var r []models.DomainName
			err := client.Call("API.Lookup", name, &r)
			if err != nil {
				return err
			}
			log.Println("Using Parent Server...")
			*result = r
			models.DomainName{}.SaveAll(cacheDatabase, r)
		}
	}

	log.Println("______Lookup Returning")
	return nil
}

func (a *API) FindByName(name string, result *[]models.DomainName) error {
	log.Println("Find By Name Called______")

	database := db.GetOpenDatabase()
	domains := models.DomainName{}.FindByName(database, name)
	fmt.Println(domains)
	*result = domains
	defer database.Close()

	log.Println("______Find By Name Returning")
	return nil
}

func (a *API) RegisterChild(child models.ServerNode, result *bool) error {
	log.Println("Register Child Called______")

	ChildrenNodes = append(ChildrenNodes, child)
	*result = true

	log.Println("______Register Child Returning")
	return nil
}
