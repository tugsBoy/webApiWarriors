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
