package types

// Greeting For all Players
type greeting struct {
	Greet      string `json:"greet"`
	AskHistory string `json:"askHistory"`
}

var ImutableGreeting = greeting{
	Greet:      "Hello Warrior!",
	AskHistory: "Do you have a Login?(Yes/No)",
}

func (g greeting) GetGreeting() greeting {
	return ImutableGreeting
}

// New Users
type newUserPrompt struct {
	PromptUserName string `json:"promptUserName"`
}

var ImutableNewUserPrompt = newUserPrompt{
	PromptUserName: "Hello! Please Provide a Username.",
}

func (nup newUserPrompt) GetNewUserPrompt() newUserPrompt {
	return ImutableNewUserPrompt
}

// Warriors structs/methods
type Warrior struct {
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Age          string `json:"age"`
	Strength     int32  `json:"strength"`
	Dex          string `json:"dex"`
	Intelligence int32  `json:"intelligence"`
	Faith        int32  `json:"faith"`
	Arcane       int32  `json:"arcane"`
	Luck         int32  `json:"luck"`
}

// Player Save File Structs
type PlayerSave struct {
	Warriors []Warrior `json:"warriors"`
}

// Struct for mockDB
type MockDB struct {
	PlayerSaves []PlayerSave
}
