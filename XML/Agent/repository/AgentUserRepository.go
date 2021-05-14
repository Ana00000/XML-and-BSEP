package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Agent/model"
	"fmt"
	"gorm.io/gorm"
)

type AgentUserRepository struct {
	Database * gorm.DB
}

func (repo * AgentUserRepository) CreateAgentUser(user *model.AgentUser) error {
	result := repo.Database.Create(user)
	fmt.Print(result)
	return nil
}