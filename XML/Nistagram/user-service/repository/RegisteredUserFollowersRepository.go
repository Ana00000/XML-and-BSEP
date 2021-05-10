package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
"fmt"
"gorm.io/gorm"
)

type RegisteredUserFollowersRepository struct {
	Database * gorm.DB
}

func (repo * RegisteredUserFollowersRepository) CreateRegisteredUserFollowers(registeredUserFollowers *model.RegisteredUserFollowers) error {
	result := repo.Database.Create(registeredUserFollowers)
	fmt.Print(result)
	return nil
}
