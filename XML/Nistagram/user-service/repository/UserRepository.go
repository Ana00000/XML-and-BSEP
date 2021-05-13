package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
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
	if repo.Database.First(&user, "username = ?", userName).RowsAffected == 0{

		return nil

	}
	return user
}

func (repo *UserRepository) FindByEmail(email string) *model.User {
	user := &model.User{}
	if repo.Database.First(&user, "email = ?", email).RowsAffected == 0{
		return nil
	}
	return user
}

func (repo *UserRepository) FindByID(ID uuid.UUID) *model.User {
	user := &model.User{}
	if repo.Database.First(&user, "id = ?", ID).RowsAffected == 0{
		return nil
	}
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

	result := repo.Database.Model(&model.User{}).Where("id = ?", user.ID)
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

func (repo *UserRepository) UpdateUserPassword(userId uuid.UUID, password string) error {
	result := repo.Database.Model(&model.User{}).Where("id = ?", userId).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo * UserRepository) FindAllFollowersInfoForUser(followers []model.ClassicUserFollowers) []model.User{

	var followerUsers []model.User

	for i := 0; i < len(followers); i++{
		var users model.User
		repo.Database.Select("*").Where("id=?", followers[i].FollowerUserId).Find(&users)
		followerUsers = append(followerUsers, users)
	}

	return followerUsers
}

func (repo * UserRepository) FindAllUsersButLoggedIn(userId uuid.UUID) []model.User{

	var users []model.User
	repo.Database.Select("*").Where("id != ?", userId).Find(&users)
	return users
}

func (repo * UserRepository) FindAllPublicUsers(publicUsers []uuid.UUID) []model.User{

	var users []model.User
	for i := 0; i < len(publicUsers); i++ {
		var user model.User
		repo.Database.Select("*").Where("id != ?", publicUsers[i]).Find(&user)
		users = append(users, user)
	}




	return users
}