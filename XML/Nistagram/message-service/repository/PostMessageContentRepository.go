package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type PostMessageContentRepository struct {
	Database * gorm.DB
}

func (repo * PostMessageContentRepository) CreatePostMessageContent(postMessageContent *model.PostMessageContent) error {
	result := repo.Database.Create(postMessageContent)
	fmt.Print(result)
	return nil
}