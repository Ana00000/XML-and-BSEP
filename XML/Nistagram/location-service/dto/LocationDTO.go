package dto

type LocationDTO struct {
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
	Country string `json:"country"`
	City string `json:"city"`
	StreetName string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
}