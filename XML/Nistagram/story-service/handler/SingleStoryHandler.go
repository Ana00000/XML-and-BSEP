package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type SingleStoryHandler struct {
	SingleStoryService * service.SingleStoryService
	StoryService * service.StoryService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//CRSINGLSTRY9023
func (handler *SingleStoryHandler) CreateSingleStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "CRSINGLSTRY9023",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoryDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	singleStoryType := model.CLOSE_FRIENDS
	switch singleStoryDTO.Type {
	case "ALL_FRIENDS":
		singleStoryType = model.ALL_FRIENDS
	case "PUBLIC":
		singleStoryType = model.PUBLIC
	}

	id := uuid.New()
	singleStory := model.SingleStory{
		Story: model.Story{
			ID:           id,
			CreationDate: time.Now(),
			Description:  singleStoryDTO.Description,
			UserId:       singleStoryDTO.UserId,
			LocationId:   singleStoryDTO.LocationId,
			IsDeleted:    false,
			IsExpired: 	  false,
			Type:         singleStoryType,
		},
	}

	err = handler.SingleStoryService.CreateSingleStory(&singleStory)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "CRSINGLSTRY9023",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating single story!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.StoryService.CreateStory(&singleStory.Story)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "CRSINGLSTRY9023",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating basic story!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	singleStoryIDJson, _ := json.Marshal(singleStory.ID)
	w.Write(singleStoryIDJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "CRSINGLSTRY9023",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created single story!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//// tab PUBLIC STORIES kada neregistroviani korisnik otvori sve PUBLIC, NOT EXPIRED I OD PUBLIC USERA
//FIDALPUBSTORISNOTREGUS9329
func (handler *SingleStoryHandler) FindAllPublicStoriesNotRegisteredUser(w http.ResponseWriter, r *http.Request) {
	//var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_valid_users/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all valid users or wrong cast response body to list ClassicUserDTO!!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	//var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all public users!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var allPublicUsers []dto.ClassicUserDTO
	if err := json.NewDecoder(resp.Body).Decode(&allPublicUsers); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ClassicUserDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var publicValidStories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllPublicStoriesNotRegisteredUser(allPublicUsers))

	//var contents = handler.StoryContentService.FindAllContentsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_stories/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoriesDTO, _ := json.Marshal(publicValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoriesDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding contents for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SingleStoryContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SingleStoryContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//var locations = handler.LocationService.FindAllLocationsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_stories/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding locations for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_stories/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding tags for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.StoryTagStoriesDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISNOTREGUS9329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryTagStoriesDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var storiesDTOS = handler.CreateStoriesDTOList(convertSingleStoriesDTOToListSingleStories(publicValidStories),contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "FIDALPUBSTORISNOTREGUS9329",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all public stories not registered user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

type UserValid struct {
	IsValid bool `json:"is_valid"`
}

// kada neregistrovani user otvori PUBLIC usera sa spiska i onda na njegovom profilu vidi PUBLIC i NOT EXPIRED STORIJE
//FIDALSTORISFORUSNOTREGUS8921
func (handler *SingleStoryHandler) FindAllStoriesForUserNotRegisteredUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	///check_if_user_valid/{userID}
	//var checkIfValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(id))
	var userValidity UserValid
	reqUrl := fmt.Sprintf("http://%s:%s/check_if_user_valid/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	err := getJson(reqUrl, &userValidity)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Failed checking if user valid or wrong cast json to UserValid!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	var checkIfValid=userValidity.IsValid
	if  checkIfValid == false {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't valid!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//fmt.Println("User IS valid")
	//var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	var profileSettings dto.ProfileSettingsDTO
	reqUrl = fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err = getJson(reqUrl, &profileSettings)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding profile settings by user id or wrong cast json to ProfileSettingsDTO!")
		//fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if profileSettings.UserVisibility == "PRIVATE_VISIBILITY"{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("User is private!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}


	var stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllStoriesForUserNotReg(uuid.MustParse(id)))
	//var contents = handler.StoryContentService.FindAllContentsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_stories/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoriesDTO, _ := json.Marshal(stories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoriesDTO))
	resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SingleStoryContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SingleStoryContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var locations = handler.LocationService.FindAllLocationsForStories(stories)
	//var locations = handler.LocationService.FindAllLocationsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_stories/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(stories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
	//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_stories/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(stories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.StoryTagStoriesDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSNOTREGUS8921",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryTagStoriesDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var storiesDTOS = handler.CreateStoriesDTOList(convertSingleStoriesDTOToListSingleStories(stories),contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "FIDALSTORISFORUSNOTREGUS8921",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all stories for user not registered user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")


}

//FIDALPUBSTORISREGUS9823
func (handler *SingleStoryHandler) FindAllPublicStoriesRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"),id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all classic users without logged in or wrong cast json to list ClassicUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	//var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all public users!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var allPublicUsers []dto.ClassicUserDTO
	if err := json.NewDecoder(resp.Body).Decode(&allPublicUsers); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ClassicUserDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var publicValidStories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllPublicStoriesNotRegisteredUser(allPublicUsers))
	//var contents = handler.StoryContentService.FindAllContentsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_stories/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoriesDTO, _ := json.Marshal(publicValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoriesDTO))
	resp,err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SingleStoryContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SingleStoryContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var locations = handler.LocationService.FindAllLocationsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_stories/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(publicValidStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_stories/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.StoryTagStoriesDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALPUBSTORISREGUS9823",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryTagStoriesDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var storiesDTOS = handler.CreateStoriesDTOList(convertSingleStoriesDTOToListSingleStories(publicValidStories),contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "FIDALPUBSTORISREGUS9823",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all public stories registered user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

type ReturnValueBool struct {
	ReturnValue bool `json:"return_value"`
}

// metoda koja se poziva kada registrovani user udje na profil nekog usera
//FIDALSTORISFORUSREGUS9322
func (handler *SingleStoryHandler) FindAllStoriesForUserRegisteredUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")
	var stories []dto.SingleStoryFullDTO

	var userValidity UserValid
	reqUrl := fmt.Sprintf("http://%s:%s/check_if_user_valid/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	err := getJson(reqUrl, &userValidity)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSREGUS9322",
			"timestamp":   time.Now().String(),
		}).Error("Failed checking if user valid or wrong cast json to UserValid!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	var checkIfValid=userValidity.IsValid
	if  checkIfValid == false {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSREGUS9322",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't valid!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	//var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	var profileSettings dto.ProfileSettingsDTO
	reqUrl = fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err = getJson(reqUrl, &profileSettings)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSREGUS9322",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding profile settings by user id or wrong cast json to ProfileSettingsDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if profileSettings.UserVisibility == "PRIVATE_VISIBILITY"{

		//PRIVATE USER

		//var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		var returnValueFollowing ReturnValueBool
		reqUrl = fmt.Sprintf("http://%s:%s/check_if_following_post_story/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id, logId)
		err = getJson(reqUrl, &returnValueFollowing)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Failed checking if following post/story or wrong cast json to ReturnValueBool!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		checkIfFollowing := returnValueFollowing.ReturnValue
		if checkIfFollowing == true{

			// PRATI GA
			//var checkIfCloseFriend = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(uuid.MustParse(id), uuid.MustParse(logId))
			var returnValueCloseFriend ReturnValueBool
			reqUrl = fmt.Sprintf("http://%s:%s/check_if_close_friend/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id, logId)
			err = getJson(reqUrl, &returnValueCloseFriend)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Failed checking if close friend or wrong cast json to ReturnValueBool!")
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			checkIfCloseFriend := returnValueCloseFriend.ReturnValue

			if checkIfCloseFriend == true{

				// NALAZI SE U CLOSE FRIENDS

				stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllStoriesForUserCloseFriend(uuid.MustParse(id)))


			} else{
				// NE NALAZI SE U CLOSE FRIENDS ALI GA PRATI - PUBLIC I ALL FRIENDS
				stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllStoriesForUserPublicAllFriends(uuid.MustParse(id)))
			}




			//var contents = handler.StoryContentService.FindAllContentsForStories(stories)
			reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_stories/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
			jsonValidStoriesDTO, _ := json.Marshal(stories)
			//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			//fmt.Println(string(jsonValidStoriesDTO))
			resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
			if err != nil || resp.StatusCode == 400 {
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Failed finding all contents for stories!")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var contents []dto.SingleStoryContentDTO
			if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Wrong cast json to list SingleStoryContentDTO!")
				w.WriteHeader(http.StatusConflict) //400
				return
			}


			//var locations = handler.LocationService.FindAllLocationsForStories(stories)
			reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_stories/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
			jsonLocationsDTO, _ := json.Marshal(stories)
			//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			//fmt.Println(string(jsonLocationsDTO))
			resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
			if err != nil || resp.StatusCode == 400 {
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Failed finding all locations for stories!")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var locations []dto.LocationDTO
			if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Wrong cast json to list LocationDTO!")
				w.WriteHeader(http.StatusConflict) //400
				return
			}


			//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
			reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_stories/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
			jsonTagsDTO, _ := json.Marshal(stories)
			//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			//fmt.Println(string(jsonTagsDTO))
			resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
			if err != nil || resp.StatusCode == 400 {
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Failed finding all tags for stories!")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var tags []dto.StoryTagStoriesDTO
			if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Wrong cast json to list StoryTagStoriesDTO!")
				w.WriteHeader(http.StatusConflict) //400
				return
			}
			var storiesDTOS = handler.CreateStoriesDTOList(convertSingleStoriesDTOToListSingleStories(stories),contents,locations,tags)

			storiesJson, _ := json.Marshal(storiesDTOS)
			w.Write(storiesJson)
			w.WriteHeader(http.StatusOK)

			handler.LogInfo.WithFields(logrus.Fields{
				"status": "success",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Info("Successfully founded all stories for user registered user!")

			w.Header().Set("Content-Type", "application/json")
			return

		}else{
			// PRIVATE USER I NE PRATI GA = NE MOZE NISTA DA VIDI
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("User doesn't following private user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}else{
		//PUBLIC USER
		//var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		var returnValueFollowing ReturnValueBool
		reqUrl = fmt.Sprintf("http://%s:%s/check_if_following_post_story/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id, logId)
		err = getJson(reqUrl, &returnValueFollowing)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Failed checking if following post/story or wrong cast json to ReturnValueBool!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		checkIfFollowing := returnValueFollowing.ReturnValue

		if checkIfFollowing == true{

			// PRATI GA
			//var checkIfCloseFriend = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(uuid.MustParse(id), uuid.MustParse(logId))
			var returnValueCloseFriend ReturnValueBool
			reqUrl = fmt.Sprintf("http://%s:%s/check_if_close_friend/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id, logId)
			err = getJson(reqUrl, &returnValueCloseFriend)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALSTORISFORUSREGUS9322",
					"timestamp":   time.Now().String(),
				}).Error("Failed checking if close friend or wrong cast json to ReturnValueBool!")
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			checkIfCloseFriend := returnValueCloseFriend.ReturnValue

			if checkIfCloseFriend == true{

				// NALAZI SE U CLOSE FRIENDS
				stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllStoriesForUserCloseFriend(uuid.MustParse(id)))


			} else{
				// NE NALAZI SE U CLOSE FRIENDS ALI GA PRATI - PUBLIC I ALL FRIENDS
				stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllStoriesForUserPublicAllFriends(uuid.MustParse(id)))
			}

		}else{
			//NE PRATI GA ALI POSTO JE PUBLIC SME DA VIDI PUBLIC STORIJE
			stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllStoriesForUserPublic(uuid.MustParse(id)))
		}
		//var contents = handler.StoryContentService.FindAllContentsForStories(stories)
		reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_stories/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
		jsonValidStoriesDTO, _ := json.Marshal(stories)
		//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		//fmt.Println(string(jsonValidStoriesDTO))
		resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
		if err != nil || resp.StatusCode == 400 {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Failed finding all contents for stories!")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var contents []dto.SingleStoryContentDTO
		if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Wrong cast json to list SingleStoryContentDTO!")
			w.WriteHeader(http.StatusConflict) //400
			return
		}


		//var locations = handler.LocationService.FindAllLocationsForStories(stories)
		reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_stories/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
		jsonLocationsDTO, _ := json.Marshal(stories)
		//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		//fmt.Println(string(jsonLocationsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
		if err != nil || resp.StatusCode == 400 {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Failed finding all locations for stories!")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var locations []dto.LocationDTO
		if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Wrong cast json to list LocationDTO!")
			w.WriteHeader(http.StatusConflict) //400
			return
		}


		//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
		reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_stories/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
		jsonTagsDTO, _ := json.Marshal(stories)
		//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		//fmt.Println(string(jsonTagsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
		if err != nil || resp.StatusCode == 400 {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Failed finding all tags for stories!")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var tags []dto.StoryTagStoriesDTO
		if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "SingleStoryHandler",
				"action":   "FIDALSTORISFORUSREGUS9322",
				"timestamp":   time.Now().String(),
			}).Error("Wrong cast json to list StoryTagStoriesDTO!")
			w.WriteHeader(http.StatusConflict) //400
			return
		}
		var storiesDTOS = handler.CreateStoriesDTOList(convertSingleStoriesDTOToListSingleStories(stories),contents,locations,tags)

		storiesJson, _ := json.Marshal(storiesDTOS)
		w.Write(storiesJson)
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORUSREGUS9322",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded all stories for user registered user!")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")


	}
}


// returns all VALID stories from FOLLOWING users (FOR HOMEPAGE)
//FIDALFOLLINGSTORIS8329
func (handler *SingleStoryHandler) FindAllFollowingStories(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"),id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ClassicUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	//var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_valid_followings_for_user/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all valid followings for user!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var followings []dto.ClassicUserFollowingsDTO
	if err := json.NewDecoder(resp.Body).Decode(&followings); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ClassicUserFollowingsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var allValidStories []dto.SingleStoryFullDTO
	var stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllFollowingStories(followings))


	for i:=0; i<len(stories);i++{
		if stories[i].Type == "PUBLIC" || stories[i].Type == "ALL_FRIENDS"{

			allValidStories = append(allValidStories, stories[i])

		}else if stories[i].Type == "CLOSE_FRIENDS"{

			//var checkIfCloseFriend = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(stories[i].UserId, uuid.MustParse(id))
			var returnValueCloseFriend ReturnValueBool
			reqUrl = fmt.Sprintf("http://%s:%s/check_if_close_friend/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), stories[i].UserId, id)
			err = getJson(reqUrl, &returnValueCloseFriend)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "FIDALFOLLINGSTORIS8329",
					"timestamp":   time.Now().String(),
				}).Error("Failed checking if close friend or wrong cast json to ReturnValueBool!")
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			checkIfCloseFriend := returnValueCloseFriend.ReturnValue
			if checkIfCloseFriend == true{

				allValidStories = append(allValidStories, stories[i])
			}
		}
	}

	//var contents = handler.StoryContentService.FindAllContentsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_stories/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoriesDTO, _ := json.Marshal(allValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoriesDTO))
	resp,err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SingleStoryContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SingleStoryContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var locations = handler.LocationService.FindAllLocationsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_stories/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(allValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_stories/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(allValidStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.StoryTagStoriesDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALFOLLINGSTORIS8329",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryTagStoriesDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var storiesDTOS = handler.CreateStoriesDTOList(convertSingleStoriesDTOToListSingleStories(allValidStories) ,contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "FIDALFOLLINGSTORIS8329",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all following stories!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

// FIND SELECTED STORY BY ID (ONLY IF NOT DELETED)!
//FIDSELECTSTRYBYIDFORREGUSRS9031
func (handler *SingleStoryHandler) FindSelectedStoryByIdForRegisteredUsers(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var story = handler.SingleStoryService.FindByID(uuid.MustParse(id))
	if story == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding single story by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if story.IsDeleted == true{

		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Founded single story is deleted!")
		w.WriteHeader(http.StatusExpectationFailed)
		return

	}

	if story.UserId != uuid.MustParse(logId){
		//POSTMAN CHECK
		//NIJE STORI OD ULOGOVANOG USERA
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Unavailable story to this user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return

	}

	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_story/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoriesDTO, _ := json.Marshal(convertSingleStoryToSingleStoryDTO(*story))
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoriesDTO))
	resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for story!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SingleStoryContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SingleStoryContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var locations = handler.LocationService.FindAllLocationsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_story/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(convertSingleStoryToSingleStoryDTO(*story))
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for story!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_story/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(convertSingleStoryToSingleStoryDTO(*story))
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for story!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.StoryTagStoriesDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryTagStoriesDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	/*
	var contents = handler.StoryContentService.FindAllContentsForStories(story)
	var locations = handler.LocationService.FindAllLocationsForStory(story)
	var tags = handler.StoryTagStoriesService.FindAllTagsForStory(story)
	 */
	var storyDTO = handler.CreateStoryDTO(story,contents,locations,tags)

	storyJson, _ := json.Marshal(storyDTO)
	w.Write(storyJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "FIDSELECTSTRYBYIDFORREGUSRS9031",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded selected story by ID for registered users!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")


}

type ReturnValueString struct {
	ReturnValue string `json:"return_value"`
}


// all stories (EXCEPT DELETED) for my current logged in user (expired and not expired, public, all_friend, close friends)
//FIDALSTORISFORLOGGUS0213
func (handler *SingleStoryHandler) FindAllStoriesForLoggedUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var stories = convertListSingleStoriesToSingleStoriesDTO(handler.SingleStoryService.FindAllStoriesForLoggedUser(uuid.MustParse(id)))
	//var contents = handler.StoryContentService.FindAllContentsForStories(stories)
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_stories/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoriesDTO, _ := json.Marshal(stories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoriesDTO))
	resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORLOGGUS0213",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SingleStoryContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORLOGGUS0213",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SingleStoryContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//var locations = handler.LocationService.FindAllLocationsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_stories/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(stories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORLOGGUS0213",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORLOGGUS0213",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_stories/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(stories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORLOGGUS0213",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for stories!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.StoryTagStoriesDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryHandler",
			"action":   "FIDALSTORISFORLOGGUS0213",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryTagStoriesDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var storiesDTOS = handler.CreateStoriesDTOList(convertSingleStoriesDTOToListSingleStories(stories),contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)
	w.WriteHeader(http.StatusOK)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "FIDALSTORISFORLOGGUS0213",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all stories for logged user!")
	w.Header().Set("Content-Type", "application/json")

}

//DTOS
//CRSTORISDTOLST0914
func (handler *SingleStoryHandler) CreateStoriesDTOList(stories []model.SingleStory, contents []dto.SingleStoryContentDTO, locations []dto.LocationDTO, tags []dto.StoryTagStoriesDTO) []dto.SelectedStoryDTO {
	var listOfStoriesDTOs []dto.SelectedStoryDTO

	for i := 0; i < len(stories); i++ {
		var storyDTO dto.SelectedStoryDTO
		storyDTO.StoryId = stories[i].ID
		storyDTO.Description = stories[i].Description
		storyDTO.CreationDate = stories[i].CreationDate
		storyDTO.UserId = stories[i].UserId



		for j := 0; j < len(contents); j++ {
			if contents[j].SingleStoryId == stories[i].ID {
				storyDTO.Path = contents[j].Path
				storyDTO.Type = contents[j].Type
			}
		}

		for k := 0; k < len(locations); k++ {
			if locations[k].ID == stories[i].LocationId {
				storyDTO.LocationId = locations[k].ID
				storyDTO.City = locations[k].City
				storyDTO.Country = locations[k].Country
				storyDTO.StreetName = locations[k].StreetName
				storyDTO.StreetNumber = locations[k].StreetNumber
			}
		}

		var listOfTags []string
		for p := 0; p < len(tags); p++ {
			if tags[p].StoryId == stories[i].ID {
				//{id}
				var  returnValueTagName ReturnValueString
				reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"),tags[p].TagId.String())
				err := getJson(reqUrl, &returnValueTagName)
				if err!=nil{
					handler.LogError.WithFields(logrus.Fields{
						"status": "failure",
						"location":   "SingleStoryHandler",
						"action":   "CRSTORISDTOLST0914",
						"timestamp":   time.Now().String(),
					}).Error("Wrong cast response body to list ReturnValueString!")
					return nil
				}
				listOfTags = append(listOfTags,returnValueTagName.ReturnValue)
			}
		}

		storyDTO.Tags = listOfTags

		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "SingleStoryHandler",
			"action":   "CRSTORISDTOLST0914",
			"timestamp":   time.Now().String(),
		}).Info("Successfully created StoryDTO objects!")

		listOfStoriesDTOs = append(listOfStoriesDTOs, storyDTO)

	}

	return listOfStoriesDTOs

}

//CRSTRYDTO0912
func (handler *SingleStoryHandler) CreateStoryDTO(story *model.SingleStory, contents []dto.SingleStoryContentDTO, locations []dto.LocationDTO, tags []dto.StoryTagStoriesDTO) dto.SelectedStoryDTO {


	var storyDTO dto.SelectedStoryDTO
	//fmt.Println("STORIES")
	storyDTO.StoryId = story.ID
	storyDTO.Description = story.Description
	storyDTO.CreationDate = story.CreationDate

	for j := 0; j < len(contents); j++ {
		if contents[j].SingleStoryId == story.ID {
			storyDTO.Path = contents[j].Path
			storyDTO.Type = contents[j].Type
		}
	}

	for k := 0; k < len(locations); k++ {
		if locations[k].ID == story.LocationId {
			storyDTO.LocationId = locations[k].ID
			storyDTO.City = locations[k].City
			storyDTO.Country = locations[k].Country
			storyDTO.StreetName = locations[k].StreetName
			storyDTO.StreetNumber = locations[k].StreetNumber
		}
	}

	var listOfTags []string
	for p := 0; p < len(tags); p++ {
		if tags[p].StoryId == story.ID {
			var  returnValueTagName ReturnValueString
			reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"),tags[p].TagId.String())
			err := getJson(reqUrl, &returnValueTagName)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "SingleStoryHandler",
					"action":   "CRSTRYDTO0912",
					"timestamp":   time.Now().String(),
				}).Error("Wrong cast response body to list ReturnValueString!")
				panic(err)
			}
			listOfTags = append(listOfTags, returnValueTagName.ReturnValue)

		}
	}

	storyDTO.Tags = listOfTags
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryHandler",
		"action":   "CRSTRYDTO0912",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created StoryDTO object!")

	return storyDTO

}

func convertListSingleStoriesToSingleStoriesDTO(singleStories []model.SingleStory) []dto.SingleStoryFullDTO{
	var singleStoriesDTO []dto.SingleStoryFullDTO
	for i := 0; i < len(singleStories); i++ {
		singleStoriesDTO=append(singleStoriesDTO,convertSingleStoryToSingleStoryDTO(singleStories[i]))
	}
	return singleStoriesDTO
}

func convertSingleStoryToSingleStoryDTO(singleStory model.SingleStory) dto.SingleStoryFullDTO{
	storyType := ""
	if singleStory.Type==model.CLOSE_FRIENDS {
		storyType="CLOSE_FRIENDS"
	} else if singleStory.Type==model.PUBLIC{
		storyType="PUBLIC"
	} else if singleStory.Type==model.ALL_FRIENDS{
		storyType="ALL_FRIENDS"
	}
	layout := "2006-01-02T15:04:05.000Z"
	var singleStoryDTO= dto.SingleStoryFullDTO{
		ID:           singleStory.ID,
		Description:  singleStory.Description,
		CreationDate: singleStory.CreationDate.Format(layout),
		UserId:       singleStory.UserId,
		LocationId:   singleStory.LocationId,
		IsDeleted:    singleStory.IsDeleted,
		IsExpired:    singleStory.IsExpired,
		Type:         storyType,
	}
	return singleStoryDTO
}

func convertSingleStoriesDTOToListSingleStories(singleStoriesDTO []dto.SingleStoryFullDTO) []model.SingleStory{
	var singleStories []model.SingleStory
	for i := 0; i < len(singleStoriesDTO); i++ {
		singleStories=append(singleStories,convertSingleStoryDTOToSingleStory(singleStoriesDTO[i]))
	}
	return singleStories
}

func convertSingleStoryDTOToSingleStory(singleStoryDTO dto.SingleStoryFullDTO) model.SingleStory{
	var storyType model.StoryType
	if singleStoryDTO.Type=="CLOSE_FRIENDS" {
		storyType=model.CLOSE_FRIENDS
	} else if singleStoryDTO.Type=="PUBLIC"{
		storyType=model.PUBLIC
	} else if singleStoryDTO.Type=="ALL_FRIENDS"{
		storyType=model.ALL_FRIENDS
	}
	layout := "2006-01-02T15:04:05.000Z"
	creationDate,_ := time.Parse(layout,singleStoryDTO.CreationDate)
	var singleStory= model.SingleStory{
		Story:	model.Story{
			ID:           singleStoryDTO.ID,
			Description:  singleStoryDTO.Description,
			CreationDate: creationDate,
			UserId:       singleStoryDTO.UserId,
			LocationId:   singleStoryDTO.LocationId,
			IsDeleted:    singleStoryDTO.IsDeleted,
			IsExpired:    singleStoryDTO.IsExpired,
			Type:         storyType,
		},
	}
	return singleStory
}

//FIDSINGSTRYFORID9102
func (handler *SingleStoryHandler) FindSingleStoryForId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	singleStory := handler.SingleStoryService.FindSingleStoryForId(uuid.MustParse(id))
	singleStoryJson, _ := json.Marshal(singleStory)
	if singleStoryJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "SingleStoryHandler",
			"action":   "FIDSINGSTRYFORID9102",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded single story for id!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(singleStoryJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "SingleStoryHandler",
		"action":   "FIDSINGSTRYFORID9102",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding single story for id!")
	w.WriteHeader(http.StatusBadRequest)
}
