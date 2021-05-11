package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type AdminRepository struct {
	Database * gorm.DB
}

func (repo * AdminRepository) CreateAdmin(admin *model.Admin) error {
	result := repo.Database.Create(admin)
	fmt.Print(result)
	return nil
}

func (repo *AdminRepository) UpdateAdminPassword(userId uuid.UUID, password string) error {
	result := repo.Database.Model(&model.Admin{}).Where("id = ?", userId).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}