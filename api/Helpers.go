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
	"time"
)

func GetClient(node models.ServerNode) *rpc.Client {
	address := node.Address
	port := node.Port

	timeout := 3 * time.Second
	_, err := net.DialTimeout("tcp", address+":"+port, timeout)
	if err != nil {
		log.Fatal("Site unreachable, error: ", err)
	}

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
	d := config.LoadConfig()
	val := d.Server.Descriptor
	fullVal := d.Server.FullDescriptor

	if val == "." {
		return ""
	} else if strings.Contains(fullVal, val) {
		return fullVal
	} else {
		client := ParentClient()
		var r string
		err := client.Call("API.GetDescriptor", "", &r)
		if err != nil {
			log.Println("Could not get parent descriptor")
			r = ""
		}
		d.Server.FullDescriptor = "." + val + r
		config.SaveConfig(d)
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
	MoveUnfittingData(client)
}

func InformManager() {
	configuration := config.LoadConfig()
	address := GetMyIP()
	port := configuration.Server.Port
	parentAddress := configuration.Parent.Address
	parentPort := configuration.Parent.Port
	managerAddress := configuration.Manager.Address
	managerPort := configuration.Manager.Port

	managerNode := models.ServerNode{Address: managerAddress, Port: managerPort}
	self := models.ServerModel{Address: address, Port: port, ParentAddress: parentAddress, ParentPort: parentPort}

	var res bool
	client := GetClient(managerNode)
	err := client.Call("API.RegisterServer", self, &res)
	if err != nil {
		log.Fatal("Could not register at the Server Manager")
	}
}

func GetMyIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "localhost"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()

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

func ClearCache(domain models.DomainName) {
	configuration := config.LoadConfig()
	managerAddress := configuration.Manager.Address
	managerPort := configuration.Manager.Port
	managerNode := models.ServerNode{Address: managerAddress, Port: managerPort}
	client := GetClient(managerNode)

	var result bool
	err := client.Call("API.RemoveFromAllCache", domain, &result)
	if err != nil {
		log.Println(err)
	}
}
