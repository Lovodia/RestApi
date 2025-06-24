package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Numbers struct {
	Values []float64 `json:"values"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
}

//ФОРМАТ ВВОДА {"values": [1, 4, 3, 6]}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default()

	if r.Method != http.MethodPost {
		logger.Warn("Unsupported HTTP method",
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
		)
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var nums Numbers
	if err := json.NewDecoder(r.Body).Decode(&nums); err != nil {
		logger.Warn("JSON decode error",
			"error", err,
			"remote_addr", r.RemoteAddr,
		)
		http.Error(w, "Invalid data format", http.StatusBadRequest)
		return
	}

	sum := 0.0
	for _, v := range nums.Values {
		sum += v
	}

	logger.Info("Sum calculated",
		"values", nums.Values,
		"sum", sum,
		"remote_addr", r.RemoteAddr,
	)

	resp := SumResponse{Sum: sum}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/post", PostHandler)
	slog.Info("Server started on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
