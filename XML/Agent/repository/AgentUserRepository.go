package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Agent/model"
	"gorm.io/gorm"
)

type AgentUserRepository struct {
	Database * gorm.DB
}

func (repo * AgentUserRepository) CreateAgentUser(agent *model.AgentUser) error {
	result := repo.Database.Create(agent)
	fmt.Print(result)
	return nil
}

func (repo * AgentUserRepository) FindAgentByUserName(userName string) *model.AgentUser {
	agent := &model.AgentUser{}
	if repo.Database.First(&agent, "username = ?", userName).RowsAffected == 0 {
		return nil
	}
	return agent
}

func (repo * AgentUserRepository) FindAgentByEmail(email string) *model.AgentUser {
	agent := &model.AgentUser{}
	if repo.Database.First(&agent, "email = ?", email).RowsAffected == 0 {
		return nil
	}
	return agent
}

func (repo * AgentUserRepository) FindAgentByID(ID uuid.UUID) *model.AgentUser {
	agent := &model.AgentUser{}
	if repo.Database.First(&agent, "id = ?", ID).RowsAffected == 0 {
		return nil
	}
	return agent
}
