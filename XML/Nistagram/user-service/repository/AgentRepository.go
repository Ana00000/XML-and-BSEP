package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
"fmt"
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