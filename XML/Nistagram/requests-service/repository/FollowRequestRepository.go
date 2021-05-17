package repository

import (
	"fmt"
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
