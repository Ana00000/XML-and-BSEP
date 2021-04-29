package model

type AdvertisementContent struct {
	Content
	Link string `json:"link" gorm:"not null"`
}