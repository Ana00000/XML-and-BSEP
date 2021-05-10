package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
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


func (repo *ClassicUserRepository) UpdateClassicUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ?", userId).Update("is_confirmed", isConfirmed)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}
