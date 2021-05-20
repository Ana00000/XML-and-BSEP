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

	ClassicUserCloseFriendsService * userService.ClassicUserCloseFirendsService
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






// returns all VALID stories from FOLLOWING users (FOR HOMEPAGE)
func (handler *SingleStoryHandler) FindAllFollowingStories(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")


	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	// returns only valid users
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	// retuns only valid FOLLOWINGS
	var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)

	// returns NOT DELETED STORIES from valid following users
	var stories = handler.SingleStoryService.FindAllFollowingStories(followings)

	//finds all conents
	var contents = handler.StoryContentService.FindAllContentsForStories(stories)


	//finds all locations
	var locations = handler.LocationService.FindAllLocationsForStories(stories)

	//find all tags
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)

	//creates a list of dtos
	var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}



// FIND SELECTED STORY BY ID (ONLY IF NOT DELETED)!
// IF PUBLIC/ IF FOLLOWING PRIVATE PROFILE
func (handler *SingleStoryHandler) FindSelectedStoryByIdForRegisteredUsers(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var story = handler.SingleStoryService.FindByID(uuid.MustParse(id))
	if story == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(story.UserId)
	if profileSettings.UserVisibility == settingsModel.PUBLIC_VISIBILITY{
		// EVERYONE CAN SELECT THIS STORY
		//finds all conents
		var contents = handler.StoryContentService.FindAllContentsForStory(story)

		//finds all locations
		var locations = handler.LocationService.FindAllLocationsForStory(story)

		//find all tags
		var tags = handler.StoryTagStoriesService.FindAllTagsForStory(story)

		//creates a list of dtos
		var storyDTO = handler.CreateStoryDTO(story,contents,locations,tags)

		storyJson, _ := json.Marshal(storyDTO)
		w.Write(storyJson)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		// CHECK IF LOGID FOLLOWING STORY USERID
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), story.UserId)
		if checkIfFollowing == true{

			//finds all conents
			var contents = handler.StoryContentService.FindAllContentsForStory(story)

			//finds all locations
			var locations = handler.LocationService.FindAllLocationsForStory(story)

			//find all tags
			var tags = handler.StoryTagStoriesService.FindAllTagsForStory(story)

			//creates a list of dtos
			var storyDTO = handler.CreateStoryDTO(story,contents,locations,tags)
			storyJson, _ := json.Marshal(storyDTO)
			w.Write(storyJson)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		}else{
			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}

}


//doneeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee

// NEREGISTROVANI

//// tab PUBLIC STORIES kada neregistroviani korisnik otvori sve PUBLIC, NOT EXPIRED I OD PUBLIC USERA
func (handler *SingleStoryHandler) FindAllPublicStoriesNotRegisteredUser(w http.ResponseWriter, r *http.Request) {

	// returns only VALID users
	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	// returns all PUBLIC users
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)

	// returns all STORIES of public and valid users
	var publicValidStories = handler.SingleStoryService.FindAllPublicStoriesNotRegisteredUser(allPublicUsers)

	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	//finds all conents
	var contents = handler.StoryContentService.FindAllContentsForStories(publicValidStories)

	//finds all locations
	var locations = handler.LocationService.FindAllLocationsForStories(publicValidStories)

	//find all tags
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(publicValidStories)

	//creates a list of dtos
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

	//finds all stories
	var stories = handler.SingleStoryService.FindAllStoriesForUserNotReg(uuid.MustParse(id))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	//finds all conents
	var contents = handler.StoryContentService.FindAllContentsForStories(stories)


	//finds all locations
	var locations = handler.LocationService.FindAllLocationsForStories(stories)

	//find all tags
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)

	//creates a list of dtos
	var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)


	storiesJson, _ := json.Marshal(storiesDTOS)
	w.Write(storiesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")


}

func (handler *SingleStoryHandler) FindAllPublicStoriesRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	// returns only VALID users but loggedIn user
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	// returns all PUBLIC users
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)

	// returns all STORIES of public and valid users
	var publicValidStories = handler.SingleStoryService.FindAllPublicStoriesNotRegisteredUser(allPublicUsers)


	//finds all conents
	var contents = handler.StoryContentService.FindAllContentsForStories(publicValidStories)

	//finds all locations
	var locations = handler.LocationService.FindAllLocationsForStories(publicValidStories)

	//find all tags
	var tags = handler.StoryTagStoriesService.FindAllTagsForStories(publicValidStories)

	//creates a list of dtos
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

	var checkIfValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(id))
	if  checkIfValid == false {
		fmt.Println("User NOT valid")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	if profileSettings.UserVisibility == settingsModel.PRIVATE_VISIBILITY{
		fmt.Println("User IS PRIVATE")

		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		if checkIfFollowing == true{

			var checkIfCloseFriend = handler.ClassicUserCloseFirends.CheckIfCloseFriend(uuid.MustParse(id), uuid.MustParse(logId))

			if checkIfCloseFriend = true{

				var stories = handler.SingleStoryService.FindAllStoriesForUserPrivateCloseFriend(uuid.MustParse(id))
				//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

				//finds all conents
				var contents = handler.StoryContentService.FindAllContentsForStories(stories)

				//finds all locations
				var locations = handler.LocationService.FindAllLocationsForStories(stories)

				//find all tags
				var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)

				//creates a list of dtos
				var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

				storiesJson, _ := json.Marshal(storiesDTOS)
				w.Write(storiesJson)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")

			}
			else{
				var stories = handler.SingleStoryService.FindAllStoriesForUserPrivate(uuid.MustParse(id))
				//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

				//finds all conents
				var contents = handler.StoryContentService.FindAllContentsForStories(stories)

				//finds all locations
				var locations = handler.LocationService.FindAllLocationsForStories(stories)

				//find all tags
				var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)

				//creates a list of dtos
				var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

				storiesJson, _ := json.Marshal(storiesDTOS)
				w.Write(storiesJson)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
			}


		}else{

			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}else{
		var stories = handler.SingleStoryService.FindAllStoriesForUserPublic(uuid.MustParse(id))
		//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

		//finds all conents
		var contents = handler.StoryContentService.FindAllContentsForStories(stories)


		//finds all locations
		var locations = handler.LocationService.FindAllLocationsForStories(stories)

		//find all tags
		var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)

		//creates a list of dtos
		var storiesDTOS = handler.CreateStoriesDTOList(stories,contents,locations,tags)

		storiesJson, _ := json.Marshal(storiesDTOS)
		w.Write(storiesJson)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

	}
}




//DTOS
func (handler *SingleStoryHandler) CreateStoriesDTOList(stories []model.SingleStory, contents []contentModel.SingleStoryContent, locations []locationModel.Location, tags []tagsModel.StoryTagStories) []dto.SelectedStoryDTO {
	var listOfStoriesDTOs []dto.SelectedStoryDTO

	for i := 0; i < len(stories); i++ {
		var storyDTO dto.SelectedStoryDTO
		storyDTO.StoryId = stories[i].ID
		storyDTO.Description = stories[i].Description
		storyDTO.CreationDate = stories[i].CreationDate



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
