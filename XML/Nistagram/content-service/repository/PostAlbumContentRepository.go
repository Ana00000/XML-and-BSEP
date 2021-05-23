package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"fmt"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"gorm.io/gorm"
)

type PostAlbumContentRepository struct {
	Database * gorm.DB
}

func (repo * PostAlbumContentRepository) CreatePostAlbumContent(postAlbumContent *model.PostAlbumContent) error {
	result := repo.Database.Create(postAlbumContent)
	fmt.Print(result)
	return nil
}

func (repo *PostAlbumContentRepository) FindAll() []model.PostAlbumContent {
	var contents []model.PostAlbumContent
	repo.Database.Select("*").Find(&contents)
	return contents
}

func (repo *PostAlbumContentRepository) FindAllContentsForPostAlbums(albums []postsModel.PostAlbum) []model.PostAlbumContent {
	var contents []model.PostAlbumContent
	var allContents = repo.FindAll()

	for i:=0;i<len(albums);i++{
		for j:=0; j<len(allContents);j++{
			if albums[i].ID == allContents[j].PostAlbumId{
				contents = append(contents, allContents[j])
			}
		}

	}

	return contents
}
