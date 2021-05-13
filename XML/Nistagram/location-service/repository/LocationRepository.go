package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"gorm.io/gorm"
)

type LocationRepository struct {
	Database * gorm.DB
}

func (repo * LocationRepository) CreateLocation(location *model.Location) error {
	result := repo.Database.Create(location)
	fmt.Print(result)
	return nil
}
