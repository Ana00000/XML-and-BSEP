package model

import (
	campaignPath "../../agent-service/model"
	icrPath "../../requests-service/model"
)

type RegisteredUser struct {
	User
	Following []RegisteredUser `gorm:"many2many:registered_user_following"`
	Followers []RegisteredUser `gorm:"many2many:registered_user_followers"`
	Campaigns []campaignPath.Campaign `gorm:"many2many:registered_user_campaigns"`
	InappropriateContentRequest []icrPath.InappropriateContentRequest `json:"inappropriate_content_request" gorm:"foreignKey:UserId"`
}
