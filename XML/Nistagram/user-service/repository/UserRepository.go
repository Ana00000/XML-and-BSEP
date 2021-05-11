package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
	"time"
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

func (repo *UserRepository) FindByEmail(email string) *model.User {
	user := &model.User{}
	repo.Database.First(&user, "email = ?", email)
	return user
}

func (repo *UserRepository) FindByID(ID uuid.UUID) *model.User {
	user := &model.User{}
	repo.Database.First(&user, "id = ?", ID)
	return user
}

func (repo *UserRepository) UpdateUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	result := repo.Database.Model(&model.User{}).Where("id = ?", userId).Update("is_confirmed", isConfirmed)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo * UserRepository) UpdateUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, user.DateOfBirth)

	result := repo.Database.Model(&model.User{}).Where("id = ?", user.ID).Update("username", user.Username).Update("phoneNumber", user.PhoneNumber).Update("firstName", user.FirstName).Update("lastName", user.LastName).Update("gender", gender).Update("dateOfBirth", dateOfBirth).Update("website", user.Website).Update("biography", user.Biography)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating profile info")
	return nil
}