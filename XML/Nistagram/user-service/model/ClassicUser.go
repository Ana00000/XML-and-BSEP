package model

type ClassicUser struct {
	User
	IsDeleted bool `json:"is_deleted" gorm:"not null"`
}
