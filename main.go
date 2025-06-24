package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Изменнеия для PULLREquest
// ЗАПУСК РАСЧЁТА ЧЕРЕЗ ТЕЛО POST ЗАПРОСА В JSON ФОРМАТЕ (НЕОБХОДИМО ПЕРЕДАТЬ СРЕЗ {"values": [1, 4, 3, 6]})
type Numbers struct {
	Values []float64 `json:"values"` // Свзяь с JSON
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //Поверка на метод запроса (используется только POST запрос)
		http.Error(w, "Only POST method is supported", http.StatusBadRequest)
		//fmt.Fprintln(w, "Only POST method is supported")
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
	fmt.Fprintln(w, sum)
}

func main() {
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe("localhost:8080", nil)
}
