package handler

import (
	contentModel "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	contentService "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	locationModel "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	locationService "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	tagModel "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	tagsService "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type StoryAlbumHandler struct {
	Service * service.StoryAlbumService
	StoryService * service.StoryService
	ClassicUserService * userService.ClassicUserService
	ClassicUserFollowingsService * userService.ClassicUserFollowingsService
	ProfileSettings *settingsService.ProfileSettingsService
	StoryAlbumContentService *contentService.StoryAlbumContentService
	LocationService *locationService.LocationService
	StoryAlbumTagStoryAlbumsService *tagsService.StoryAlbumTagStoryAlbumsService
	TagService *tagsService.TagService
}

func (handler *StoryAlbumHandler) CreateStoryAlbum(w http.ResponseWriter, r *http.Request) {
	var storyAlbumDTO dto.StoryAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyAlbumType := model.CLOSE_FRIENDS
	switch storyAlbumDTO.Type {
	case "ALL_FRIENDS":
		storyAlbumType = model.ALL_FRIENDS
	case "PUBLIC":
		storyAlbumType = model.PUBLIC
	}

	id := uuid.New()
	storyAlbum := model.StoryAlbum{
		Story : model.Story{
			ID:          	id,
			CreationDate: 	time.Now(),
			Description:    storyAlbumDTO.Description,
			UserId:      	storyAlbumDTO.UserId,
			LocationId:     storyAlbumDTO.LocationId,
			IsDeleted:      false,
			Type:      		storyAlbumType,
		},
	}

	err = handler.Service.CreateStoryAlbum(&storyAlbum)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.StoryService.CreateStory(&storyAlbum.Story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	storyAlbumIDJson, _ := json.Marshal(storyAlbum.ID)
	w.Write(storyAlbumIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryAlbumHandler) FindAllAlbumStoriesForLoggedUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var albumStories = handler.Service.FindAllAlbumStoriesForUser(uuid.MustParse(id))
	var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbums(albumStories)
	var locations = handler.LocationService.FindAllLocationsForStoryAlbums(albumStories)
	var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(albumStories)
	var albumStoriesDTOS = handler.CreateStoryAlbumsDTOList(albumStories,contents,locations,tags)

	albumStoriesJson, _ := json.Marshal(albumStoriesDTOS)
	w.Write(albumStoriesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryAlbumHandler) CreateStoryAlbumsDTOList(albums []model.StoryAlbum, contents []contentModel.StoryAlbumContent, locations []locationModel.Location, tags []tagModel.StoryAlbumTagStoryAlbums) []dto.SelectedStoryAlbumDTO {
	var listOfStoryAlbumsDTOs []dto.SelectedStoryAlbumDTO

	for i := 0; i < len(albums); i++ {
		var storyAlbumDTO dto.SelectedStoryAlbumDTO
		storyAlbumDTO.StoryAlbumId = albums[i].ID
		storyAlbumDTO.Description = albums[i].Description
		storyAlbumDTO.CreationDate = albums[i].CreationDate
		storyAlbumDTO.UserId = albums[i].UserId

		for j := 0; j < len(contents); j++ {
			if contents[j].StoryAlbumId == albums[i].ID {
				storyAlbumDTO.Path = append(storyAlbumDTO.Path, contents[j].Path)

				if contents[j].Type == contentModel.VIDEO {
					storyAlbumDTO.Type = append(storyAlbumDTO.Type, "VIDEO")
				} else if contents[j].Type == contentModel.PICTURE {
					storyAlbumDTO.Type = append(storyAlbumDTO.Type, "PICTURE")
				}
			}
		}

		for k := 0; k < len(locations); k++ {
			if locations[k].ID == albums[i].LocationId {
				storyAlbumDTO.LocationId = locations[k].ID
				storyAlbumDTO.City = locations[k].City
				storyAlbumDTO.Country = locations[k].Country
				storyAlbumDTO.StreetName = locations[k].StreetName
				storyAlbumDTO.StreetNumber = locations[k].StreetNumber
			}
		}

		var listOfTags []string
		for p := 0; p < len(tags); p++ {
			if tags[p].StoryAlbumId == albums[i].ID {
				listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))
			}
		}

		storyAlbumDTO.Tags = listOfTags
		listOfStoryAlbumsDTOs = append(listOfStoryAlbumsDTOs, storyAlbumDTO)
	}

	return listOfStoryAlbumsDTOs
}

func (handler *StoryAlbumHandler) FindSelectedStoryAlbumByIdForLoggedUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id") //story album id
	logId := r.URL.Query().Get("logId") //loged user id

	var storyAlbum = handler.Service.FindByID(uuid.MustParse(id))
	if storyAlbum == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if storyAlbum.IsDeleted == true{
		fmt.Println("Deleted story album")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if storyAlbum.UserId != uuid.MustParse(logId){
		fmt.Println("Story album doesnt belong to user")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbum(storyAlbum)
	var locations = handler.LocationService.FindAllLocationsForStoryAlbum(storyAlbum)
	var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbum(storyAlbum)

	var storyAlbumDTO = handler.CreateStoryAlbumDTO(storyAlbum,contents,locations,tags)

	storyAlbumJson, _ := json.Marshal(storyAlbumDTO)
	w.Write(storyAlbumJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *StoryAlbumHandler) CreateStoryAlbumDTO(album *model.StoryAlbum, contents []contentModel.StoryAlbumContent, locations []locationModel.Location, tags []tagModel.StoryAlbumTagStoryAlbums) dto.SelectedStoryAlbumDTO {
	var storyAlbumDTO dto.SelectedStoryAlbumDTO

	storyAlbumDTO.StoryAlbumId = album.ID
	storyAlbumDTO.Description = album.Description
	storyAlbumDTO.CreationDate = album.CreationDate


	for j := 0; j < len(contents); j++ {
		if contents[j].StoryAlbumId == album.ID {
			storyAlbumDTO.Path = append(storyAlbumDTO.Path, contents[j].Path)

			if contents[j].Type == contentModel.VIDEO {
				storyAlbumDTO.Type = append(storyAlbumDTO.Type, "VIDEO")
			} else if contents[j].Type == contentModel.PICTURE {
				storyAlbumDTO.Type = append(storyAlbumDTO.Type, "PICTURE")
			}
		}
	}

	for k := 0; k < len(locations); k++ {
		if locations[k].ID == album.LocationId {
			storyAlbumDTO.LocationId = locations[k].ID
			storyAlbumDTO.City = locations[k].City
			storyAlbumDTO.Country = locations[k].Country
			storyAlbumDTO.StreetName = locations[k].StreetName
			storyAlbumDTO.StreetNumber = locations[k].StreetNumber
		}
	}

	var listOfTags []string
	for p := 0; p < len(tags); p++ {
		if tags[p].StoryAlbumId == album.ID {
			listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))
		}
	}

	storyAlbumDTO.Tags = listOfTags
	return storyAlbumDTO
}

func (handler *StoryAlbumHandler) FindAllPublicAlbumStoriesRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidStoryAlbums = handler.Service.FindAllPublicAlbumStoriesNotRegisteredUser(allPublicUsers)
	var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbums(publicValidStoryAlbums)
	var locations = handler.LocationService.FindAllLocationsForStoryAlbums(publicValidStoryAlbums)
	var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(publicValidStoryAlbums)
	var storyAlbumsDTOS = handler.CreateStoryAlbumsDTOList(publicValidStoryAlbums,contents,locations,tags)

	storyAlbumsJson, _ := json.Marshal(storyAlbumsDTOS)
	w.Write(storyAlbumsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}