package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
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

func (service *SinglePostService) FindByID(ID uuid.UUID) *model.SinglePost {
	post := service.Repo.FindByID(ID)
	return post
}

func (service *SinglePostService) FindAllPostsForUser(ID uuid.UUID) []model.SinglePost {
	posts := service.Repo.FindAllPostsForUser(ID)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllFollowingPosts(followings []dto.ClassicUserFollowingsFullDTO) []model.SinglePost {
	posts := service.Repo.FindAllFollowingPosts(followings)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllPublicAndFriendsPostsValid(allValidUsers []dto.ClassicUserDTO) []model.SinglePost {
	posts := service.Repo.FindAllPublicAndFriendsPostsValid(allValidUsers)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllPostsByIds(postsIds []uuid.UUID) []model.SinglePost {
	posts := service.Repo.FindAllPostsByIds(postsIds)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllPublicPostsByIds(postsIds []uuid.UUID, allValidUsers []userModel.ClassicUser) []model.SinglePost {
	posts := service.Repo.FindAllPublicPostsByIds(postsIds,allValidUsers)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllPostIdsWithLocationId(locationId uuid.UUID) []model.SinglePost {
	posts := service.Repo.FindAllPostIdsWithLocationId(locationId)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllPublicAndFriendsPosts(postsList []model.SinglePost, alValidUsers []userModel.ClassicUser) []model.SinglePost {
	posts := service.Repo.FindAllPublicAndFriendsPosts(postsList, alValidUsers)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostService) FindAllPostsForUsers(users []userModel.ClassicUser) []model.SinglePost {
	posts := service.Repo.FindAllPostsForUsers(users)
	if posts != nil {
		return posts
	}
	return nil
}