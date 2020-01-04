package main

import (
	"dnsServer/api"
	"dnsServer/config"
	"fmt"
	"log"
	"regexp"
	"time"
)

func main() {
	config.Start()
	fmt.Println("Welcome to the DomaInator server.")
	configuration := config.LoadConfig()

	portRegex, _ := regexp.Compile("^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])(?::([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]))?$")
	domainRegex, _ := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9].[a-zA-Z]{2,}$")
	ipRegex, _ := regexp.Compile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")

	var res string

	for true {
		log.Printf("Current port = \"%s\" press 'Enter' to continue or provide new port:", configuration.Server.Port)
		_, _ = fmt.Scanln(&res)

		if res == "" {
			break
		}

		if portRegex.MatchString(res) {
			configuration.Server.Port = res
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	for true {
		log.Printf("Parent address = \"%s\" press 'Enter' to continue or provide new address:", configuration.Parent.Address)
		_, _ = fmt.Scanln(&res)

		if res == "" {
			break
		}

		if domainRegex.MatchString(res) || ipRegex.MatchString(res) {
			configuration.Parent.Address = res
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	for true {
		log.Printf("Parent port = \"%s\" press 'Enter' to continue or provide new port:", configuration.Parent.Port)
		_, _ = fmt.Scanln(&res)

		if res == "" {
			break
		}

		if portRegex.MatchString(res) {
			configuration.Parent.Port = res
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	config.SaveConfig(configuration)
	log.Printf("Parent set to: %s", configuration.Parent)
	go api.Serve()

	time.Sleep(1 * time.Second)
	for true {
		log.Printf("Type 'exit' or 'stop' to stop serving.")
		_, _ = fmt.Scanln(&res)

		if res == "exit" || res == "stop" {
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	//fmt.Println(models.DomainName{}.FindByName("www.google.com"))
	//configuration := config.LoadConfig()

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
