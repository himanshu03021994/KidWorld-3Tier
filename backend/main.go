package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	// 1. Host the Frontend Website
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	// 2. Host the Backend API
	http.HandleFunc("/api/register", registerHandler)

	fmt.Println("🚀 Kid World is LIVE!")
	fmt.Println("👉 Open Chrome and go to: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}