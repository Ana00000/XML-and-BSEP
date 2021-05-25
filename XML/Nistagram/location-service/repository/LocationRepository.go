package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"gorm.io/gorm"
	"strings"
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

func (repo *LocationRepository) FindByLocationDTO(locationDTO dto.LocationDTO) *model.Location {
	location := &model.Location{}
	if repo.Database.First(&location, "street_name=? and street_number=? and city=? and country=? and longitude=? and latitude=?", locationDTO.StreetName, locationDTO.StreetNumber, locationDTO.City, locationDTO.Country,locationDTO.Longitude,locationDTO.Latitude).RowsAffected == 0 {
		return nil
	}
	fmt.Println(location)
	return location
}


func (repo *LocationRepository) FindAllLocationsForPosts(allPosts []dto.SinglePostDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for i:=0;i<len(allPosts);i++{
		for j:=0; j<len(allLocations);j++{
			if allPosts[i].LocationId == allLocations[j].ID && !ExsistInList(allLocations[j],locations){
				locations = append(locations, allLocations[j])
			}
		}

	}
	return locations
}

func ExsistInList(location model.Location, allLocations []model.Location) bool{
	for i := 0; i < len(allLocations); i++ {
		if allLocations[i].ID == location.ID{
			return true
		}
	}
	return false
}


func (repo *LocationRepository) FindAllLocationsForPost(post *dto.SinglePostDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for j:=0; j<len(allLocations);j++{
			if post.LocationId == allLocations[j].ID && ExsistInList(allLocations[j],locations){
				locations = append(locations, allLocations[j])
			}
		}

	return locations
}

func (repo *LocationRepository) FindAllLocationsForStories(allStories []dto.SingleStoryDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for i:=0;i<len(allStories);i++{
		for j:=0; j<len(allLocations);j++{
			if allStories[i].LocationId == allLocations[j].ID{
				locations = append(locations, allLocations[j])
			}
		}

	}
	return locations
}

func (repo *LocationRepository) FindAllLocationsForStory(story *dto.SingleStoryDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for j:=0; j<len(allLocations);j++{
		if story.LocationId == allLocations[j].ID{
			locations = append(locations, allLocations[j])
		}
	}


	return locations
}

func (repo *LocationRepository) FindLocationIdByLocationString(locationString string) model.Location {
	var locationStringParts = strings.Split(locationString, ",")

	var country = locationStringParts[0]
	var city = locationStringParts[1]
	var streetName = locationStringParts[2]
	var streetNumber = locationStringParts[3]


	var location   = model.Location{}

	var allLocations = repo.FindAll()

	for i:=0; i<len(allLocations);i++{
		if allLocations[i].Country == country && allLocations[i].City == city && allLocations[i].StreetName == streetName && allLocations[i].StreetNumber == streetNumber{
			location =  allLocations[i]
		}
	}

	return location

}

func (repo *LocationRepository) FindAllLocationsForPostAlbums(albums []dto.PostAlbumFullDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for i:=0;i<len(albums);i++{
		for j:=0; j<len(allLocations);j++{
			if albums[i].LocationId == allLocations[j].ID{
				locations = append(locations, allLocations[j])
			}
		}

	}
	return locations
}

func (repo *LocationRepository) FindAllLocationsForPostAlbum(album *dto.PostAlbumFullDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for j:=0; j<len(allLocations);j++{
		if album.LocationId == allLocations[j].ID{
			locations = append(locations, allLocations[j])
		}
	}

	return locations
}

func (repo *LocationRepository) FindAllLocationsForStoryAlbums(albums []dto.StoryAlbumFullDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for i:=0;i<len(albums);i++{
		for j:=0; j<len(allLocations);j++{
			if albums[i].LocationId == allLocations[j].ID{
				locations = append(locations, allLocations[j])
			}
		}

	}
	return locations
}

func (repo *LocationRepository) FindAllLocationsForStoryAlbum(album *dto.StoryAlbumFullDTO) []model.Location {
	var locations []model.Location
	var allLocations = repo.FindAll()

	for j:=0; j<len(allLocations);j++{
		if album.LocationId == allLocations[j].ID{
			locations = append(locations, allLocations[j])
		}
	}

	return locations
}