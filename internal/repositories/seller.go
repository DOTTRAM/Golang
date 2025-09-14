package repositories

import (
	"Webserver/internal/models"
	"context"
	"errors"
	"strconv"
)

type SellerRepository struct {
	// Убрать зависимость от handlers!
	storage []*models.Seller
}

func NewSellerRepository() *SellerRepository {
	return &SellerRepository{
		storage: make([]*models.Seller, 0),
	}
}

func (r *SellerRepository) Create(ctx context.Context, id int, name string) (*models.Seller, error) {
	seller := &models.Seller{
		Id:   id,
		Name: name,
	}

	r.storage = append(r.storage, seller)
	return seller, nil
}

func (r *SellerRepository) GetByID(ctx context.Context, id string) (*models.Seller, error) {
	// Конвертируем string ID в int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	for _, seller := range r.storage {
		if seller.Id == intID {
			return seller, nil
		}
	}

	return nil, nil // Seller not found
}

func (r *SellerRepository) GetAll(ctx context.Context) ([]*models.Seller, error) {
	if len(r.storage) == 0 {
		return nil, errors.New("no sellers found")
	}
	return r.storage, nil
}

func (r *SellerRepository) Update(ctx context.Context, seller *models.Seller) error {
	if seller == nil {
		return errors.New("seller cannot be nil")
	}

	for i, existingSeller := range r.storage {
		if existingSeller.Id == seller.Id {
			r.storage[i] = seller // Обновляем продавца
			return nil
		}
	}

	return errors.New("seller not found")
}

func (r *SellerRepository) Delete(ctx context.Context, id string) error {
	// Конвертируем string ID в int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	for i, seller := range r.storage {
		if seller.Id == intID {
			// Удаляем элемент из slice
			r.storage = append(r.storage[:i], r.storage[i+1:]...)
			return nil
		}
	}

	return errors.New("seller not found")
}

// Дополнительные методы для обратной совместимости
func (r *SellerRepository) SearchbyId(ctx context.Context, id string) (*models.Seller, error) {
	return r.GetByID(ctx, id)
}

func (r *SellerRepository) SearchAll(ctx context.Context) ([]*models.Seller, error) {
	return r.GetAll(ctx)
}
