package main

import (
	"fmt"
	"net/http"

	"github.com/tugsBoy/webApiWarriors/api"
)

func main() {
	api.RegisterRoutes()
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
