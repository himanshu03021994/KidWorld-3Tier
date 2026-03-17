package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type KidProfile struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	MobileNo string `json:"mobile"`
	Town     string `json:"town"`
	State    string `json:"state"`
}

type AppResponse struct {
	Message  string `json:"message"`
	Language string `json:"language"`
	Module   string `json:"module"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Please send a POST request", http.StatusMethodNotAllowed)
		return
	}

	var profile KidProfile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Invalid data format", http.StatusBadRequest)
		return
	}

	state := strings.ToLower(profile.State)
	language := "English"
	module := "Standard ABCs"

	switch state {
	case "bihar", "uttar pradesh", "madhya pradesh":
		language = "Hindi"
		module = "Varnamala (क से कबूतर)"
	case "maharashtra":
		language = "Marathi"
		module = "Marathi Mulakshare"
	case "gujarat":
		language = "Gujarati"
		module = "Gujarati Kakko"
	}

	response := AppResponse{
		Message:  fmt.Sprintf("Welcome to Kid World, %s!", profile.Name),
		Language: language,
		Module:   module,
	}

	json.NewEncoder(w).Encode(response)
}

	func main() {
	// 1. New Regional Alphabet Route
	http.HandleFunc("/get-alphabet", func(w http.ResponseWriter, r *http.Request) {
		// Allow any website to read this data (CORS bypass)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// Get the state from the URL, default to Bihar if none is provided
		state := r.URL.Query().Get("state")
		data, exists := RegionalData[state]
		if !exists {
			data = RegionalData["Bihar"] 
		}
		
		json.NewEncoder(w).Encode(data)
	})

	// 2. Cloud-Ready Port Setup (Crucial for Render)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback for your local laptop
	}

	fmt.Println("Kid World Server starting on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}