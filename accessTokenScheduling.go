package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	tokenMutex  sync.Mutex
	accessToken string
)

func main() {
	// Initial token generation
	refreshToken()

	// Simulate making API calls using the access token
	makeAPICalls()

	// Schedule token refresh every 5 minutes in the background
	go scheduleTokenRefresh()

	// Keep the main application running
	select {}
}

func getToken() string {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	return accessToken
}

func setToken(newToken string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	accessToken = newToken
}

func refreshToken() {
	// Obtain a new token from your authentication endpoint (implement your logic here)
	newToken := fetchNewToken()

	// Update the access token
	setToken(newToken)

	fmt.Println("Token refreshed:", newToken)
}

func fetchNewToken() string {
	// Simulate fetching a new token from an authentication endpoint (replace with actual logic)
	return "new_access_token"
}

func makeAPICalls() {
	// Use the current access token for making API calls (read-only operation)
	currentToken := getToken()

	// Simulate making API calls using the current token
	fmt.Println("Making API calls with token:", currentToken)
}

func scheduleTokenRefresh() {
	for {
		// Sleep for 5 minutes
		time.Sleep(5 * time.Minute)

		// Refresh the token in the background
		go refreshToken()
	}
}
