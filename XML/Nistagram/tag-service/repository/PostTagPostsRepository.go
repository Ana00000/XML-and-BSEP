package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type PostTagPostsRepository struct {
	Database * gorm.DB
}

func (repo * PostTagPostsRepository) CreatePostTagPosts(postTagPosts *model.PostTagPosts) error {
	result := repo.Database.Create(postTagPosts)
	fmt.Print(result)
	return nil
}

func (repo *PostTagPostsRepository) FindAll() []model.PostTagPosts {
	var tags []model.PostTagPosts
	repo.Database.Select("*").Find(&tags)
	return tags
}

func (repo *PostTagPostsRepository) FindTagById(ID uuid.UUID) model.Tag{
	var tag model.Tag
	repo.Database.Select("*").Where("id=?", ID).Find(&tag)
	return tag
}

//FindAllTagsForPostsTagPosts
func (repo *PostTagPostsRepository) FindAllTagsForPostsTagPosts(allPosts []dto.SinglePostDTO) []model.PostTagPosts {
	var tags []model.PostTagPosts
	var allTags = repo.FindAll()

	for i:=0;i<len(allPosts);i++{
		for j:=0; j<len(allTags);j++{
			if allPosts[i].ID == allTags[j].PostId{
				tags = append(tags, allTags[j])
			}
		}

	}
	return tags
}

func (repo *PostTagPostsRepository) FindAllTagsForPosts(allPosts []dto.SinglePostDTO) []model.Tag {
	var tags []model.Tag
	var allTags = repo.FindAll()

	for i:=0;i<len(allPosts);i++{
		for j:=0; j<len(allTags);j++{
			if allPosts[i].ID == allTags[j].PostId && !ExsistInListTags(repo.FindTagById(allTags[j].TagId),tags){
				tags = append(tags, repo.FindTagById(allTags[j].TagId))
			}
		}

	}
	return tags
}

func ExsistInListPostTagPosts(postTagPosts model.PostTagPosts, allPostTagPosts []model.PostTagPosts) bool{
	for i := 0; i < len(allPostTagPosts); i++ {
		if allPostTagPosts[i].ID == postTagPosts.ID{
			return true
		}
	}
	return false
}


func ExsistInListTags(tag model.Tag, allTags []model.Tag) bool{
	for i := 0; i < len(allTags); i++ {
		if allTags[i].ID == tag.ID{
			return true
		}
	}
	return false
}


func (repo *PostTagPostsRepository) FindAllTagsForPost(post *dto.SinglePostDTO) []model.PostTagPosts {
	var tags []model.PostTagPosts
	var allTags = repo.FindAll()

	for j:=0; j<len(allTags);j++{
			if post.ID == allTags[j].PostId && !ExsistInListPostTagPosts(allTags[j],tags){
				tags = append(tags, allTags[j])
			}
		}

	return tags
}

func (repo *PostTagPostsRepository) FindAllPostIdsWithTagId(tagId uuid.UUID) []uuid.UUID {
	var postIds []uuid.UUID
	var allPosts = repo.FindAll()

	for i:=0; i <len(allPosts);i++{
		if allPosts[i].TagId == tagId{
			postIds = append(postIds, allPosts[i].PostId)
		}
	}

	return postIds
}




