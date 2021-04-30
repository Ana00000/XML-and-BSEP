package repository

import (
	"../model"
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