package repositories

import (
	"Webserver/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type SellerRepository struct {
}

func NewSellerRepository() *SellerRepository {
	return &SellerRepository{}
}

func (h *SellerRepository) Create(ctx context.Context, id int, name string) (*models.Car, error) {

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
