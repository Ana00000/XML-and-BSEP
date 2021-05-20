package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type ClassicUserCloseFriendsRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserCloseFriendsRepository) CreateClassicUserCloseFriends(classicUserCloseFriends *model.ClassicUserCloseFriends) error {
	result := repo.Database.Create(classicUserCloseFriends)
	fmt.Print(result)
	return nil
}