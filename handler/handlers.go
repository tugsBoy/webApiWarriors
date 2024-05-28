package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tugsBoy/webApiWarriors/types"
)

func GreetUser(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(W).Encode(types.ImutableGreeting.GetGreeting()); err != nil {
		http.Error(W, err.Error(), http.StatusInternalServerError)
	}
}

func NewUser(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Contetn-Type", "application/json")
	if err := json.NewEncoder(W).Encode(types.ImutableNewUserPrompt.PromptUserName); err != nil {
		http.Error(W, err.Error(), http.StatusInternalServerError)
	}
}
