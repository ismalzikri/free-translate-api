package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateHandler(t *testing.T) {
	// Create a test request payload
	requestPayload := TranslateRequest{
		Text: "Hello",
		To:   "es",
	}
	payloadBytes, _ := json.Marshal(requestPayload)

	// Create a test request
	req, err := http.NewRequest("POST", "/translate", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the TranslateHandler function directly
	handler := http.HandlerFunc(TranslateHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, rr.Code)
	}

	// Parse the response body
	var response TranslateResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Check the response fields
	if response.TranslatedText == "" {
		t.Error("TranslatedText field is empty")
	}
	if response.Status != true {
		t.Error("Status field is not true")
	}
	if response.Message != "" {
		t.Error("Message field is not empty")
	}
}
