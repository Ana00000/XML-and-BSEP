package model

type RegisteredUser struct {
	User
	Following []RegisteredUser `gorm:"many2many:registered_user_following"`
	Followers []RegisteredUser `gorm:"many2many:registered_user_followers"`
}
