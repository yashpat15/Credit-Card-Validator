package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateCardHandler(t *testing.T) {
	// Create a request with a valid credit card number.
	req, err := http.NewRequest("POST", "/validate", bytes.NewBufferString(`{"number":"4532015112830366"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Record the response using httptest.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(validateCardHandler)

	// Serve HTTP request to the handler.
	handler.ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the JSON response and evaluate the results.
	var result map[string]bool
	if err := json.Unmarshal(rr.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if isValid, exists := result["valid"]; !exists || !isValid {
		t.Errorf("handler returned unexpected body: got %v want %v", result, map[string]bool{"valid": true})
	}
}
