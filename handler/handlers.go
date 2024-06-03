package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/tugsBoy/webApiWarriors/services"
	"github.com/tugsBoy/webApiWarriors/types"
)

func GreetUser(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(W).Encode(types.ImutableGreeting.GetGreeting()); err != nil {
		http.Error(W, err.Error(), http.StatusInternalServerError)
	}
}

// type LoginHandler struct {
// 	Login services.ProcessLoginer
// }

func HandleLogin(W http.ResponseWriter, r *http.Request) {
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(W, "Invalid JSON input.", http.StatusBadRequest)
		return
	}
	playerSave, err := services.ProcessLogin(user)
	if err != nil {
		if errors.Is(err, types.ErrUserNotFound) || errors.Is(err, types.ErrBlankUserID) {
			http.Error(W, err.Error(), http.StatusNotFound)
		} else {
			http.Error(W, "Internal Server Error", http.StatusInternalServerError)
		}
		log.Printf("Error fetching user save data: %V", err)
		return
	}
	W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(W).Encode(playerSave)
}

// func NewUser(W http.ResponseWriter, r *http.Request) {
// 	W.Header().Set("Contetn-Type", "application/json")
// 	if err := json.NewEncoder(W).Encode(types.ImutableNewUserPrompt.GetNewUserPrompt()); err != nil {
// 		http.Error(W, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func ExistingUser(W http.ResponseWriter, r *http.Request) {
// 	W.Header().Set("Contetn-Type", "application/json")
// 	userSave, err := types.PlayerSave.GetPlayerSave("bob", db )
// }
