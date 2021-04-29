package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database * gorm.DB
}

func (repo * UserRepository) CreateUser(user *model.User) error {
	result := repo.Database.Create(user)
	fmt.Print(result)
	return nil
}