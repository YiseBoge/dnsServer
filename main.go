package main

import (
	"dnsServer/config"
	"dnsServer/models"
	"fmt"
)

func main() {
	config.Start()
	//a := config.LoadConfig()

	fmt.Println("Start")

	var domain = models.DomainName{Name: "www.google.com", Address: "8.8.8.10"}
	models.CreateDomainName(domain)
	fmt.Println(models.DomainNamesByName("www.google.com"))

	//var domain1 = models.DomainNameById(0)
	//fmt.Println(domain1)
	//var domain3 = models.DomainNameByName("www.google.com")
	//fmt.Println(domain3)

	//var domains = models.AllDomainNames()
	//fmt.Println(domains)
	//models.DeleteDomainName(domain1)
	//fmt.Println(domain1)

	//var domain1 = models.DomainNameById(1)
	//fmt.Println(domain1)
	//models.DeleteDomainName(domain1)
	//fmt.Println(models.AllDomainNames())

	//var domain2 = models.DomainName{Name: "www.apple.com", Address:"1.2.3.4"}
	//models.CreateDomainName(domain2)
	//fmt.Println(domain2)

	fmt.Println("End")

}
