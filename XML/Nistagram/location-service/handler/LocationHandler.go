package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	"net/http"
	_ "strconv"
)

type LocationHandler struct {
	Service *service.LocationService
}

func (handler *LocationHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var locationDTO dto.LocationDTO
	err := json.NewDecoder(r.Body).Decode(&locationDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var foundLocation = handler.Service.FindByLocationDTO(locationDTO)
	if foundLocation != nil {
		fmt.Println(foundLocation.StreetName)
		locationIDJson, _ := json.Marshal(foundLocation.ID)
		w.Write(locationIDJson)
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	fmt.Println("Nije pronasao lokaciju")
	location := model.Location{
		ID:           uuid.UUID{},
		Longitude:    locationDTO.Longitude,
		Latitude:     locationDTO.Latitude,
		Country:      locationDTO.Country,
		City:         locationDTO.City,
		StreetName:   locationDTO.StreetName,
		StreetNumber: locationDTO.StreetNumber,
	}

	err = handler.Service.CreateLocation(&location)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	locationIDJson, _ := json.Marshal(location.ID)
	w.Write(locationIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var location = handler.Service.FindByID(uuid.MustParse(id))
	if location == nil {
		fmt.Println("Location not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	locationJson, _ := json.Marshal(location)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

type ReturnValueId struct {
	ID uuid.UUID `json:"id"`
}

func (handler *LocationHandler) FindLocationIdByLocationString(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	locationString := vars["locationString"]

	var location = handler.Service.FindLocationIdByLocationString(locationString)

	locationId := ReturnValueId{
		ID: location.ID,
	}
	locationJson, _ := json.Marshal(locationId)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStories(w http.ResponseWriter, r *http.Request) {
	var singleStoriesDTO []dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoriesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStories(singleStoriesDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStory(&singleStoryDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPosts(w http.ResponseWriter, r *http.Request) {
	var singlePostsDTO []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPosts(singlePostsDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPost(&singlePostDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStoryAlbums(w http.ResponseWriter, r *http.Request) {
	var storyAlbumsFullDTO []dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumsFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStoryAlbums(storyAlbumsFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStoryAlbum(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStoryAlbum(&storyAlbumFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPostAlbums(w http.ResponseWriter, r *http.Request) {
	var postAlbumsFullDTO []dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumsFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPostAlbums(postAlbumsFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPostAlbum(&postAlbumFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
