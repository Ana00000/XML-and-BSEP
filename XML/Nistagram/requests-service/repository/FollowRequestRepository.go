package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"gorm.io/gorm"
)

type FollowRequestRepository struct {
	Database * gorm.DB
}

func (repo * FollowRequestRepository) CreateFollowRequest(followRequest *model.FollowRequest) error {
	result := repo.Database.Create(followRequest)
	fmt.Print(result)
	return nil
}

func (repo *FollowRequestRepository) FindById(id uuid.UUID) *model.FollowRequest {
	request := &model.FollowRequest{}
	if repo.Database.First(&request, "id = ?", id).RowsAffected == 0{
		return nil
	}
	return request
}

func (repo * FollowRequestRepository) FindAllFollowRequestsForUser(userId uuid.UUID) []model.FollowRequest{
	var requests []model.FollowRequest
	repo.Database.Select("*").Where("classic_user_id = ?", userId).Find(&requests)
	return requests
}
