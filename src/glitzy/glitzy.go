package glitzy

import (
	"fmt"
	"time"
)

// Add will add a new password to the database
func Add() {
	fmt.Println("Add a new password")
	time.Sleep(1 * time.Second)
}

// Wipe will remove all the passwords
func Wipe() {
	fmt.Println("Clean the entire passwords")
}
