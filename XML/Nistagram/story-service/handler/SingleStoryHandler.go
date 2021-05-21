package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	contentModel "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	contentService "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	locationModel "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	locationService "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	settingsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	tagsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	tagsService "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type SingleStoryHandler struct {
	SingleStoryService * service.SingleStoryService
	StoryService * service.StoryService

	ClassicUserService * userService.ClassicUserService
	ClassicUserFollowingsService * userService.ClassicUserFollowingsService
	ProfileSettings *settingsService.ProfileSettingsService
	StoryContentService *contentService.SingleStoryContentService
	LocationService *locationService.LocationService
	StoryTagStoriesService *tagsService.StoryTagStoriesService
	TagService *tagsService.TagService

	ClassicUserCloseFriendsService * userService.ClassicUserCloseFriendsService
}

func (handler *SingleStoryHandler) CreateSingleStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.StoryService.CreateStory(&singleStory.Story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	singleStoryIDJson, _ := json.Marshal(singleStory.ID)
	w.Write(singleStoryIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

// NEREGISTROVANI

//// tab PUBLIC STORIES kada neregistroviani korisnik otvori sve PUBLIC, NOT EXPIRED I OD PUBLIC USERA
func (handler *SingleStoryHandler) FindAllPublicStoriesNotRegisteredUser(w http.ResponseWriter, r *http.Request) {


	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidStories = handler.SingleStoryService.FindAllPublicStoriesNotRegisteredUser(allPublicUsers)
	var contents = handler.StoryContentService.FindAllContentsForStories(publicValidStories)
	var locations = handler.LocationService.FindAllLocationsForStories(publicValidStories)
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(publicValidStories)
	var storiesDTOS = handler.CreateStoriesDTOList(publicValidStories,contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


// kada neregistrovani user otvori PUBLIC usera sa spiska i onda na njegovom profilu vidi PUBLIC i NOT EXPIRED STORIJE
func (handler *SingleStoryHandler) FindAllStoriesForUserNotRegisteredUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var checkIfValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(id))
	if  checkIfValid == false {
		fmt.Println("User NOT valid")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	fmt.Println("User IS valid")
	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	if profileSettings.UserVisibility == settingsModel.PRIVATE_VISIBILITY{
		fmt.Println("User IS PRIVATE")
		w.WriteHeader(http.StatusExpectationFailed)
	}


	var stories = handler.SingleStoryService.FindAllStoriesForUserNotReg(uuid.MustParse(id))
	var contents = handler.StoryContentService.FindAllContentsForStories(stories)
	var locations = handler.LocationService.FindAllLocationsForStories(stories)
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
	var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)


	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")


}

func (handler *SingleStoryHandler) FindAllPublicStoriesRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidStories = handler.SingleStoryService.FindAllPublicStoriesNotRegisteredUser(allPublicUsers)
	var contents = handler.StoryContentService.FindAllContentsForStories(publicValidStories)
	var locations = handler.LocationService.FindAllLocationsForStories(publicValidStories)
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(publicValidStories)
	var storiesDTOS = handler.CreateStoriesDTOList(publicValidStories,contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


// metoda koja se poziva kada registrovani user udje na profil nekog usera
func (handler *SingleStoryHandler) FindAllStoriesForUserRegisteredUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")
	var stories []model.SingleStory

	var checkIfValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(id))
	if  checkIfValid == false {
		fmt.Println("User NOT valid")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	if profileSettings.UserVisibility == settingsModel.PRIVATE_VISIBILITY{

		//PRIVATE USER

		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		if checkIfFollowing == true{

			// PRATI GA

			var checkIfCloseFriend = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(uuid.MustParse(id), uuid.MustParse(logId))
			if checkIfCloseFriend == true{

				// NALAZI SE U CLOSE FRIENDS

				stories = handler.SingleStoryService.FindAllStoriesForUserCloseFriend(uuid.MustParse(id))


			} else{
				// NE NALAZI SE U CLOSE FRIENDS ALI GA PRATI - PUBLIC I ALL FRIENDS
				stories = handler.SingleStoryService.FindAllStoriesForUserPublicAllFriends(uuid.MustParse(id))
			}

			var contents = handler.StoryContentService.FindAllContentsForStories(stories)
			var locations = handler.LocationService.FindAllLocationsForStories(stories)
			var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
			var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

			storiesJson, _ := json.Marshal(storiesDTOS)
			w.Write(storiesJson)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")


		}else{
			// PRIVATE USER I NE PRATI GA = NE MOZE NISTA DA VIDI

			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}else{
		//PUBLIC USER
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		if checkIfFollowing == true{

			// PRATI GA
			var checkIfCloseFriend = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(uuid.MustParse(id), uuid.MustParse(logId))
			if checkIfCloseFriend == true{

				// NALAZI SE U CLOSE FRIENDS
				stories = handler.SingleStoryService.FindAllStoriesForUserCloseFriend(uuid.MustParse(id))


			} else{
				// NE NALAZI SE U CLOSE FRIENDS ALI GA PRATI - PUBLIC I ALL FRIENDS
				stories = handler.SingleStoryService.FindAllStoriesForUserPublicAllFriends(uuid.MustParse(id))
			}

		}else{
			//NE PRATI GA ALI POSTO JE PUBLIC SME DA VIDI PUBLIC STORIJE
			stories = handler.SingleStoryService.FindAllStoriesForUserPublic(uuid.MustParse(id))
		}
		var contents = handler.StoryContentService.FindAllContentsForStories(stories)
		var locations = handler.LocationService.FindAllLocationsForStories(stories)
		var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
		var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

		storiesJson, _ := json.Marshal(storiesDTOS)
		w.Write(storiesJson)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")


	}
}


// returns all VALID stories from FOLLOWING users (FOR HOMEPAGE)
func (handler *SingleStoryHandler) FindAllFollowingStories(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)
	var allValidStories []model.SingleStory
	var stories = handler.SingleStoryService.FindAllFollowingStories(followings)


	for i:=0; i<len(stories);i++{
		if stories[i].Type == model.PUBLIC || stories[i].Type == model.ALL_FRIENDS{

			allValidStories = append(allValidStories, stories[i])

		}else if stories[i].Type == model.CLOSE_FRIENDS{

			var checkIfCloseFriend = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(stories[i].UserId, uuid.MustParse(id))
			if checkIfCloseFriend == true{

				allValidStories = append(allValidStories, stories[i])
			}
		}
	}

	var contents = handler.StoryContentService.FindAllContentsForStories(stories)
	var locations = handler.LocationService.FindAllLocationsForStories(stories)
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
	var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

// FIND SELECTED STORY BY ID (ONLY IF NOT DELETED)!
func (handler *SingleStoryHandler) FindSelectedStoryByIdForRegisteredUsers(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var story = handler.SingleStoryService.FindByID(uuid.MustParse(id))
	if story == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if story.IsDeleted == true{

		fmt.Println("Deleted story")
		w.WriteHeader(http.StatusExpectationFailed)
		return

	}

	if story.UserId != uuid.MustParse(logId){
		//POSTMAN CHECK
		//NIJE STORI OD ULOGOVANOG USERA

		fmt.Println("Unavailable story to this user")
		w.WriteHeader(http.StatusExpectationFailed)
		return

	}


	var contents = handler.StoryContentService.FindAllContentsForStory(story)
	var locations = handler.LocationService.FindAllLocationsForStory(story)
	var tags = handler.StoryTagStoriesService.FindAllTagsForStory(story)
	var storyDTO = handler.CreateStoryDTO(story,contents,locations,tags)

	storyJson, _ := json.Marshal(storyDTO)
	w.Write(storyJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")


}


// all stories (EXCEPT DELETED) for my current logged in user (expired and not expired, public, all_friend, close friends)
func (handler *SingleStoryHandler) FindAllStoriesForLoggedUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var stories = handler.SingleStoryService.FindAllStoriesForLoggedUser(uuid.MustParse(id))
	var contents = handler.StoryContentService.FindAllContentsForStories(stories)
	var locations = handler.LocationService.FindAllLocationsForStories(stories)
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
	var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

//DTOS
func (handler *SingleStoryHandler) CreateStoriesDTOList(stories []model.SingleStory, contents []contentModel.SingleStoryContent, locations []locationModel.Location, tags []tagsModel.StoryTagStories) []dto.SelectedStoryDTO {
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

				if contents[j].Type == contentModel.VIDEO{
					storyDTO.Type = "VIDEO"
				}else if contents[j].Type == contentModel.PICTURE{
					storyDTO.Type = "PICTURE"
				}

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
				listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].StoryTagId))
			}
		}

		storyDTO.Tags = listOfTags

		listOfStoriesDTOs = append(listOfStoriesDTOs, storyDTO)

	}

	return listOfStoriesDTOs

}

func (handler *SingleStoryHandler) CreateStoryDTO(story *model.SingleStory, contents []contentModel.SingleStoryContent, locations []locationModel.Location, tags []tagsModel.StoryTagStories) dto.SelectedStoryDTO {


	var storyDTO dto.SelectedStoryDTO
	fmt.Println("STORIES")
	storyDTO.StoryId = story.ID
	storyDTO.Description = story.Description
	storyDTO.CreationDate = story.CreationDate

	for j := 0; j < len(contents); j++ {
		if contents[j].SingleStoryId == story.ID {
			storyDTO.Path = contents[j].Path

			if contents[j].Type == contentModel.VIDEO{
				storyDTO.Type = "VIDEO"
			}else if contents[j].Type == contentModel.PICTURE{
				storyDTO.Type = "PICTURE"
			}
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
			listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].StoryTagId))

		}
	}

	storyDTO.Tags = listOfTags


	return storyDTO

}

