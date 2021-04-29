package dto

type CampaignDTO struct {
	ExposureTime string `json:"exposure_time" gorm:"not null"`
}