package model

import (
	postPath "../../post-service/model"
	settingsPath "../../settings-service/model"
	storyPath "../../story-service/model"
)

type ClassicUser struct {
	RegisteredUser
	Stories         []storyPath.Story          `json:"stories" gorm:"foreignKey:UserId"`
	StoryHighlights []storyPath.StoryHighlight `json:"story_highlights" gorm:"foreignKey:UserId"`
	Posts           []postPath.Post            `json:"posts" gorm:"foreignKey:UserID"`
	PostCollections []postPath.PostCollection  `json:"post_collections" gorm:"foreignKey:UserID"`
	Activities      []postPath.Activity        `json:"activities" gorm:"foreignKey:UserID"`
	Comments []postPath.Comment `json:"comments" gorm:"foreignKey:UserID"`
	IsBlocked bool `json:"is_blocked" gorm:"not null"`
	UserCategory UserCategory `json:"user_category" gorm:"not null"`
	OfficialDocumentPath string `json:"official_document_path" gorm:"not null"`
	Settings settingsPath.ProfileSettings `json:"settings" gorm:"foreignKey:UserId"`
}