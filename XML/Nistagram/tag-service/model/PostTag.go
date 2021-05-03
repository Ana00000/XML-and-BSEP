package model

import (
	postPath "../../post-service/model"
)

type PostTag struct {
	Tag
	Posts []postPath.Post `gorm:"many2many:post_tag_posts"`
}
