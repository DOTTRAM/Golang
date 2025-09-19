package domain

import "errors"

type Seller struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Seller) CalcPrice(car *Car) (int, error) {
	if car == nil {
		return 0, errors.New("Car is nil")
	}

	if car.Name == "BMW" {
		return 10000000, nil
	}
	return 0, nil
}
