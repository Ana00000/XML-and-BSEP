package dto

type MultiUseCampaignDTO struct {
	ExposureTime string `json:"exposure_time" gorm:"not null"`
	ExpiryTime string `json:"expiry_time" gorm:"not null"`
	Frequency int `json:"expiry_time" gorm:"not null"`
}