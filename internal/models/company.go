package models

type Company struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Lemons   int    `json:"lemons"`
	Diamonds int    `json:"diamonds"`
	IsActive bool   `json:"isActive"`
}
