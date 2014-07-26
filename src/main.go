package main

import (
	"flag"
	"fmt"
	"./ng"
	"strings"
)

var api ng.Api

func main() {
	config := ng.Config{}

	flag.StringVar(&(config.Token), "token", "", "API Key from cloudflare.com account settings")
	flag.StringVar(&(config.Email), "email", "", "email from cloudflare.com account settings")
	flag.StringVar(&(config.Domain), "domain", "", "domain you'd like to update. For instance, sub.example.com or example.com")
	var ipv4Only bool
	flag.BoolVar(&ipv4Only, "ipv4only", false, "set this flag to true if you want to use only IPv4")
	flag.Parse()

	if !config.Validate() {
		return
	}

	ip, ipType := getIP(ipv4Only)

	if ip == "" {
		fmt.Println("Unable to detect any type of IP. Probably there is no internet connection.")
		return
	}

	if ipType == "AAAA" {
		fmt.Println("IPv6 detected: ", ip)
	} else {
		fmt.Println("IPv4 detected: ", ip)
	}

	subDomains := strings.Split(config.Domain, ".")
	var rootDomain string
	if len(subDomains) == 2 {
		rootDomain = config.Domain
	} else {
		rootDomain = strings.Join(subDomains[len(subDomains)-2:], ".")
	}

	api = ng.NewApi(config)
	result, err := api.RecLoadAll(rootDomain)

	if err != nil {
		fmt.Println("Error while getting "+rootDomain+" info: ", err)
		return
	}

	var obj ng.ObjVO

	for _, v := range result.GetObjs() {
		if v.GetName() == config.Domain && v.GetType() == ipType {
			obj = v
			break;
		}
	}

	if obj == nil {
		fmt.Println("Domain", config.Domain, "with Type", ipType, "doesn't exist. Creating...")
		obj, err = api.RecNew(rootDomain, config.Domain, ip, ipType)
		if err != nil {
			fmt.Println("Error while creating", config.Domain, ":", err)
			return
		}
		fmt.Println("Domain", config.Domain, "was created")
	}

	if obj.GetContent() != ip || obj.GetType() != ipType || obj.GetServiceMode() != 1 {
		fmt.Println("Updating", config.Domain, "IP")
		err = api.RecEdit(rootDomain, obj.GetName(), obj.GetRecID(), ip, ipType)
		if err != nil {
			fmt.Println("Error while updating domain info:", err)
		} else {
			fmt.Println("Domain IP was updated")
		}
	} else {
		fmt.Println("Domain is up to date")
	}
}

func getIP(ipv4Only bool) (string, string) {

	ip := ""
	var ipType string

	if !ipv4Only {
		ipType = "AAAA"
		ip = ng.GetIpv6()
	}

	if ip == "" {
		ipType = "A"
		ip = ng.GetIpv4()
	}

	return ip, ipType
}
