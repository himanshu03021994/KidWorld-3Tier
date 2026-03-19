package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type LanguageProfile struct {
	Script    string   `json:"script"`
	LangCode  string   `json:"langCode"`
	Direction string   `json:"direction"`
	Alphabets []string `json:"alphabets"`
}

var RegionalRegistry = map[string]LanguageProfile{
	"Bihar":       {"Devanagari", "hi-IN", "ltr", []string{"अ", "आ", "इ", "ई", "उ", "ऊ", "ए", "ऐ", "ओ", "औ"}},
	"Maharashtra": {"Devanagari", "mr-IN", "ltr", []string{"अ", "आ", "इ", "ई", "उ", "ऊ", "ए", "ऐ", "ओ", "औ"}},
	"West Bengal": {"Bengali", "bn-IN", "ltr", []string{"অ", "আ", "ই", "ঈ", "উ", "ঊ", "এ", "ঐ", "ও", "ঔ"}},
	"Punjab":      {"Gurmukhi", "pa-IN", "ltr", []string{"ੳ", "ਅ", "ੲ", "ਸ", "ਹ", "ਕ", "ਖ", "ਗ", "ਘ", "ਙ"}},
	"Gujarat":     {"Gujarati", "gu-IN", "ltr", []string{"અ", "આ", "ઇ", "ઈ", "ઉ", "ઊ", "એ", "ઐ", "ઓ", "ઔ"}},
	"Tamil Nadu":  {"Tamil", "ta-IN", "ltr", []string{"அ", "ஆ", "இ", "ஈ", "உ", "ஊ", "எ", "ஏ", "ஐ", "ஒ"}},
	"Telangana":   {"Telugu", "te-IN", "ltr", []string{"అ", "ఆ", "ఇ", "ఈ", "ఉ", "ఊ", "ఎ", "ఏ", "ఐ", "ఒ"}},
	"Karnataka":   {"Kannada", "kn-IN", "ltr", []string{"అ", "ఆ", "ఇ", "ఈ", "ఉ", "ఊ", "ఎ", "ఏ", "ఐ", "ఒ"}},
	"Kerala":      {"Malayalam", "ml-IN", "ltr", []string{"അ", "ആ", "ഇ", "ഈ", "ഉ", "ഊ", "എ", "ഏ", "ഐ", "ഒ"}},
	"Odisha":      {"Odia", "or-IN", "ltr", []string{"ଅ", "ଆ", "ଇ", "ଈ", "ଉ", "ଊ", "ଏ", "ଐ", "ଓ", "ঔ"}},
	"Jharkhand":   {"Ol Chiki", "sat-IN", "ltr", []string{"ᱚ", "ᱛ", "ᱜ", "ᱝ", "ᱞ", "ᱟ", "ᱠ", "ᱡ", "ᱢ", "ᱣ"}},
}

func getAlphabet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	state := r.URL.Query().Get("state")
	profile, flux := RegionalRegistry[state]
	if !flux {
		profile = RegionalRegistry["Bihar"] // Default fallback
	}
	json.NewEncoder(w).Encode(profile)
}

func main() {
	http.HandleFunc("/get-alphabet", getAlphabet)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
