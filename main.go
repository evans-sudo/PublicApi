package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Email    string `json:"email"`
	Datetime string `json:"datetime"`
	Github   string `json:"github"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// CORS headers for all methods
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	// Set Content-Type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Prepare the response data
	response := Response{
		Email:    "chiamaevans10@gmail.com",
		Datetime: time.Now().Format(time.RFC3339),
		Github:   "https://github.com/evans-sudo/PublicApi",
	}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
	}
}


func main() {
	http.HandleFunc("/api/info", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
