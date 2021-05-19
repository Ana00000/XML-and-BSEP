package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type ClassicUserRepository struct {
	Database *gorm.DB
}

func (repo *ClassicUserRepository) CreateClassicUser(classicUser *model.ClassicUser) error {
	result := repo.Database.Create(classicUser)
	fmt.Print(result)
	return nil
}

func (repo *ClassicUserRepository) UpdateClassicUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ? and is_deleted = ?", userId, false).Update("is_confirmed", isConfirmed)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *ClassicUserRepository) UpdateClassicUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ? and is_confirmed = ? and is_deleted = ?", user.ID, true, false)
	result.Update("username", user.Username)

	fmt.Println(result.RowsAffected)
	result.Update("phone_number", user.PhoneNumber)
	fmt.Println(result.RowsAffected)
	result.Update("first_name", user.FirstName)
	fmt.Println(result.RowsAffected)
	result.Update("last_name", user.LastName)
	fmt.Println(result.RowsAffected)
	result.Update("gender", gender)
	fmt.Println(result.RowsAffected)
	result.Update("date_of_birth", user.DateOfBirth)
	fmt.Println(result.RowsAffected)
	result.Update("website", user.Website)
	fmt.Println(result.RowsAffected)
	result.Update("biography", user.Biography)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating profile info")
	return nil
}

func (repo *ClassicUserRepository) UpdateClassicUserPassword(userId uuid.UUID, salt string, password string) error {
	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ? and is_confirmed = ? and is_deleted = ?", userId, true, false).Update("salt", salt).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

//SELECTED USER FOR CLASSIC USER VIEW (IF ADMIN IT SHOUD BE ABLE TO SEE DELETED USERS?)

func (repo *ClassicUserRepository) FindSelectedUserById(id uuid.UUID) *dto.SelectedUserDTO {
	user := &model.ClassicUser{}

	if repo.Database.First(&user, "id = ? and is_confirmed = ? and is_deleted = ?", id, true, false).RowsAffected == 0 {
		return nil
	}

	return repo.ConvertFromUserToSelectedUserDTO(user)
}

func (repo *ClassicUserRepository) ConvertFromUserToSelectedUserDTO(user *model.ClassicUser) *dto.SelectedUserDTO {
	userDTO := &dto.SelectedUserDTO{}
	userDTO.Username = user.Username
	userDTO.FirstName = user.FirstName
	userDTO.LastName = user.LastName
	userDTO.Biography = user.Biography
	userDTO.Website = user.Website
	return userDTO
}

func (repo *ClassicUserRepository) FindClassicUserByUserName(userName string) *model.ClassicUser {
	user := &model.ClassicUser{}
	if repo.Database.First(&user, "username = ? and is_confirmed = ? and is_deleted = ? ", userName, true, false).RowsAffected == 0 {
		return nil
	}
	return user
}

func (repo *ClassicUserRepository) FindAllUsersButLoggedIn(userId uuid.UUID) []model.ClassicUser {

	var users []model.ClassicUser
	repo.Database.Select("*").Where("id != ? and is_confirmed = ? and is_deleted = ? ", userId, true, false).Find(&users)
	return users
}