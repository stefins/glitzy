package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// User hold the info of a credential
type User struct {
	serviceName string
	username    string
	password    string
}

// GetInfo input the user info from terminal
func GetInfo() User {
	var service string
	var username string
	var password string
	for {
		service = getNormalString("Enter the Service Name: ")
		if service != "" {
			break
		}
	}
	for {
		username = getNormalString("Enter the username: ")
		if username != "" {
			break
		}
	}
	for {
		password = getProtectedString("Enter the password: ")
		if password != "" {
			break
		}
	}

	return User{service, username, password}
}

// getNormalString input normal text from terminal
func getNormalString(promptText string) string {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%v", err)
	}
	value = strings.ReplaceAll(value, "\n", "")
	return value
}

// getProtectedString input protected text from terminal
func getProtectedString(promptText string) string {
	fmt.Print(promptText)
	val, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatalf("%v", err)
	}
	value := strings.ReplaceAll(string(val), "\n", "")
	return value
}
