package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"fmt"
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