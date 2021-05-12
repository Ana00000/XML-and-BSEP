package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	Database * gorm.DB
}

func (repo * PostRepository) CreatePost(post *model.Post) error {
	result := repo.Database.Create(post)
	fmt.Print(result)
	return nil
}

func (repo * PostRepository) UpdatePost(post *dto.PostUpdateDTO) error {
	result := repo.Database.Model(&model.Post{}).Where("id = ?", post.ID).Update("description", post.Description).Update("location_id", post.LocationID)
	fmt.Print(result)
	return nil
}
