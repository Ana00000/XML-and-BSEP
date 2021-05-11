package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type ClassicUserFollowersRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserFollowersRepository) CreateClassicUserFollowers(classicUserFollowers *model.ClassicUserFollowers) error {
	result := repo.Database.Create(classicUserFollowers)
	fmt.Print(result)
	return nil
}
