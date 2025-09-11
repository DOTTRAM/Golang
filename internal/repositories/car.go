package repositories

import (
	"Webserver/internal/models"
	"context"
	"fmt"
)

var storage = make([]*models.Car, 0)

func Create(id int, name string) *models.Car {
	car := models.Car{
		Id:   id,
		Name: name,
	}

	storage = append(storage, &car) // Добавляем в сторадж
	return &car

}

func GetAll(ctx context.Context) ([]*models.Car, error) {
	return storage, nil
}

func GetById(ctx context.Context, id int) (*models.Car, error) {
	var car *models.Car

	for _, storagedCar := range storage {
		if id == storagedCar.Id {
			car = storagedCar
		}
	}

	return car, nil
}

func Update(ctx context.Context, car *models.Car) error {
	for i, m := range storage {
		if car.Id == m.Id {
			storage[i] = car
		}
	}
	return nil
}

func Delete(ctx context.Context, id int) (*models.Car, error) {

	originalLength := len(storage)

	for i := 0; i < originalLength; i++ {
		if storage[i].Id == id {
			// Сохраняем найденную машину
			// Удаляем элемент из слайса
			storage = append(storage[:i], storage[i+1:]...)

			// НЕМЕДЛЕННО возвращаем успешный результат

		}
	}

	// Если дошли сюда - машина не найдена
	return nil, fmt.Errorf("car with id %d not found", id)
}
