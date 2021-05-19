package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"gorm.io/gorm"
)

type LocationRepository struct {
	Database * gorm.DB
}

func (repo * LocationRepository) CreateLocation(location *model.Location) error {
	result := repo.Database.Create(location)
	fmt.Print(result)
	return nil
}

func (repo *LocationRepository) FindByID(ID uuid.UUID) *model.Location {
	location := &model.Location{}
	repo.Database.First(&location, "id = ?", ID)
	return location
}

func (repo *LocationRepository) FindAll() []model.Location {
	var locations []model.Location
	repo.Database.Select("*").Find(&locations)
	return locations
}


func (repo *LocationRepository) FindAllLocationsForPosts(allPosts []postsModel.Post) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for i:=0;i<len(allPosts);i++{
		for j:=0; j<len(allLocations);j++{
			if allPosts[i].LocationId == allLocations[j].ID{
				locations = append(locations, allLocations[j])
			}
		}

	}
	return locations
}