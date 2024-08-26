// handler.go
package main

import (
	"encoding/json"
	"net/http"
)

// CardInfo holds the credit card number.
type CardInfo struct {
	Number string `json:"number"`
}

// handles the POST requests for card validation.
func validateCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var card CardInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&card); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid := ValidateCreditCard(card.Number)
	response := map[string]bool{"valid": isValid}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
