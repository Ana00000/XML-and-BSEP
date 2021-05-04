package model

type RegisteredUser struct {
	User
	//Followings []RegisteredUser `gorm:"many2many:registered_user_following"`
	//Followers []RegisteredUser `gorm:"many2many:registered_user_followers"`
	//Campaigns []campaignPath.Campaign `gorm:"many2many:registered_user_campaigns"`
	//InappropriateContentRequest []icrPath.InappropriateContentRequest `gorm:"foreignKey:UserId"`
}
