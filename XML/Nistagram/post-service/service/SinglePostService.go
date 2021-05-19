package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
)

type SinglePostService struct {
	Repo * repository.SinglePostRepository
}

func (service * SinglePostService) CreateSinglePost(singlePost *model.SinglePost) error {
	err := service.Repo.CreateSinglePost(singlePost)
	if err != nil {
		return err
	}
	return nil
}

func (service *SinglePostService) FindByID(ID uuid.UUID) *model.Post {
	post := service.Repo.FindByID(ID)
	return post
}

func (service *SinglePostService) FindAllPostsForUser(ID uuid.UUID) []model.Post {
	posts := service.Repo.FindAllPostsForUser(ID)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllFollowingPosts(followings []userModel.ClassicUserFollowings) []model.Post {
	posts := service.Repo.FindAllFollowingPosts(followings)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllPublicPostsNotRegisteredUser(allValidUsers []userModel.ClassicUser) []model.Post {
	posts := service.Repo.FindAllPublicPostsNotRegisteredUser(allValidUsers)
	if posts != nil {
		return posts
	}
	return nil
}
