package ng

import (
	"fmt"
	"strings"
)

type Config struct {
	Token   string
	Email   string
	Domain  string
}

func (this *Config) Validate() bool {

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
