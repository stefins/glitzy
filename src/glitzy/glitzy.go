package glitzy

import (
	"fmt"

	"github.com/iamstefin/glitzy/src/utils"
)

// Add will add a new password to the database
func Add() (err error) {
	fmt.Println(utils.GetInfo())
	return
}

// Wipe will remove all the passwords
func Wipe() (err error) {
	fmt.Println("Clean the entire passwords")
	return
}
