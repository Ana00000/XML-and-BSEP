package model

import (
	postContentPath "../../content-service/model"
)

type PostAlbum struct {
	Post
	PostContents []postContentPath.PostAlbumContent `json:"post_contents" gorm:"foreignKey:PostAlbumId"`
}
