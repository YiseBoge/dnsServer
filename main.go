package main

import (
	"dnsServer/api"
	"dnsServer/config"
	"dnsServer/models"
	"fmt"
)

//var database []models.DomainName

//func (a *API) GetDB(empty string, reply *[]models.DomainName) error {
//	database = append(database, models.DomainName{Name: "www.apple.com", Address: "8.8.8.8"})
//	*reply = database
//	return nil
//}

func main() {
	fmt.Println(models.DomainName{}.FindByName("www.google.com"))
	config.Start()
	api.Serve()

	//fmt.Println(models.DomainName{}.FindByName("www.google.com"))
	//a := config.LoadConfig()

	//fmt.Println("Start")

	//var domain = models.DomainNameById(6)
	//domain.Name = "trial Name"
	//domain.UpdateDomainName()
	//fmt.Println(models.DomainNamesByName("www.google.com"))

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

	//var domain2 = models.DomainName{Name: "www.apple.com", Address:"8.8.8.8"}
	//var domain2 = models.DomainName{}.FindById(30)
	//if domain2.ID > 0{
	//	domain2.Address = "1.2.3.4"
	//	domain2.Save()
	//}

	//fmt.Println(models.DomainName{}.FindAll())
	//fmt.Println("End")

}
