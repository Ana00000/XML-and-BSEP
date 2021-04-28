package service

import (
	"../model"
	"../repository"
)

type ProductService struct {
	Repo * repository.ProductRepository
}

func (service * ProductService) CreateProduct(product *model.Product) error {
	service.Repo.CreateProduct(product)
	return nil
}