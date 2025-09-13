package repositories

import (
	"Webserver/internal/handlers"
	"Webserver/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type SellerRepository struct {
	h *handlers.SellerHandler
}

func NewSellerRepository() *SellerRepository {
	return &SellerRepository{}
}

var storage1 = make([]*models.Seller, 0)

func (h *SellerRepository) Create(ctx context.Context, id int, name string) (*models.Seller, error) {
	seller := models.Seller{
		Id:   id,
		Name: name,
	}
	storage1 = append(storage1, &seller)
	return &seller, nil
}

func (h *SellerRepository) Update(c *gin.Context) {

}

func (h *SellerRepository) Delete(c *gin.Context) error {

	return nil
}

func (h *SellerRepository) SearchbyId(c *gin.Context) {

}

func (h *SellerRepository) SearchAll(c *gin.Context) {

}

func (h *SellerRepository) GetByID(ctx context.Context, id string) (interface{}, interface{}) {

}

func (h *SellerRepository) GetAll(ctx context.Context) (interface{}, interface{}) {

}
