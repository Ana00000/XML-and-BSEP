package dto

import "github.com/google/uuid"

type LocationDTO struct {
	ID uuid.UUID `json:"id"`
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
	Country string `json:"country"`
	City string `json:"city"`
	StreetName string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
}

