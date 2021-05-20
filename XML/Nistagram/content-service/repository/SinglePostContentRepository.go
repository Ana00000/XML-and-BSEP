package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"gorm.io/gorm"
)

type SinglePostContentRepository struct {
	Database * gorm.DB
}

func (repo * SinglePostContentRepository) CreateSinglePostContent(singlePostContent *model.SinglePostContent) error {
	result := repo.Database.Create(singlePostContent)
	fmt.Print(result)
	return nil
}

func (repo *SinglePostContentRepository) FindAll() []model.SinglePostContent {
	var posts []model.SinglePostContent
	repo.Database.Select("*").Find(&posts)
	return posts
}


func (repo *SinglePostContentRepository) FindAllContentsForPosts(allPosts []postsModel.SinglePost) []model.SinglePostContent {
	var contents []model.SinglePostContent
	var allContents = repo.FindAll()

	for i:=0;i<len(allPosts);i++{
		for j:=0; j<len(allContents);j++{
			if allPosts[i].ID == allContents[j].SinglePostId{
				contents = append(contents, allContents[j])
			}
		}

	}

	return contents
}

func (repo *SinglePostContentRepository) FindAllContentsForPost(post *postsModel.SinglePost) []model.SinglePostContent {
	var contents []model.SinglePostContent
	var allContents = repo.FindAll()

	for j:=0; j<len(allContents);j++{
			if post.ID == allContents[j].SinglePostId{
				contents = append(contents, allContents[j])
			}
		}

	return contents
}
