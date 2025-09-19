package services

import (
	"Webserver/internal/domain"
	"Webserver/internal/repositories"
	"context"
	"errors"
)

type SellerService struct {
	repo *repositories.SellerRepository
}

// NewSellerService Constructor
func NewSellerService(repo *repositories.SellerRepository) *SellerService {
	return &SellerService{repo: repo}
}

func (s *SellerService) Create(ctx context.Context, name string) (*domain.Seller, error) {
	// Валидация входных данных
	if name == "" {
		return nil, errors.New("name is required")
	}

	if len(name) < 2 || len(name) > 100 {
		return nil, errors.New("name must be between 2 and 100 characters")
	}

	// Создаем объект seller
	seller := &domain.Seller{
		Name: name,
	}

	// Сохраняем в репозитории (передаем только контекст и объект)
	createdSeller, err := s.repo.Create(ctx, seller)
	if err != nil {
		return nil, err
	}

	return createdSeller, nil
}

func (s *SellerService) GetByID(ctx context.Context, id string) (*domain.Seller, error) {
	// Валидация ID
	if id == "" {
		return nil, errors.New("id is required")
	}

	seller, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *SellerService) GetAll(ctx context.Context) ([]*domain.Seller, error) {
	sellers, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err // Возвращаем оригинальную ошибку
	}

	return sellers, nil
}

func (s *SellerService) Update(ctx context.Context, id string, name string) (*domain.Seller, error) {
	// Валидация
	if id == "" {
		return nil, errors.New("id is required")
	}
	if name == "" {
		return nil, errors.New("name is required")
	}
	if len(name) < 2 || len(name) > 100 {
		return nil, errors.New("name must be between 2 and 100 characters")
	}

	// Получаем существующего продавца
	existingSeller, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Обновляем данные
	existingSeller.Name = name

	err = s.repo.Update(ctx, existingSeller)
	if err != nil {
		return nil, err
	}

	return existingSeller, nil
}

func (s *SellerService) Delete(ctx context.Context, id string) error {
	// Валидация
	if id == "" {
		return errors.New("id is required")
	}

	// Проверяем существование
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Удаляем
	return s.repo.Delete(ctx, id)
}
