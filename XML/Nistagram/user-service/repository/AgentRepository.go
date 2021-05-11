package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
	"time"
)

type AgentRepository struct {
	Database * gorm.DB
}

func (repo * AgentRepository) CreateAgent(agent *model.Agent) error {
	result := repo.Database.Create(agent)
	fmt.Print(result)
	return nil
}

func (repo * AgentRepository) UpdateAgentProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, user.DateOfBirth)

	result := repo.Database.Model(&model.Agent{}).Where("id = ?", user.ID).Update("username", user.Username).Update("phoneNumber", user.PhoneNumber).Update("firstName", user.FirstName).Update("lastName", user.LastName).Update("gender", gender).Update("dateOfBirth", dateOfBirth).Update("website", user.Website).Update("biography", user.Biography)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating agent profile info")
	return nil
}

func (repo *AgentRepository) UpdateAgentPassword(userId uuid.UUID, password string) error {
	result := repo.Database.Model(&model.Agent{}).Where("id = ?", userId).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}