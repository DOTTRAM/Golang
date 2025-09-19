package seller

import (
	"Webserver/internal/repositories"
	"context"
)

type GetCarsBySellerUseCase struct {
	r *repositories.SellerRepository
}

func (uc *GetCarsBySellerUseCase) GetCarsBySeller(ctx context.Context) error {
	uc.r.Create()
}
