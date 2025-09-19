package services

import (
	"Webserver/internal/domain"
	"Webserver/internal/repositories"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Бизнес Логика

func Create(ctx context.Context, name string) (*domain.Car, error) {
	if name == "Chery" {
		fmt.Println("Бюджетные машины")
		return nil, errors.New("Не в этой категории автомобилей")
	}

	if name == "BMW" {
		fmt.Println("Премиум класс")

	}

	id := rand.Int()
	car := repositories.Create(id, name)

	return car, nil

}

func GetAll(ctx context.Context) ([]*domain.Car, error) {
	cars, err := repositories.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func GetById(ctx context.Context, id int) (*domain.Car, error) {
	car, err := repositories.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func Update(ctx context.Context, id int, name *string, price *int) (*domain.Car, error) {
	currentCar, _ := GetById(ctx, id)

	if name != nil {
		currentCar.Name = *name
	}

	if price != nil {
		currentCar.Price = *price
	}

	repositories.Update(ctx, currentCar)

	return currentCar, nil
}

func SetYear(ctx context.Context, id int, year int) (*domain.Car, error) {
	if year < 0 {
		return nil, errors.New("че ты пишешь")
	}

	if year > time.Now().Year() {
		return nil, errors.New("че ты пишешь2")
	}

	// поиск машины
	currentCar, _ := GetById(ctx, id)

	// установка значения
	currentCar.Year = year

	// сохрание данных
	repositories.Update(ctx, currentCar)

	return currentCar, nil
}

func Delete(ctx context.Context, id int) (*domain.Car, error) {
	car, err := repositories.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return car, nil
}
