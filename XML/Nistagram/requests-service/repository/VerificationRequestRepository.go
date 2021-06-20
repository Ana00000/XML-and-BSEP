package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"gorm.io/gorm"
)

type VerificationRequestRepository struct {
	Database *gorm.DB
}

func (repo *VerificationRequestRepository) CreateVerificationRequest(verificationRequest *model.VerificationRequest) error {
	result := repo.Database.Create(verificationRequest)
	fmt.Print(result)
	return nil
}

func (repo *VerificationRequestRepository) FindById(id uuid.UUID) *model.VerificationRequest {
	request := &model.VerificationRequest{}
	if repo.Database.First(&request, "id = ?", id).RowsAffected == 0 {
		return nil
	}
	return request
}

func (repo *VerificationRequestRepository) FindAllPendingVerificationRequests() []model.VerificationRequest {
	var requests []model.VerificationRequest

	repo.Database.Select("*").Where("verification_request_status = ?", 0).Find(&requests)
	return requests
}


func (repo *VerificationRequestRepository) UpdateVerificationRequestAccepted(verificationRequestId uuid.UUID) error {
	result := repo.Database.Model(&model.VerificationRequest{}).Where("id = ?", verificationRequestId).Update("verification_request_status", model.ACCEPTED)
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo *VerificationRequestRepository) UpdateVerificationRequestRejected(verificationRequestId uuid.UUID) error {
	result := repo.Database.Model(&model.VerificationRequest{}).Where("id = ?", verificationRequestId).Update("verification_request_status", model.REJECT)
	fmt.Println(result.RowsAffected)
	return nil
}