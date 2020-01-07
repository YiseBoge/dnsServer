package main

import (
	"dnsServer/api"
	"dnsServer/config"
	"dnsServer/models"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"log"
	"regexp"
	"time"
)

func main() {
	//fmt.Println(api.GetMyIP())
	//return
	config.Start()
	fmt.Println("Welcome to the DomaInator server.")
	configuration := config.LoadConfig()

	portRegex, _ := regexp.Compile("^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])(?::([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]))?$")
	domainRegex, _ := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9].[a-zA-Z]{2,}$")
	ipRegex, _ := regexp.Compile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")
	descRegex, _ := regexp.Compile("^[a-zA-Z0-9_.]*$")

	var res1 string
	for true {
		log.Printf("Current port = \"%s\" press 'Enter' to continue or provide new port:", configuration.Server.Port)
		_, _ = fmt.Scanln(&res1)

		if res1 == "" {
			break
		}

		if portRegex.MatchString(res1) {
			configuration.Server.Port = res1
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	var res2 string
	for true {
		log.Printf("Current descriptor = \"%s\" press 'Enter' to continue or provide new descriptor:", configuration.Server.Descriptor)
		_, _ = fmt.Scanln(&res2)

		if res2 == "" {
			break
		}

		if descRegex.MatchString(res2) {
			configuration.Server.Descriptor = res2
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	var res3 string
	for true {
		log.Printf("Parent address = \"%s\" press 'Enter' to continue or provide new address:", configuration.Parent.Address)
		_, _ = fmt.Scanln(&res3)

		if res3 == "" {
			break
		}

		if domainRegex.MatchString(res3) || ipRegex.MatchString(res3) {
			configuration.Parent.Address = res3
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	var res4 string
	for true {
		log.Printf("Parent port = \"%s\" press 'Enter' to continue or provide new port:", configuration.Parent.Port)
		_, _ = fmt.Scanln(&res4)

		if res4 == "" {
			break
		}

		if portRegex.MatchString(res4) {
			configuration.Parent.Port = res4
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

	config.SaveConfig(configuration)
	log.Printf("Parent set to: %s", configuration.Parent)
	api.InformParent()
	gocron.Every(uint64(configuration.Timeout)).Hours().Do(models.ClearTimedOut, configuration.Timeout)
	go api.Serve()

	time.Sleep(1 * time.Second)

	var res string
	for true {
		log.Printf("Type 'exit' or 'stop' to stop serving.")
		_, _ = fmt.Scanln(&res)

		if res == "exit" || res == "stop" {
			break
		}
		log.Printf("**Bad input, Please try again**")
	}

}
