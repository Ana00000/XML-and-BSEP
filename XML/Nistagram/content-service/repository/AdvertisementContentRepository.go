package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type AdvertisementContentRepository struct {
	Database * gorm.DB
}

func (repo * AdvertisementContentRepository) CreateAdvertisementContent(advertisementContent *model.AdvertisementContent) error {
	result := repo.Database.Create(advertisementContent)
	fmt.Print(result)
	return nil
}