package usecase

import "api-go/model"

type ProductUseCase struct {
}

func NewProductuUseCase() *productUseCase {
	return &productUseCase{}
}

func (pu *productUseCase) GetProducts([]model.Product, error) {
	return []model.Product{}, nil
}
