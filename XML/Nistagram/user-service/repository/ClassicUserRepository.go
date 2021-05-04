package repository

import (
"../model"
"fmt"
"gorm.io/gorm"
)

type ClassicUserRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserRepository) CreateClassicUser(classicUser *model.ClassicUser) error {
	result := repo.Database.Create(classicUser)
	fmt.Print(result)
	return nil
}
