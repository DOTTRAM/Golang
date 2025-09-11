package models

type Car struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}
