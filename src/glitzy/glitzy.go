package glitzy

import (
	"fmt"
	"log"

	"github.com/iamstefin/glitzy/src/models"
	"github.com/iamstefin/glitzy/src/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Add will add a new password to the database
func Add() (err error) {
	db, err := initDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = db.Create(utils.GetInfo()).Error
	if err != nil {
		log.Fatalf("Failed to add password!")
	}
	fmt.Println("Data added!")
	return
}

// Wipe will remove all the passwords
func Wipe() (err error) {
	db, err := initDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%v Records Deleted \n", db.Exec("DELETE FROM users").RowsAffected)
	return
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}
