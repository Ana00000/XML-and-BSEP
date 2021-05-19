package repository

import (
	"fmt"
	"github.com/google/uuid"
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

func (repo *PostRepository) FindAllPosts() []model.Post {
	var posts []model.Post
	repo.Database.Select("*").Find(&posts)
	return posts
}

/*
func (repo *PostRepository) FindAllValidPosts() []model.Post {
	var allPosts []model.Post = repo.FindAllPosts()

	for i:= 0; i< len(allPosts); i++{

	}
}*/

func (repo *PostRepository) FindAllPostsForUser(userId uuid.UUID) []model.Post {
	var posts []model.Post
	repo.Database.Select("*").Where("user_id = ?", userId).Find(&posts)
	return posts
}

func (repo *PostRepository) FindByID(ID uuid.UUID) *model.Post {
	post := &model.Post{}
	if repo.Database.First(&post, "id = ?", ID).RowsAffected == 0 {
		return nil
	}
	return post
}







