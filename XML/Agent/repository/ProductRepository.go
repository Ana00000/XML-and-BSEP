package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Agent/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Database *gorm.DB
}

func (repo *ProductRepository) CreateProduct(product *model.Product) error {
	result := repo.Database.Create(product)
	fmt.Print(result)
	return nil
}
