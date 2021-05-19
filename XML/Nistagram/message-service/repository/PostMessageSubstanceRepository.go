package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"gorm.io/gorm"
)

type PostMessageSubstanceRepository struct {
	Database * gorm.DB
}

func (repo * PostMessageSubstanceRepository) CreatePostMessageSubstance(postMessageSubstance *model.PostMessageSubstance) error {
	result := repo.Database.Create(postMessageSubstance)
	fmt.Print(result)
	return nil
}