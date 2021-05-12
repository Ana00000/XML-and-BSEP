package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"fmt"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	Database * gorm.DB
}

func (repo * ActivityRepository) CreateActivity(activity *model.Activity) error {
	result := repo.Database.Create(activity)
	fmt.Print(result)
	return nil
}