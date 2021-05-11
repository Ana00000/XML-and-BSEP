package service

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
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