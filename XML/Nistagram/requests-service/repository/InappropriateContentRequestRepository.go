package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"gorm.io/gorm"
)

type InappropriateContentRequestRepository struct {
	Database *gorm.DB
}

func (repo *InappropriateContentRequestRepository) CreateInappropriateContentRequest(inappropriateContentRequest *model.InappropriateContentRequest) error {
	result := repo.Database.Create(inappropriateContentRequest)
	fmt.Print(result)
	return nil
}

func (repo * InappropriateContentRequestRepository) CreateICR(icr *model.InappropriateContentRequest) error {
	result := repo.Database.Create(icr)
	fmt.Print(result)
	return nil
}