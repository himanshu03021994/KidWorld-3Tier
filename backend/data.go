package main

type Alphabet struct {
	Letter string `json:"letter"`
}

var RegionalData = map[string][]Alphabet{
	"Bihar":       {{Letter: "अ"}, {Letter: "आ"}, {Letter: "इ"}, {Letter: "ई"}},
	"Punjab":      {{Letter: "ਅ"}, {Letter: "ਆ"}, {Letter: "ਇ"}, {Letter: "ਈ"}},
	"Bengal":      {{Letter: "অ"}, {Letter: "আ"}, {Letter: "ই"}, {Letter: "ঈ"}},
	"TamilNadu":   {{Letter: "அ"}, {Letter: "ஆ"}, {Letter: "இ"}, {Letter: "ஈ"}},
	"Andhra":      {{Letter: "అ"}, {Letter: "ఆ"}, {Letter: "ఇ"}, {Letter: "ఈ"}},
	"Maharashtra": {{Letter: "अ"}, {Letter: "आ"}, {Letter: "इ"}, {Letter: "ई"}},
	"Karnataka":   {{Letter: "ಅ"}, {Letter: "ಆ"}, {Letter: "ಇ"}, {Letter: "ಈ"}},
}