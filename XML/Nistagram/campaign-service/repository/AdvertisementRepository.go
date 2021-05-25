package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"fmt"
	"gorm.io/gorm"
)

type AdvertisementRepository struct {
	Database * gorm.DB
}

func (repo * AdvertisementRepository) CreateAdvertisement(advertisement *model.Advertisement) error {
	result := repo.Database.Create(advertisement)
	fmt.Print(result)
	return nil
}