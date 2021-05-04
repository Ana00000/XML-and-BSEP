package model

type CommentTag struct {
	Tag
	//Comments []commentPath.Comment `gorm:"many2many:comment_tag_comments"`
}