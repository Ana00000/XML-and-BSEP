package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type AgentRepository struct {
	Database * gorm.DB
}

func (repo * AgentRepository) CreateAgent(agent *model.Agent) error {
	result := repo.Database.Create(agent)
	fmt.Print(result)
	return nil
}

func (repo *AgentRepository) UpdateAgentPassword(userId uuid.UUID, password string) error {
	result := repo.Database.Model(&model.Agent{}).Where("id = ?", userId).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}