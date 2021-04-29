package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type VerificationRequestRepository struct {
	Database * gorm.DB
}

func (repo * VerificationRequestRepository) CreateVerificationRequest(verificationRequest *model.VerificationRequest) error {
	result := repo.Database.Create(verificationRequest)
	fmt.Print(result)
	return nil
}