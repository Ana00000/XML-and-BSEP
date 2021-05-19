package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type TagRepository struct {
	Database * gorm.DB
}

func (repo * TagRepository) CreateTag(tag *model.Tag) error {
	result := repo.Database.Create(tag)
	fmt.Print(result)
	return nil
}


func (repo *TagRepository) FindTagNameById(ID uuid.UUID) string{
	tag := &model.Tag{}
	if repo.Database.First(&tag, "id = ?", ID).RowsAffected == 0 {
		return ""
	}
	return tag.Name
}
