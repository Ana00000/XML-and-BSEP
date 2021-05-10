package model

type RegisteredUser struct {
	ClassicUser
	IsBlocked bool `json:"is_blocked" gorm:"not null"`
	RegisteredUserCategory RegisteredUserCategory `json:"registered_user_category" gorm:"not null"`
	OfficialDocumentPath string `json:"official_document_path" gorm:"not null"`
}