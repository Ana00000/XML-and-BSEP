package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type RegisteredUserRepository struct {
	Database * gorm.DB
}

func (repo * RegisteredUserRepository) CreateRegisteredUser(registeredUser *model.RegisteredUser) error {
	result := repo.Database.Create(registeredUser)
	fmt.Print(result)
	return nil
}