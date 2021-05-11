package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type RecoveryPasswordTokenRepository struct {
	Database * gorm.DB
}

func (repo * RecoveryPasswordTokenRepository) CreateRecoveryPasswordToken(recoveryPasswordToken *model.RecoveryPasswordToken) error {
	result := repo.Database.Create(recoveryPasswordToken)
	fmt.Print(result)
	return nil
}

func (repo *RecoveryPasswordTokenRepository) FindByToken(token uuid.UUID) *model.RecoveryPasswordToken {
	recoveryPasswordToken := &model.RecoveryPasswordToken{}
	repo.Database.First(&recoveryPasswordToken, "recovery_password_token = ?", token)
	return recoveryPasswordToken
}

func (repo *RecoveryPasswordTokenRepository) UpdateRecoveryPasswordTokenValidity(token uuid.UUID, isValid bool) error {
	result := repo.Database.Model(&model.RecoveryPasswordToken{}).Where("recovery_password_token = ?", token).Update("is_valid", isValid)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}