package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Email    string `json:"email"`
	Datetime string `json:"datetime"`
	Github   string `json:"github"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")


	response := Response{
		Email:    "chiamaevans10@gmail.com",
		Datetime: time.Now().Format(time.RFC3339),
		Github:   "https://github.com/evans_sudo/PublicApi",
	}

	json.NewEncoder(w).Encode(response)
	
}

func main() {
	http.HandleFunc("/api/info", handler)
	http.ListenAndServe(":8080", nil)
}
