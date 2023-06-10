package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"miniproject3/intenal/service"
	"miniproject3/intenal/usecase"
)

func main() {
	apiURL := "https://reqres.in/api/users?page=2" // Update the URL here
	userService := service.NewUserService(apiURL)
	userUseCase := usecase.NewUserUseCase(userService)

	userID := 1
	user, err := userUseCase.GetUserByID(userID)
	if err != nil {
		log.Fatalf("Failed to get user with ID %d: %v", userID, err)
	}

	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
}

func serveHTTP() {
	// Handle HTTP requests here
}

func setupRoutes() {
	// Define your HTTP routes here
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	setupRoutes()

	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
