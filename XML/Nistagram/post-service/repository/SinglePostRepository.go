package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type SinglePostRepository struct {
	Database * gorm.DB
}

func (repo * SinglePostRepository) CreateSinglePost(singlePost *model.SinglePost) error {
	result := repo.Database.Create(singlePost)
	fmt.Print(result)
	return nil
}

func (repo *SinglePostRepository) FindAllPosts() []model.SinglePost {
	var posts []model.SinglePost
	repo.Database.Select("*").Find(&posts)
	return posts
}

func (repo *SinglePostRepository) FindByID(ID uuid.UUID) *model.SinglePost {
	post := &model.SinglePost{}
	if repo.Database.First(&post, "id = ? and is_deleted = ?", ID, false).RowsAffected == 0 {
		return nil
	}
	return post
}

// USED WHEN CLICKING ON A SELECTED USER (YOU CAN SELECT FROM A LIST OF ONLY VALID USERS)
func (repo *SinglePostRepository) FindAllPostsForUser(userId uuid.UUID) []model.SinglePost {
	var posts []model.SinglePost
	repo.Database.Select("*").Where("user_id = ? and is_deleted = ?", userId, false).Find(&posts)
	return posts
}


// FIND ALL NOT DELETED VALID POSTS THAT LOGGED IN USER FOLLOWS
func (repo *SinglePostRepository) FindAllFollowingPosts(followings []userModel.ClassicUserFollowings) []model.SinglePost {
	var allPosts = repo.FindAllPosts()
	var allFollowingPosts []model.SinglePost

	for i:= 0; i< len(allPosts); i++{
		for j := 0; j < len(followings); j++{
			if (allPosts[i].UserID == followings[j].FollowingUserId) && (allPosts[i].IsDeleted == false){
				allFollowingPosts = append(allFollowingPosts, allPosts[i])
			}
		}
	}
	return allFollowingPosts
}

func (repo *SinglePostRepository) FindAllPublicPostsNotRegisteredUser(allValidUsers []userModel.ClassicUser) []model.SinglePost {
	var allPosts = repo.FindAllPosts()
	var allPublicPosts []model.SinglePost

	for i:=0;i<len(allPosts);i++{
		for j:=0; j<len(allValidUsers);j++{
			if allPosts[i].UserID == allValidUsers[j].ID {
				allPublicPosts = append(allPublicPosts, allPosts[i])
			}
		}
	}

	return allPublicPosts
}