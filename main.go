package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Numbers struct {
	Values []float64 `json:"values"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
}

// ввод в формате {"values": [1, 6, 3, 6]}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		statusCode := http.StatusMethodNotAllowed
		responseBody := "Only POST method is supported"
		log.Printf("Error %d: %s", statusCode, responseBody)
		http.Error(w, responseBody, statusCode)
		return
	}

	var nums Numbers
	if err := json.NewDecoder(r.Body).Decode(&nums); err != nil {
		statusCode := http.StatusBadRequest
		responseBody := "Invalid data format"
		log.Printf("Error %d: %s, decode error: %v", statusCode, responseBody, err)
		http.Error(w, responseBody, statusCode)
		return
	}

	sum := 0.0
	for _, v := range nums.Values {
		sum += v
	}

	resp := SumResponse{Sum: sum}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		statusCode := http.StatusInternalServerError
		responseBody := "Failed to encode response"
		log.Printf("Error %d: %s, encode error: %v", statusCode, responseBody, err)
		http.Error(w, responseBody, statusCode)
		return
	}
}

func main() {
	http.HandleFunc("/post", PostHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
