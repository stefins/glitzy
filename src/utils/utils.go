package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/iamstefin/glitzy/src/models"
	"github.com/manifoldco/promptui"
	"gorm.io/gorm"
)

// GetInfo input the user info from terminal
func GetInfo() *models.User {
	ServiceName := GetNormalString("Service Name ")
	Username := GetNormalString("Username ")
	Password := getProtectedString("Password")
	return &models.User{Name: ServiceName, Username: Username, Password: Password}
}

// GetNormalString input normal text from terminal
func GetNormalString(promptText string) string {
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

// AddOrCheckMainPassword add a main password to the application
func AddOrCheckMainPassword(db *gorm.DB) {
	if (db.Find(&models.Main{}).RowsAffected == 0) {
		fmt.Println("Main Password not configured!")
		mainPass := getProtectedString("Enter the Main Password")
		mainPassRe := getProtectedString("Enter the Main Password again")
		if mainPass != mainPassRe {
			log.Fatalf("Password Mismatch!")
		}
		hash := sha256.New()
		hash.Write([]byte(mainPass))
		err := db.Create(&models.Main{PasswordHash: hex.EncodeToString(hash.Sum(nil))}).Error
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		var dbInfo models.Main
		db.First(&dbInfo)
		userInputMainPass := getProtectedString("Enter the Main Password")
		hash := sha256.New()
		hash.Write([]byte(userInputMainPass))
		if dbInfo.PasswordHash == hex.EncodeToString(hash.Sum(nil)) {
			fmt.Println("Authenticated with Main Password")
		} else {
			fmt.Println("Invalid Password!")
			os.Exit(0)
		}
	}
}
