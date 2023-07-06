package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bregydoc/gtranslate"
	"github.com/rs/cors"
)

type TranslateRequest struct {
	Text string `json:"text"`
	To   string `json:"to"`
}

type TranslateResponse struct {
	TranslatedText string `json:"translatedText,omitempty"`
	Status         bool   `json:"status"`
	Message        string `json:"message"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request TranslateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	translated, err := gtranslate.TranslateWithParams(request.Text, gtranslate.TranslationParams{
		From: "auto",
		To:   request.To,
	})
	if err != nil {
		sendErrorResponse(w, "Translation failed", http.StatusInternalServerError)
		return
	}

	response := TranslateResponse{
		TranslatedText: translated,
		Status:         true,
		Message:        "",
	}

	sendJSONResponse(w, response, http.StatusOK)
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := TranslateResponse{
		Status:  false,
		Message: message,
	}
	sendJSONResponse(w, response, statusCode)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/translate", TranslateHandler)

	c := cors.Default().Handler(mux)

	fmt.Println("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", c))
}
