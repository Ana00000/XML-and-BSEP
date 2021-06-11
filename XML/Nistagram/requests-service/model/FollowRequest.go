package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowRequest struct {
	ID                  uuid.UUID           `json:"id"`
	ClassicUserId       uuid.UUID           `json:"classic_user_id" gorm:"not null"`
	FollowerUserId      uuid.UUID           `json:"follower_user_id" gorm:"not null"`
	FollowRequestStatus FollowRequestStatus `json:"follow_request_status" gorm:"not null"`
}

func (followRequest *FollowRequest) BeforeCreate(scope *gorm.DB) error {
	followRequest.ID = uuid.New()
	return nil
}
