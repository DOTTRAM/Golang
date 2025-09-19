package repositories

import (
	"Webserver/internal/database"
	"Webserver/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type SellerRepository struct {
	db *database.Database
}

func NewSellerRepository(db *database.Database) *SellerRepository {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	result := r.db.DB.WithContext(ctx).Create(seller)
	if result.Error != nil {
		return nil, result.Error
	}
	return seller, nil
}

func (r *SellerRepository) GetByID(ctx context.Context, id string) (*domain.Seller, error) {
	var seller domain.Seller
	result := r.db.DB.WithContext(ctx).First(&seller, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("seller not found")
		}
		return nil, result.Error
	}
	return &seller, nil
}

func (r *SellerRepository) GetAll(ctx context.Context) ([]*domain.Seller, error) {
	var sellers []*domain.Seller
	result := r.db.DB.WithContext(ctx).Find(&sellers)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(sellers) == 0 {
		return nil, errors.New("no sellers found")
	}
	return sellers, nil
}

func (r *SellerRepository) Update(ctx context.Context, seller *domain.Seller) error {
	result := r.db.DB.WithContext(ctx).Save(seller)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SellerRepository) Delete(ctx context.Context, id string) error {
	result := r.db.DB.WithContext(ctx).Delete(&domain.Seller{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("seller not found")
	}
	return nil
}

// Дополнительные методы для обратной совместимости
func (r *SellerRepository) SearchbyId(ctx context.Context, id string) (*domain.Seller, error) {
	return r.GetByID(ctx, id)
}

func (r *SellerRepository) SearchAll(ctx context.Context) ([]*domain.Seller, error) {
	return r.GetAll(ctx)
}
