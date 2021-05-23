package repository

import (
	"fmt"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type PostAlbumTagPostAlbumsRepository struct {
	Database * gorm.DB
}

func (repo * PostAlbumTagPostAlbumsRepository) CreatePostAlbumTagPostAlbums(postAlbumTagPostAlbums *model.PostAlbumTagPostAlbums) error {
	result := repo.Database.Create(postAlbumTagPostAlbums)
	fmt.Print(result)
	return nil
}

func (repo *PostAlbumTagPostAlbumsRepository) FindAll() []model.PostAlbumTagPostAlbums {
	var tags []model.PostAlbumTagPostAlbums
	repo.Database.Select("*").Find(&tags)
	return tags
}

func (repo *PostAlbumTagPostAlbumsRepository) FindAllTagsForPostAlbumTagPostAlbums(albums []postsModel.PostAlbum) []model.PostAlbumTagPostAlbums {
	var tags []model.PostAlbumTagPostAlbums
	var allTags = repo.FindAll()

	for i:=0;i<len(albums);i++{
		for j:=0; j<len(allTags);j++{
			if albums[i].ID == allTags[j].PostAlbumId{
				tags = append(tags, allTags[j])
			}
		}

	}
	return tags
}

func (repo *PostAlbumTagPostAlbumsRepository) FindAllTagsForPostAlbum(album *postsModel.PostAlbum) []model.PostAlbumTagPostAlbums {
	var tags []model.PostAlbumTagPostAlbums
	var allTags = repo.FindAll()

	for j:=0; j<len(allTags);j++{
		if album.ID == allTags[j].PostAlbumId{
			tags = append(tags, allTags[j])
		}
	}

	return tags
}