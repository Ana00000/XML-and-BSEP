package model

type ClassicUser struct {
	RegisteredUser
	IsBlocked bool `json:"is_blocked" gorm:"not null"`
	UserCategory UserCategory `json:"user_category" gorm:"not null"`
	OfficialDocumentPath string `json:"official_document_path" gorm:"not null"`
}