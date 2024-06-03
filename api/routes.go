package api

import (
	"net/http"

	"github.com/tugsBoy/webApiWarriors/handler"
)

func RegisterRoutes() {
	http.HandleFunc("/", handler.GreetUser)
	http.HandleFunc("/login", handler.HandleLogin)
}
