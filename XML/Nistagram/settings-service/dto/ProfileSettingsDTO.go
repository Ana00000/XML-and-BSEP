package dto

import (
	"../model"
)

type ProfileSettingsDTO struct{
	UserId string `json:"userId"`
	UserVisibility model.UserVisibility `json:"userVisibility"`
	MessageApprovalType model.MessageApprovalType `json:"messageApprovalType"`
}
