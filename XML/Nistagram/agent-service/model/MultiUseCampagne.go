package model

import "time"

type MultiUseCampaign struct {
	Campaign
	ExpiryTime time.Time `json:"expiry_time" gorm:"not null"`
	Frequency int `json:"expiry_time" gorm:"not null"`
}