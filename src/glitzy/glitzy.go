package glitzy

import (
	"fmt"
	"log"
	"os"

	"github.com/iamstefin/glitzy/src/models"
	"github.com/iamstefin/glitzy/src/utils"
	"github.com/manifoldco/promptui"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Add will add a new password to the database
func Add() (err error) {
	db, err := initDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	utils.AddOrCheckMainPassword(db)
	err = db.Create(utils.GetInfo()).Error
	if err != nil {
		log.Fatalf("Failed to add password!")
	}
	fmt.Println("Data added!")
	return
}

// Search will search for the passwords
func Search() (err error) {
	db, err := initDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	searchTerm := utils.GetNormalString("Enter the search term")
	var rtrmdl []models.User
	if db.Where("name LIKE ?", "%"+searchTerm+"%").Find(&rtrmdl).RowsAffected == 0 {
		fmt.Println("Password not found!")
		os.Exit(0)
	}
	templates := &promptui.SelectTemplates{
		Active:   " \U00002705 {{ .Name | cyan }} ({{ .Username }})",
		Inactive: " {{ .Name | cyan }} ({{ .Username }})",
		Selected: " \U00002705 {{ .Name }} ({{ .Username }})",
	}
	prompt := promptui.Select{
		Label:     "Select",
		Items:     rtrmdl,
		HideHelp:  true,
		Templates: templates,
	}
	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("The Password for %v Is %v\n", rtrmdl[i].Username, rtrmdl[i].Password)
	return nil
}

// Wipe will remove all the passwords
func Wipe() (err error) {
	db, err := initDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	utils.AddOrCheckMainPassword(db)
	fmt.Printf("%v Records Deleted \n", db.Exec("DELETE FROM users").RowsAffected)
	os.Remove("test.db")
	return
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Main{})
	return db, nil
}
