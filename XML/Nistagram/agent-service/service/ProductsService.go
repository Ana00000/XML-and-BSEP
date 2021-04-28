package service

import (
	"../repository"
)

type ProductService struct {
	Repo * repository.ProductRepository
}