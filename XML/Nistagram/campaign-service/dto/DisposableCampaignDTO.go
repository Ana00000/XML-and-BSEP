package dto

type DisposableCampaignDTO struct {
	ExposureTime string `json:"exposure_time" gorm:"not null"`
}