package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
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


func (repo *PostRepository) FindByID(ID uuid.UUID) *model.Post {
	post := &model.Post{}
	if repo.Database.First(&post, "id = ?", ID).RowsAffected == 0 {
		return nil
	}
	return post
}

// USED WHEN CLICKING ON A SELECTED USER (YOU CAN SELECT FROM A LIST OF ONLY VALID USERS)
func (repo *PostRepository) FindAllPostsForUser(userId uuid.UUID) []model.Post {
	var posts []model.Post
	repo.Database.Select("*").Where("user_id = ?", userId).Find(&posts)
	return posts
}


// FIND ALL VALID POSTS THAT LOGGED IN USER FOLLOWS
func (repo *PostRepository) FindAllFollowingPosts(followings []userModel.ClassicUserFollowings) []model.Post {
	var allPosts = repo.FindAllPosts()
	var allFollowingPosts []model.Post

	for i:= 0; i< len(allPosts); i++{
		for j := 0; j < len(followings); i++{
			if allPosts[i].UserID == followings[i].FollowingUserId{
				allFollowingPosts = append(allFollowingPosts, allPosts[i])
			}
		}
	}
	return allFollowingPosts
}







