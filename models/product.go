package models

type Product struct {
	ID          *int    `json:"id"`
	Description string  `json:"description"`
	UN          string  `json:"un"`
	Value       float64 `json:"value"`
}

type Products struct {
	Data []Product `json:"data"`
}
