package main

import (
	"fmt"
	"net/http"
)

func greetUserHandler(W http.ResponseWriter, r *http.Request) {
	fmt.Fprint(W, "Hello, Warrior!")
}

func main() {
	http.HandleFunc("/", greetUserHandler)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
