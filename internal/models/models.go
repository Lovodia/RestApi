package models

type Numbers struct {
	Values []float64 `json:"values"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
}

type MultiplyResponse struct {
	Multiply float64 `json:"multiply"`
}
