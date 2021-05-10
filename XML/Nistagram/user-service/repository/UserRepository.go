package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
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


func (repo * UserRepository) FindAllUsers() []model.User{
	var users []model.User
	repo.Database.Select("*").Find(&users)
	return users
}

func (repo *UserRepository) FindByUserName(userName string) *model.User {
	user := &model.User{}
	repo.Database.First(&user, "username = ?", userName)
	return user
}

func (repo *UserRepository) UpdateUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	result := repo.Database.Model(&model.User{}).Where("id = ?", userId).Update("is_confirmed", isConfirmed)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}