package repository

import (
	"../model"
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