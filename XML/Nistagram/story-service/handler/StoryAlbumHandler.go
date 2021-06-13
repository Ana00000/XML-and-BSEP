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

type StoryAlbumHandler struct {
	Service * service.StoryAlbumService
	StoryService * service.StoryService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}
//CRSTRYALB8542
func (handler *StoryAlbumHandler) CreateStoryAlbum(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "CRSTRYALB8542",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-story-album-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "CRSTRYALB8542",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "CRSTRYALB8542",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	var storyAlbumDTO dto.StoryAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "CRSTRYALB8542",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryAlbumDTO!")
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
		Story: model.Story{
			ID:           id,
			CreationDate: time.Now(),
			Description:  storyAlbumDTO.Description,
			UserId:       storyAlbumDTO.UserId,
			LocationId:   storyAlbumDTO.LocationId,
			IsDeleted:    false,
			Type:         storyAlbumType,
		},
	}

	err = handler.Service.CreateStoryAlbum(&storyAlbum)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "CRSTRYALB8542",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating story album!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.StoryService.CreateStory(&storyAlbum.Story)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "CRSTRYALB8542",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating basic story!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	storyAlbumIDJson, _ := json.Marshal(storyAlbum.ID)
	w.Write(storyAlbumIDJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumHandler",
		"action":   "CRSTRYALB8542",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created story album!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
//FIDALALBSTORISFORLOGGUS8293
func (handler *StoryAlbumHandler) FindAllAlbumStoriesForLoggedUser(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-album-stories-for-logged-user-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	id := r.URL.Query().Get("id")

	var albumStories = handler.Service.FindAllAlbumStoriesForUser(uuid.MustParse(id))
	//var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbums(albumStories)
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_story_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoryAlbumsDTO, _ := json.Marshal(albumStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoryAlbumsDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoryAlbumsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for story album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.StoryAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var locations = handler.LocationService.FindAllLocationsForStoryAlbums(albumStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_story_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(albumStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for story album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(albumStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_story_album_tag_story_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(albumStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for story album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(publicValidAlbumStories)
	var tags []dto.StoryAlbumTagStoryAlbumsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumTagStoryAlbumsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var albumStoriesDTOS = handler.CreateStoryAlbumsDTOList(albumStories, contents, locations, tags)

	albumStoriesJson, _ := json.Marshal(albumStoriesDTOS)
	w.Write(albumStoriesJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumHandler",
		"action":   "FIDALALBSTORISFORLOGGUS8293",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all album stories for logged user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
//CRSTRYALBMSDTOLST0330
func (handler *StoryAlbumHandler) CreateStoryAlbumsDTOList(albums []model.StoryAlbum, contents []dto.StoryAlbumContentFullDTO, locations []dto.LocationDTO, tags []dto.StoryAlbumTagStoryAlbumsDTO) []dto.SelectedStoryAlbumDTO {
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
				storyAlbumDTO.Type = append(storyAlbumDTO.Type, contents[j].Type)
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
				var returnValueTagName ReturnValueString
				reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"), tags[p].TagId.String())
				err := getJson(reqUrl, &returnValueTagName)
				if err!=nil{
					handler.LogError.WithFields(logrus.Fields{
						"status": "failure",
						"location":   "StoryAlbumHandler",
						"action":   "CRSTRYALBMSDTOLST0330",
						"timestamp":   time.Now().String(),
					}).Error("Failed finding tag name by id or wrong cast json to ReturnValueString!")
					//fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
					return nil
				}
				listOfTags = append(listOfTags, returnValueTagName.ReturnValue)

				//listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))
			}
		}

		storyAlbumDTO.Tags = listOfTags
		listOfStoryAlbumsDTOs = append(listOfStoryAlbumsDTOs, storyAlbumDTO)
	}

	return listOfStoryAlbumsDTOs
}
//FIDSELECTSTRYALBBYIDFORLOGGUS983
func (handler *StoryAlbumHandler) FindSelectedStoryAlbumByIdForLoggedUser(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-selected-story-album-by-id-for-logged-user-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}
	/*
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	id := r.URL.Query().Get("id")       //story album id
	logId := r.URL.Query().Get("logId") //loged user id

	var storyAlbum = handler.Service.FindByID(uuid.MustParse(id))
	if storyAlbum == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Story album not found!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if storyAlbum.IsDeleted == true{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Story album is deleted!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if storyAlbum.UserId != uuid.MustParse(logId){
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Story album doesnt belong to user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	storyAlbumFullDTO := convertStoryAlbumToStoryAlbumDTO(*storyAlbum)
	///find_all_contents_for_story_album/
	//var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbum(storyAlbum)
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_story_album/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoryAlbumsDTO, _ := json.Marshal(storyAlbumFullDTO)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoryAlbumsDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoryAlbumsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for story album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.StoryAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	///find_locations_for_story_album/ POST LOCATON
	//var locations = handler.LocationService.FindAllLocationsForStoryAlbum(storyAlbum)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_story_album/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(storyAlbumFullDTO)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for story album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	///find_all_tags_for_story_album/ POST TAG
	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbum(storyAlbum)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_story_album_tag_story_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(storyAlbumFullDTO)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for story album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.StoryAlbumTagStoryAlbumsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumTagStoryAlbumsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var storyAlbumDTO = handler.CreateStoryAlbumDTO(storyAlbum, contents, locations, tags)

	storyAlbumJson, _ := json.Marshal(storyAlbumDTO)
	w.Write(storyAlbumJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumHandler",
		"action":   "FIDSELECTSTRYALBBYIDFORLOGGUS983",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded selected story album by ID for logged user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
//CRSTRYALBDTO9810
func (handler *StoryAlbumHandler) CreateStoryAlbumDTO(album *model.StoryAlbum, contents []dto.StoryAlbumContentFullDTO, locations []dto.LocationDTO, tags []dto.StoryAlbumTagStoryAlbumsDTO) dto.SelectedStoryAlbumDTO {
	var storyAlbumDTO dto.SelectedStoryAlbumDTO

	storyAlbumDTO.StoryAlbumId = album.ID
	storyAlbumDTO.Description = album.Description
	storyAlbumDTO.CreationDate = album.CreationDate

	for j := 0; j < len(contents); j++ {
		if contents[j].StoryAlbumId == album.ID {
			storyAlbumDTO.Path = append(storyAlbumDTO.Path, contents[j].Path)
			storyAlbumDTO.Type = append(storyAlbumDTO.Type, contents[j].Type)
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
			var returnValueTagName ReturnValueString
			reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"), tags[p].TagId.String())
			err := getJson(reqUrl, &returnValueTagName)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "StoryAlbumHandler",
					"action":   "CRSTRYALBDTO9810",
					"timestamp":   time.Now().String(),
				}).Error("Failed finding tag name by ID or wrong cast json to ReturnValueString!")
				panic(err)
			}
			listOfTags = append(listOfTags, returnValueTagName.ReturnValue)
			//listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))
		}
	}

	storyAlbumDTO.Tags = listOfTags
	return storyAlbumDTO
}
//FIDALPUBALBSTORISREGUS9012
func (handler *StoryAlbumHandler) FindAllPublicAlbumStoriesRegisteredUser(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-public-album-stories-registered-user-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}
	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	id := r.URL.Query().Get("id")

	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all classic users but logged in or wrong cast json to list ClassicUserDTO!")
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
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
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
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ClassicUserDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var publicValidStoryAlbums = handler.Service.FindAllPublicAlbumStoriesNotRegisteredUser(allPublicUsers)
	//var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbums(publicValidStoryAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_story_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoryAlbumsDTO, _ := json.Marshal(allPublicUsers)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoryAlbumsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoryAlbumsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.StoryAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var locations = handler.LocationService.FindAllLocationsForStoryAlbums(publicValidStoryAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_story_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidStoryAlbums)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(publicValidStoryAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_story_album_tag_story_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidStoryAlbums)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(publicValidAlbumStories)
	var tags []dto.StoryAlbumTagStoryAlbumsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISREGUS9012",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumTagStoryAlbumsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var storyAlbumsDTOS = handler.CreateStoryAlbumsDTOList(publicValidStoryAlbums, contents, locations, tags)

	storyAlbumsJson, _ := json.Marshal(storyAlbumsDTOS)
	w.Write(storyAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumHandler",
		"action":   "FIDALPUBALBSTORISREGUS9012",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all public album stories registered user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
//FIDALPUBALBSTORISNOTREGUS9021
func (handler *StoryAlbumHandler) FindAllPublicAlbumStoriesNotRegisteredUser(w http.ResponseWriter, r *http.Request) {

	//var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_valid_users/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all valid users or wrong cast json to list ClassicUserDTO!")
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
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
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
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ClassicUserDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	var publicValidAlbumStories = handler.Service.FindAllPublicAlbumStoriesNotRegisteredUser(allPublicUsers)

	//var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbums(publicValidAlbumStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_story_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoryAlbumsDTO, _ := json.Marshal(publicValidAlbumStories)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoryAlbumsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoryAlbumsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding contents for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.StoryAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var locations = handler.LocationService.FindAllLocationsForStoryAlbums(publicValidAlbumStories)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_story_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidAlbumStories)
	//fmt.Println("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding locations for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_story_album_tag_story_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidAlbumStories)
	//fmt.Println("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding tags for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(publicValidAlbumStories)
	var tags []dto.StoryAlbumTagStoryAlbumsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALPUBALBSTORISNOTREGUS9021",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumTagStoryAlbumsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var storyAlbumsDTOS = handler.CreateStoryAlbumsDTOList(publicValidAlbumStories, contents, locations, tags)

	storyAlbumsJson, _ := json.Marshal(storyAlbumsDTOS)
	w.Write(storyAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumHandler",
		"action":   "FIDALPUBALBSTORISNOTREGUS9021",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all public album stories not registered user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
//FIDALFOLLINGSTRYALBMS0910
// returns all VALID story albums from FOLLOWING users (FOR HOMEPAGE)
func (handler *StoryAlbumHandler) FindAllFollowingStoryAlbums(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-following-story-albums-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	id := r.URL.Query().Get("id")

	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all classic users but logged in or wrong cast json to list ClassicUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_valid_followings_for_user/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	//fmt.Println("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
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
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ClassicUserFollowingsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var allValidStoryAlbums []model.StoryAlbum
	var storyAlbums = handler.Service.FindAllFollowingStoryAlbums(followings)

	for i := 0; i < len(storyAlbums); i++ {
		if storyAlbums[i].Type == model.PUBLIC || storyAlbums[i].Type == model.ALL_FRIENDS {

			allValidStoryAlbums = append(allValidStoryAlbums, storyAlbums[i])

		} else if storyAlbums[i].Type == model.CLOSE_FRIENDS {
			//var checkIfCloseFriend = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(storyAlbums[i].UserId, uuid.MustParse(id))
			var returnValueCloseFriend ReturnValueBool
			reqUrl = fmt.Sprintf("http://%s:%s/check_if_close_friend/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), storyAlbums[i].UserId, id)
			err = getJson(reqUrl, &returnValueCloseFriend)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "StoryAlbumHandler",
					"action":   "FIDALFOLLINGSTRYALBMS0910",
					"timestamp":   time.Now().String(),
				}).Error("Failed checking if close friend or wrong cast json to list ReturnValueBool!")
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			checkIfCloseFriend := returnValueCloseFriend.ReturnValue

			if checkIfCloseFriend == true {

				allValidStoryAlbums = append(allValidStoryAlbums, storyAlbums[i])
			}
		}
	}
	var allValidStoryAlbumsDTOs = convertListStoryAlbumsToStoryAlbumsDTO(allValidStoryAlbums)
	//var contents = handler.StoryAlbumContentService.FindAllContentsForStoryAlbums(storyAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_story_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidStoryAlbumsDTO, _ := json.Marshal(allValidStoryAlbumsDTOs)
	//fmt.Println("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonValidStoryAlbumsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoryAlbumsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all contents for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.StoryAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var locations = handler.LocationService.FindAllLocationsForStoryAlbums(storyAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_story_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(allValidStoryAlbumsDTOs)
	//fmt.Println("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all locations for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(storyAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_story_album_tag_story_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(allValidStoryAlbumsDTOs)
	//fmt.Println("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding all tags for story albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	//var tags = handler.StoryAlbumTagStoryAlbumsService.FindAllTagsForStoryAlbumTagStoryAlbums(publicValidAlbumStories)
	var tags []dto.StoryAlbumTagStoryAlbumsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALFOLLINGSTRYALBMS0910",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumTagStoryAlbumsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var storyAlbumsDTOS = handler.CreateStoryAlbumsDTOList(storyAlbums, contents, locations, tags)

	storyAlbumsJson, _ := json.Marshal(storyAlbumsDTOS)
	w.Write(storyAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumHandler",
		"action":   "FIDALFOLLINGSTRYALBMS0910",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all following story albums!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func convertListStoryAlbumsToStoryAlbumsDTO(storyAlbums []model.StoryAlbum) []dto.StoryAlbumFullDTO {
	var storyAlbumsDTO []dto.StoryAlbumFullDTO
	for i := 0; i < len(storyAlbums); i++ {
		storyAlbumsDTO = append(storyAlbumsDTO, convertStoryAlbumToStoryAlbumDTO(storyAlbums[i]))
	}
	return storyAlbumsDTO
}

func convertStoryAlbumToStoryAlbumDTO(storyAlbum model.StoryAlbum) dto.StoryAlbumFullDTO {
	layout := "2006-01-02T15:04:05.000Z"
	var storyAlbumDTO = dto.StoryAlbumFullDTO{
		ID:           storyAlbum.ID,
		Description:  storyAlbum.Description,
		CreationDate: storyAlbum.CreationDate.Format(layout),
		UserId:       storyAlbum.UserId,
		LocationId:   storyAlbum.LocationId,
		IsDeleted:    storyAlbum.IsDeleted,
	}
	return storyAlbumDTO
}

func convertStoryAlbumsDTOToListStoryAlbums(storyAlbumsDTO []dto.StoryAlbumFullDTO) []model.StoryAlbum {
	var storyAlbums []model.StoryAlbum
	for i := 0; i < len(storyAlbumsDTO); i++ {
		storyAlbums = append(storyAlbums, convertStoryAlbumDTOToStoryAlbum(storyAlbumsDTO[i]))
	}
	return storyAlbums
}

func convertStoryAlbumDTOToStoryAlbum(storyAlbumDTO dto.StoryAlbumFullDTO) model.StoryAlbum {
	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, storyAlbumDTO.CreationDate)
	var storyAlbum = model.StoryAlbum{
		Story: model.Story{
			ID:           storyAlbumDTO.ID,
			Description:  storyAlbumDTO.Description,
			CreationDate: creationDate,
			UserId:       storyAlbumDTO.UserId,
			LocationId:   storyAlbumDTO.LocationId,
			IsDeleted:    storyAlbumDTO.IsDeleted,
			IsExpired:    storyAlbumDTO.IsExpired,
			Type:         0,
		},
	}
	return storyAlbum
}
