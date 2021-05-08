package service

import (
	"../model"
	"../repository"
)

type ProductService struct {
	Repo * repository.ProductRepository
}

func (service * ProductService) CreateProduct(product *model.Product) error {
	err := service.Repo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}