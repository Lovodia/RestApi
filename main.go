package main

import (
	"encoding/json"
	"net/http"
)

// Изменнеия для PULLREquest
// ЗАПУСК РАСЧЁТА ЧЕРЕЗ ТЕЛО POST ЗАПРОСА В JSON ФОРМАТЕ (НЕОБХОДИМО ПЕРЕДАТЬ СРЕЗ {"values": [1, 4, 3, 6]})
type Numbers struct {
	Values []float64 `json:"values"` // Свзяь с JSON
}
type SumResponse struct {
	Sum float64 `json:"sum"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //Поверка на метод запроса (используется только POST запрос)
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed) // 405
		return
	}

	var nums Numbers
	if err := json.NewDecoder(r.Body).Decode(&nums); err != nil { // Декодеривание JSON
		http.Error(w, "invalid data format", http.StatusBadRequest) //Ошибка при декодировании
		return
	}

	sum := 0.0
	for _, v := range nums.Values {
		sum += v
	}
	resp := SumResponse{Sum: sum}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil { //ошибка при декодировании
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe("localhost:8080", nil)
}
