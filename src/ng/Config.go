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
	Domain   string
	IPv4Only bool
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
	flag.StringVar(&(this.Token), "token", "", "API Key from cloudflare.com account settings")
	flag.StringVar(&(this.Email), "email", "", "email from cloudflare.com account settings")
	flag.StringVar(&(this.Domain), "domain", "", "domain you'd like to update. For instance, sub.example.com or example.com")
	flag.BoolVar(&(this.IPv4Only), "ipv4only", false, "set this flag to true if you want to use only IPv4")
	flag.Parse()
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

	if this.Domain == "" {
		message += "Please provide domain\n"
	} else if (len(strings.Split(this.Domain, ".")) < 2) {
		message += "Please provide valid domain\n"
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
