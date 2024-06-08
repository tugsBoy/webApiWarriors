package types

import "errors"

//import "errors"

// Greeting For all Players
type greeting struct {
	Greet        string `json:"greet"`
	PromptUserID string `json:"promptUserID"`
}

var ImutableGreeting = greeting{
	Greet:        "Hello Warrior!",
	PromptUserID: "Please Provide your user ID",
}

func (g greeting) GetGreeting() greeting {
	return ImutableGreeting
}

type User struct {
	LoginID string `json:"loginID"`
}

type UserCreateSuccess struct {
	Success string `json:"success"`
}

type UserCreateFail struct {
	Fail string `json:"fail"`
}

// Warriors structs/methods
type Warrior struct {
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Age          int32  `json:"age"`
	Strength     int32  `json:"strength"`
	Dex          string `json:"dex"`
	Intelligence int32  `json:"intelligence"`
	Faith        int32  `json:"faith"`
	Arcane       int32  `json:"arcane"`
	Luck         int32  `json:"luck"`
}

// Player Save File Structs and methods
type PlayerSave struct {
	LoginId  string    `json:"loginID"`
	Warriors []Warrior `json:"warriors"`
}

var ErrUserNotFound = errors.New("user not found")
var ErrBlankUserID = errors.New("blank user ID")
