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

func (repo * FollowRequestRepository) FindAllFollowRequestsForFollower(userId uuid.UUID) []model.FollowRequest{
	var requests []model.FollowRequest


	repo.Database.Select("*").Where("follower_user_id = ?", userId).Find(&requests)
	return requests
}


func (repo * FollowRequestRepository) FindAllPendingFollowRequestsForUser(userId uuid.UUID) []model.FollowRequest{
	var allRequests = repo.FindAllFollowRequestsForFollower(userId)
	var pendingRequests []model.FollowRequest

	for i := 0; i < len(allRequests); i++{
		if allRequests[i].FollowRequestStatus == model.PENDING{
			pendingRequests = append(pendingRequests, allRequests[i])
		}
	}

	return pendingRequests
}



func (repo * FollowRequestRepository) FindFollowRequest(classicUserId uuid.UUID, followerUserId uuid.UUID ) *model.FollowRequest{
	request := &model.FollowRequest{}

	if repo.Database.First(&request, "classic_user_id = ? and follower_user_id = ?", classicUserId, followerUserId).RowsAffected == 0{
		return nil
	}
	return request

}

func (repo *FollowRequestRepository) UpdateFollowRequestPending(followRequestId uuid.UUID) error {
	result := repo.Database.Model(&model.FollowRequest{}).Where("id = ?", followRequestId).Update("follow_request_status", model.PENDING)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *FollowRequestRepository) UpdateFollowRequestAccepted(followRequestId uuid.UUID) error {
	result := repo.Database.Model(&model.FollowRequest{}).Where("id = ?", followRequestId).Update("follow_request_status", model.ACCEPTED)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *FollowRequestRepository) UpdateFollowRequestRejected(followRequestId uuid.UUID) error {
	result := repo.Database.Model(&model.FollowRequest{}).Where("id = ?", followRequestId).Update("follow_request_status", model.REJECT)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}
