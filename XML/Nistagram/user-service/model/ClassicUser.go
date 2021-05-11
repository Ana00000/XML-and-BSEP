package model

type ClassicUser struct {
	User
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
}
