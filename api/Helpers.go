package api

import (
	"dnsServer/config"
	"dnsServer/db"
	"dnsServer/models"
	"github.com/jinzhu/gorm"
	"log"
	"net"
	"net/rpc"
	"strings"
)

func GetClient(node models.ServerNode) *rpc.Client {
	address := node.Address
	port := node.Port

	client, err := rpc.DialHTTP("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ParentClient() *rpc.Client {
	parentAddress := config.LoadConfig().Parent.Address
	parentPort := config.LoadConfig().Parent.Port
	client := GetClient(models.ServerNode{Address: parentAddress, Port: parentPort})
	return client
}

func GetMyDescriptor() string {
	val := config.LoadConfig().Server.Descriptor

	if val == "." {
		return ""
	} else {
		client := ParentClient()
		var r string
		err := client.Call("API.GetDescriptor", "", &r)
		if err != nil {
			log.Println("Could not get parent descriptor")
			r = ""
		}
		return "." + val + r
	}
}

func InformParent() {
	address := GetMyIP()
	if address == "" {
		log.Fatal("IP returned Nothing")
	}
	port := config.LoadConfig().Server.Port
	descriptor := GetMyDescriptor()

	if descriptor == "" {
		return
	}

	self := models.ServerNode{Address: address, Port: port, Descriptor: descriptor}
	var res bool
	client := ParentClient()
	err := client.Call("API.RegisterChild", self, &res)
	if err != nil {
		log.Fatal("Could not register at the Parent")
	}
}

func InformManager() {
	configuration := config.LoadConfig()
	address := GetMyIP()
	port := configuration.Server.Port
	parentAddress := configuration.Parent.Address
	parentPort := configuration.Parent.Port

	managerNode := models.ServerNode{Address: configuration.Manager.Address, Port: configuration.Manager.Port}
	self := models.ServerModel{Address: address, Port: port, ParentAddress: parentAddress, ParentPort: parentPort}

	var res bool
	client := GetClient(managerNode)
	err := client.Call("API.RegisterServer", self, &res)
	if err != nil {
		log.Fatal("Could not register at the Server Manager")
	}
}

func GetMyIP() string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("Could not find my IP")
	}
	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "localhost"
}

func MoveUnfittingData(client *rpc.Client) {
	descriptor := GetMyDescriptor()
	database := db.GetOpenDatabase()
	domains := models.DomainName{}.FindAll(database)
	var result bool

	for _, domain := range domains {
		if !strings.HasSuffix(domain.Name, descriptor) {
			go __MoveData(client, domain, result, database)
		}
	}
}

func __MoveData(client *rpc.Client, domain models.DomainName, result bool, database *gorm.DB) {
	err := client.Call("API.Register", domain, &result)
	if err != nil {
		log.Println(err)
	}
	domain.Delete(database)
}
