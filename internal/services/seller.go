package services

import (
	"Webserver/internal/models"
	"Webserver/internal/repositories"
	"context"
	"errors"
	"math/rand"

	_ "github.com/gin-gonic/gin"
)

type SellerService struct {
	repo *repositories.SellerRepository
}

// NewSellerHandler Constructor
func NewSellerService(repo *repositories.SellerRepository) *SellerService {
	return &SellerService{repo: repo}
}

func (s *SellerService) Create(ctx context.Context, name string) (*models.Seller, error) {
	// Валидация входных данных
	if name == "" {
		return nil, errors.New("name is required")
	}

	if len(name) < 2 || len(name) > 100 {
		return nil, errors.New("name must be between 2 and 100 characters")
	}

	// Сохраняем в репозитории
	id := rand.Int()
	seller, err := s.repo.Create(ctx, id, name)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *SellerService) GetByID(ctx context.Context, id string) (*models.Seller, error) {
	// Валидация ID
	if id == "" {
		return nil, errors.New("id is required")
	}

	seller, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if seller == nil {
		return nil, errors.New("seller not found")
	}

	return seller, nil
}

func (s *SellerService) GetAll(ctx context.Context) ([]*models.Seller, error) {
	sellers, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("Введено пустое значение")
	}

	return sellers, nil
}

func (s *SellerService) Update(ctx context.Context, id string, name string) (*models.Seller, error) {
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

	// Проверяем существование продавца
	existingSeller, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingSeller == nil {
		return nil, errors.New("seller not found")
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
	existingSeller, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingSeller == nil {
		return errors.New("seller not found")
	}
	// Удаляем
	return s.repo.Delete(ctx, id)
}
