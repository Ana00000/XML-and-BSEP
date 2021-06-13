package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	"net/http"
	"os"
	_ "strconv"
	"strings"
	"time"
)

type LocationHandler struct {
	Service * service.LocationService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func Request(url string, token string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	tokenString := "Bearer "+token
	req.Header.Set("Authorization", tokenString)
	resp, err := http.DefaultClient.Do(req)
	return resp
}

func (handler *LocationHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "CRLON001",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-location-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "CRLON001",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	var locationDTO dto.LocationDTO
	err := json.NewDecoder(r.Body).Decode(&locationDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "CRLON001",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var foundLocation = handler.Service.FindByLocationDTO(locationDTO)
	if foundLocation != nil {
		fmt.Println(foundLocation.StreetName)
		locationIDJson, _ := json.Marshal(foundLocation.ID)
		w.Write(locationIDJson)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "CRLON001",
			"timestamp":   time.Now().String(),
		}).Error("Location already exists!")
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
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "CRLON001",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating location!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	locationIDJson, _ := json.Marshal(location.ID)
	w.Write(locationIDJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "CRLON001",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created location!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var location = handler.Service.FindByID(uuid.MustParse(id))
	if  location == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIBYIDI111",
			"timestamp":   time.Now().String(),
		}).Error("Location not found!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	locationJson, _ := json.Marshal(location)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIBYIDI111",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found location!")
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

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FILOIDBYLOSTR670",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found location!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStories(w http.ResponseWriter, r *http.Request) {
	var singleStoriesDTO []dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoriesDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIALLOFOSTH070",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoriesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStories(singleStoriesDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOSTH070",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIALLOFOSTD780",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoryDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStory(&singleStoryDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOSTD780",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPosts(w http.ResponseWriter, r *http.Request) {
	var singlePostsDTO []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "LocationHandler",
		"action":   "FIALLOFOPOP066",
		"timestamp":   time.Now().String(),
	}).Error("Wrong cast json to SinglePostsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPosts(singlePostsDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOPOP066",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIALLOFOPOG429",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SinglePostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPost(&singlePostDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOPOG429",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStoryAlbums(w http.ResponseWriter, r *http.Request) {
	var storyAlbumsFullDTO []dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumsFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIALLOFOSTALK777",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryAlbumsFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStoryAlbums(storyAlbumsFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOSTALK777",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForStoryAlbum(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIALLOFOSTALJ927",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForStoryAlbum(&storyAlbumFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOSTALJ927",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPostAlbums(w http.ResponseWriter, r *http.Request) {
	var postAlbumsFullDTO []dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumsFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIALLOFOPOALP969",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumsFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPostAlbums(postAlbumsFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOPOALP969",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindAllLocationsForPostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationHandler",
			"action":   "FIALLOFOPOALA073",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var locations = handler.Service.FindAllLocationsForPostAlbum(&postAlbumFullDTO)

	locationJson, _ := json.Marshal(locations)
	w.Write(locationJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "LocationHandler",
		"action":   "FIALLOFOPOALA073",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found locations!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
