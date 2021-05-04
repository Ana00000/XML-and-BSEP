package model

type SinglePost struct {
	Post
	//PostContent postContentPath.SinglePostContent `json:"post_content" gorm:"foreignKey:SinglePostId"`
}
