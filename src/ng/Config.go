package ng

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Token    string
	Email    string
	Domains  []string
	IPv4Only bool
	Interval int
}

func NewConfig() (Config, error) {
	result := Config{}

	if len(os.Args[1:]) == 0 || indexOfString(os.Args[1:], "-config") != -1 {
		result.readConfigFile()
	} else {
		result.readFlags()
	}

	if !result.validate() {
		return result, errors.New("")
	} else {
		return result, nil
	}
}

func (this *Config) readFlags() {
	var domains string
	flag.StringVar(&(this.Token), "token", "", "API Key from cloudflare.com account settings")
	flag.StringVar(&(this.Email), "email", "", "email from cloudflare.com account settings")
	flag.StringVar(&domains, "domains", "", "domain you'd like to update. For instance, sub.example.com or example.com")
	flag.BoolVar(&(this.IPv4Only), "ipv4only", false, "set this flag to true if you want to use only IPv4")
	flag.IntVar(&(this.Interval), "interval", 0, "interval in minutes between updates. If 0 update processed once.")
	flag.Parse()

	domains = strings.TrimSpace(domains)
	this.Domains = strings.Split(domains, ",")
}

func (this *Config) readConfigFile() {
	var filePath string
	flag.StringVar(&filePath, "config", "config.json", "Path to config file")
	flag.Parse()

	configContent, err := ioutil.ReadFile(string(filePath))

	if err != nil {
		fmt.Println("Error while reading config file", filePath, err)
		return
	}

	err = json.Unmarshal(configContent, this)

	if err != nil {
		fmt.Println("Error while parsing config file", filePath, err)
	}
}

func (this *Config) validate() bool {

	message := ""

	if this.Token == "" {
		message += "Please provide token\n"
	}

	if (this.Email == "") {
		message += "Please provide email\n"
	}

	if len(this.Domains) == 0 {
		message += "Please provide at least one domain name\n"
	}

	for index, domain := range this.Domains {
		if domain == "" {
			message += "One of the domain name is empty\n"
			this.Domains = append(this.Domains[:index], this.Domains[index+1:]...)
		} else if (len(strings.Split(domain, ".")) < 2) {
			message += "Domain " + domain + " is invalid\n"
			this.Domains = append(this.Domains[:index], this.Domains[index+1:]...)
		}
	}

	if message == "" {
		return true;
	} else {
		fmt.Print(message)
		return false
	}
}

func indexOfString(slice []string, value string) int {
	for index, v := range slice {
		if strings.Contains(v, value) {
			return index
		}
	}
	return -1
}
