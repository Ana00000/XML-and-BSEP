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

func (repo *TagRepository) FindAll() []model.Tag {
	var tags []model.Tag
	repo.Database.Select("*").Find(&tags)
	return tags
}

func (repo *TagRepository) FindAllHashTags() []model.Tag {
	var tags []model.Tag
	repo.Database.Select("*").Where("tag_type = ?", 1).Find(&tags)
	return tags
}


func (repo *TagRepository) FindTagByName(name string) *model.Tag {
	tag := &model.Tag{}
	if repo.Database.First(&tag, "name = ?", name).RowsAffected == 0 {
		return nil
	}
	return tag
}

func (repo *TagRepository) FindTagByType(tagType string) *model.Tag {
	tag := &model.Tag{}
	if repo.Database.First(&tag, "tag_type = ?", tagType).RowsAffected == 0 {
		return nil
	}
	return tag
}

func (repo * TagRepository) FindTagForId(tagId uuid.UUID) model.Tag{
	var tag model.Tag
	repo.Database.Select("*").Where("id = ?", tagId).Find(&tag)
	return tag
}