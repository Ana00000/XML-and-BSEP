package repository

import (
"../model"
"fmt"
"gorm.io/gorm"
)

type RegisteredUserFollowingsRepository struct {
	Database * gorm.DB
}

func (repo * RegisteredUserFollowingsRepository) CreateRegisteredUserFollowings(registeredUserFollowings *model.RegisteredUserFollowings) error {
	result := repo.Database.Create(registeredUserFollowings)
	fmt.Print(result)
	return nil
}
