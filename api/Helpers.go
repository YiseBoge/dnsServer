package api

import (
	"dnsServer/config"
	"dnsServer/models"
	"log"
	"net"
	"net/rpc"
)

func GetClient(node models.ServerNode) *rpc.Client {
	parentAddress := node.Address
	parentPort := node.Port

	client, err := rpc.DialHTTP("tcp", parentAddress+":"+parentPort)
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

	self := models.ServerNode{Address: address, Port: port, Descriptor: descriptor}
	var res bool
	client := ParentClient()
	err := client.Call("API.RegisterChild", self, &res)
	if err != nil {
		log.Fatal("Could not register at the Parent")
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
	return ""
}
