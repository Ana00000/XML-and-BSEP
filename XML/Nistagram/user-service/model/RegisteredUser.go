package model

type RegisteredUser struct {
	ClassicUser
	RegisteredUserCategory RegisteredUserCategory `json:"registered_user_category" gorm:"not null"`
	OfficialDocumentPath string `json:"official_document_path" gorm:"not null"`
}