package utils

import (
	"errors"
	"log"

	"github.com/iamstefin/glitzy/src/models"
	"github.com/manifoldco/promptui"
)

// GetInfo input the user info from terminal
func GetInfo() (user models.User) {
	user.ServiceName = getNormalString("Service Name ")
	user.Username = getNormalString("Username ")
	user.Password = getProtectedString("Password")
	return user
}

// getNormalString input normal text from terminal
func getNormalString(promptText string) string {
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("Must be more than 4 characters")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    promptText,
		Validate: validate,
	}
	result, err := prompt.Run()

	if err != nil {
		log.Fatalf("%v", err)
	}
	return result
}

// getProtectedString input protected text from terminal
func getProtectedString(promptText string) string {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("Password must have more than 6 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    promptText,
		Validate: validate,
		Mask:     '.',
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("%v", err)
	}

	return result
}
