package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// LetterData represents a single alphabet entry
type LetterData struct {
	Letter string `json:"letter"`
	Word   string `json:"word"`
	Emoji  string `json:"emoji"`
}

// RegionalData stores the alphabets for each supported language
var RegionalData = map[string][]LetterData{
	"Hindi": {
		{"अ", "Anaar (Pomegranate)", "🍎"},
		{"आ", "Aam (Mango)", "🥭"},
		{"इ", "Imli (Tamarind)", "🫘"},
		{"ई", "Eekh (Sugarcane)", "🎋"},
		{"उ", "Ullu (Owl)", "🦉"},
		{"ऊ", "Oon (Wool)", "🧶"},
		{"ए", "Ek (One)", "1️⃣"},
		{"ऐ", "Ainak (Spectacles)", "👓"},
		{"ओ", "Okhli (Mortar)", "🥣"},
		{"औ", "Aurat (Woman)", "👩"},
		{"क", "Kabootar (Pigeon)", "🕊️"},
		{"ख", "Khargosh (Rabbit)", "🐇"},
		{"ग", "Gamla (Flower Pot)", "🪴"},
		{"घ", "Ghar (House)", "🏠"},
	},
	"Punjabi": {
		{"ੳ", "Ooth (Camel)", "🐪"},
		{"ਅ", "Anaar (Pomegranate)", "🍎"},
		{"ੲ", "Itt (Brick)", "🧱"},
		{"ਸ", "Saeb (Apple)", "🍏"},
		{"ਹ", "Hathi (Elephant)", "🐘"},
		{"ਕ", "Kabootar (Pigeon)", "🕊️"},
		{"ਖ", "Khamb (Feather)", "🪶"},
		{"ਗ", "Gamla (Flower Pot)", "🪴"},
		{"ਘ", "Ghar (House)", "🏠"},
	},
	"Marathi": {
		{"अ", "Ananas (Pineapple)", "🍍"},
		{"आ", "Aai (Mother)", "👩‍👦"},
		{"इ", "Imarat (Building)", "🏢"},
		{"ई", "Idlimbu (Lemon)", "🍋"},
		{"उ", "Undir (Mouse)", "🐁"},
		{"ऊ", "Oos (Sugarcane)", "🎋"},
		{"ए", "Edka (Ram)", "🐏"},
		{"ऐ", "Airan (Anvil)", "🗜️"},
		{"ओ", "Oze (Load)", "🎒"},
		{"औ", "Aushadh (Medicine)", "💊"},
	},
	"Gujarati": {
		{"અ", "Ananas (Pineapple)", "🍍"},
		{"આ", "Aag (Fire)", "🔥"},
		{"ઇ", "Iyal (Caterpillar)", "🐛"},
		{"ઈ", "Eesh (God)", "✨"},
		{"ઉ", "Undar (Mouse)", "🐁"},
		{"ઊ", "Oon (Wool)", "🧶"},
		{"એ", "Ek (One)", "1️⃣"},
		{"ઐ", "Airavat (Elephant)", "🐘"},
		{"ઓ", "Oshiku (Pillow)", "🛌"},
		{"ઔ", "Aushadh (Medicine)", "💊"},
	},
}

// enableCORS is a helper function to allow frontend to talk to backend
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	// Route 1: Setup User Profile & Language
	http.HandleFunc("/setup", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		var profile map[string]interface{}
		json.NewDecoder(r.Body).Decode(&profile)

		name := "Kid"
		state := "Bihar" // Defaults
		if n, ok := profile["name"].(string); ok { name = n }
		if s, ok := profile["state"].(string); ok { state = s }

		language := "Hindi"
		module := "Varnamala"

		if state == "Punjab" {
			language = "Punjabi"
			module = "Gurmukhi"
		} else if state == "Maharashtra" {
			language = "Marathi"
			module = "Mulakhshare"
		} else if state == "Gujarat" {
			language = "Gujarati"
			module = "Kakko"
		}

		response := map[string]string{
			"message":  "Welcome " + name + "!",
			"language": language,
			"module":   module,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Route 2: Fetch Alphabet Data for the Game
	http.HandleFunc("/get-alphabet", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		state := r.URL.Query().Get("state")
		language := "Hindi" 
		
		if state == "Punjab" { language = "Punjabi" } else if state == "Maharashtra" { language = "Marathi" } else if state == "Gujarat" { language = "Gujarati" }

		data := RegionalData[language]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}