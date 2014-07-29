package main

import (
	"fmt"
	"./ng"
	"strings"
	"time"
)

var api ng.Api

func main() {
	config, err := ng.NewConfig()

	if err != nil {
		return
	}

	ip, ipType := getIP(config.IPv4Only)

	if ip == "" {
		fmt.Println("Unable to detect any type of IP. Probably there is no internet connection.")
		return
	}

	if ipType == "AAAA" {
		fmt.Println("IPv6 detected: ", ip)
	} else {
		fmt.Println("IPv4 detected: ", ip)
	}

	api = ng.NewApi(config)
	updateDomains(config.Domains, ip, ipType)

	if config.Interval > 0 {

		ticker := time.NewTicker(time.Minute * time.Duration(config.Interval)).C

		for {
			select {
			case <-ticker:
				updateDomains(config.Domains, ip, ipType)
			}
		}
	}
}

func updateDomains(domains []string, ip string, ipType string) {
	for _, domain := range domains {
		updateDomain(domain, ip, ipType)
	}
}

func updateDomain(domain string, ip string, ipType string) {
	subDomains := strings.Split(domain, ".")
	var rootDomain string
	if len(subDomains) == 2 {
		rootDomain = domain
	} else {
		rootDomain = strings.Join(subDomains[len(subDomains)-2:], ".")
	}

	result, err := api.RecLoadAll(rootDomain)

	if err != nil {
		fmt.Println("Error while getting "+rootDomain+" info: ", err)
		return
	}

	var obj ng.ObjVO

	for _, v := range result.GetObjs() {
		if v.GetName() == domain && v.GetType() == ipType {
			obj = v
			break;
		}
	}

	if obj == nil {
		fmt.Println("Domain", domain, "with Type", ipType, "doesn't exist. Creating...")
		obj, err = api.RecNew(rootDomain, domain, ip, ipType)
		if err != nil {
			fmt.Println("Error while creating", domain, ":", err)
			return
		}
		fmt.Println("Domain", domain, "was created")
	}

	if obj.GetContent() != ip || obj.GetType() != ipType || obj.GetServiceMode() != 1 {
		fmt.Println("Updating", domain, "IP")
		err = api.RecEdit(rootDomain, obj.GetName(), obj.GetRecID(), ip, ipType)
		if err != nil {
			fmt.Println("Error while updating domain info:", err)
		} else {
			fmt.Println("Domain IP was updated")
		}
	} else {
		fmt.Println("Domain", domain, "is up to date")
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
