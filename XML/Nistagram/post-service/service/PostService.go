package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
)

type PostService struct {
	Repo * repository.PostRepository
}

func (service * PostService) CreatePost(post *model.Post) error {
	err := service.Repo.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (service * PostService) UpdatePost(post *dto.PostUpdateDTO) error {
	result := service.Repo.UpdatePost(post)
	fmt.Print(result)
	return nil
}

func (service *PostService) FindByID(ID uuid.UUID) *model.Post {
	post := service.Repo.FindByID(ID)
	return post
}

func (service *PostService) FindAllPostsForUser(ID uuid.UUID) []model.Post {
	posts := service.Repo.FindAllPostsForUser(ID)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *PostService) FindAllFollowingPosts(followings []userModel.ClassicUserFollowings) []model.Post {
	posts := service.Repo.FindAllFollowingPosts(followings)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *PostService) FindAllPublicPostsNotRegisteredUser(allValidUsers []userModel.ClassicUser) []model.Post {
	posts := service.Repo.FindAllPublicPostsNotRegisteredUser(allValidUsers)
	if posts != nil {
		return posts
	}
	return nil
}



