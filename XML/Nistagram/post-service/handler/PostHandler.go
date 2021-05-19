package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	contentModel "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	contentService "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	locationModel "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	locationService "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	settingsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	tagsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	tagsService "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type PostHandler struct {
	PostService *service.PostService
	ClassicUserService * userService.ClassicUserService
	ClassicUserFollowingsService * userService.ClassicUserFollowingsService
	ProfileSettings *settingsService.ProfileSettingsService
	PostContentService *contentService.SinglePostContentService
	LocationService *locationService.LocationService
	PostTagPostsService *tagsService.PostTagPostsService
	TagService *tagsService.TagService
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostDTO
	err := json.NewDecoder(r.Body).Decode(&postDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post := model.Post{
		ID:           uuid.UUID{},
		Description:  postDTO.Description,
		CreationDate: time.Now(),
		UserID:       postDTO.UserID,
		LocationId:   postDTO.LocationID,
		IsDeleted:    false,
	}

	err = handler.PostService.CreatePost(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostUpdateDTO

	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.PostService.UpdatePost(&postDTO)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


// for selected user (you can only select VALID users)
func (handler *PostHandler) FindAllPostsForUserNotRegisteredUser(w http.ResponseWriter, r *http.Request) {
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

	//finds all posts
	var posts = handler.PostService.FindAllPostsForUser(uuid.MustParse(id))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	//finds all conents
	var contents = handler.PostContentService.FindAllContentsForPosts(posts)


	//finds all locations
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)

	//find all tags
	var tags = handler.PostTagPostsService.FindAllTagsForPosts(posts)

	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(posts,contents,locations,tags)


	postsJson, _ := json.Marshal(postsDTOS)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")


}


func (handler *PostHandler) CreatePostsDTOList(posts []model.Post, contents []contentModel.SinglePostContent, locations []locationModel.Location, tags []tagsModel.PostTagPosts) []dto.SelectedPostDTO {
	var listOfPostsDTOs []dto.SelectedPostDTO

	for i := 0; i < len(posts); i++ {
		var postDTO dto.SelectedPostDTO
		fmt.Println("POSTS")
		postDTO.PostId = posts[i].ID
		postDTO.Description = posts[i].Description
		postDTO.CreationDate = posts[i].CreationDate

		for j := 0; j < len(contents); j++ {
			if contents[j].SinglePostId == posts[i].ID {
				postDTO.Path = contents[i].Path
			}
		}

		for k := 0; k < len(locations); k++ {
			if locations[k].ID == posts[i].LocationId {
				postDTO.LocationId = locations[k].ID
				postDTO.City = locations[k].City
				postDTO.Country = locations[k].Country
				postDTO.StreetName = locations[k].StreetName
				postDTO.StreetNumber = locations[k].StreetNumber
			}
		}

		var listOfTags []string
		for p := 0; p < len(tags); p++ {
			if tags[p].PostId == posts[i].UserID {
				listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].PostTagId))

			}
		}

		postDTO.Tags = listOfTags

		listOfPostsDTOs = append(listOfPostsDTOs, postDTO)

	}

	return listOfPostsDTOs

}




func (handler *PostHandler) FindAllPostsForUserRegisteredUser(w http.ResponseWriter, r *http.Request) {
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

		// CHECK IF LOGID FOLLOWING POST USERID
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPost(uuid.MustParse(logId), uuid.MustParse(id))
		if checkIfFollowing == true{
			var posts = handler.PostService.FindAllPostsForUser(uuid.MustParse(id))
			//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

			postsJson, _ := json.Marshal(posts)
			w.Write(postsJson)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

		}else{

			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}else{
		var posts = handler.PostService.FindAllPostsForUser(uuid.MustParse(id))
		//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

		postsJson, _ := json.Marshal(posts)
		w.Write(postsJson)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

	}
}


// returns all VALID posts from FOLLOWING users (FOR HOMEPAGE)
func (handler *PostHandler) FindAllFollowingPosts(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")


	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	// returns only valid users
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	// retuns only valid FOLLOWINGS
	var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)

	// returns NOT DELETED POSTS from valid following users
	var posts = handler.PostService.FindAllFollowingPosts(followings)

	postsJson, _ := json.Marshal(posts)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

// FIND SELECTED POST BY ID (ONLY IF NOT DELETED)!
// IF PUBLIC/ IF FOLLOWING PRIVATE PROFILE
func (handler *PostHandler) FindSelectedPostByIdForNotRegisteredUsers(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var post = handler.PostService.FindByID(uuid.MustParse(id))
	if post == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	if profileSettings.UserVisibility == settingsModel.PUBLIC_VISIBILITY{
		// EVERYONE CAN SELECT THIS POST
		postJson, _ := json.Marshal(post)
		w.Write(postJson)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		// FOR POSTMAN CHECK (should redirect)
		fmt.Println("Profile is private")
		w.WriteHeader(http.StatusExpectationFailed)
	}



}


// FIND SELECTED POST BY ID (ONLY IF NOT DELETED)!
// IF PUBLIC/ IF FOLLOWING PRIVATE PROFILE
func (handler *PostHandler) FindSelectedPostByIdForRegisteredUsers(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var post = handler.PostService.FindByID(uuid.MustParse(id))
	if post == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	if profileSettings.UserVisibility == settingsModel.PUBLIC_VISIBILITY{
		// EVERYONE CAN SELECT THIS POST
		postJson, _ := json.Marshal(post)
		w.Write(postJson)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		// CHECK IF LOGID FOLLOWING POST USERID
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPost(uuid.MustParse(logId), post.UserID)
		if checkIfFollowing == true{
			postJson, _ := json.Marshal(post)
			w.Write(postJson)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		}else{
			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}

}

//FIND ALL PUBLIC POSTS (for not registered users)
func (handler *PostHandler) FindAllPublicPostsNotRegisteredUser(w http.ResponseWriter, r *http.Request) {

	// returns only VALID users
	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()

	// returns all PUBLIC users
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)

	// returns all POSTS of public and valid users
	var publicValidPosts = handler.PostService.FindAllPublicPostsNotRegisteredUser(allPublicUsers)
	postJson, _ := json.Marshal(publicValidPosts)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


//FindAllPublicPostsRegisteredUser
func (handler *PostHandler) FindAllPublicPostsRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	// returns only VALID users but loggedIn user
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	// returns all PUBLIC users
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)

	// returns all POSTS of public and valid users
	var publicValidPosts = handler.PostService.FindAllPublicPostsNotRegisteredUser(allPublicUsers)
	postJson, _ := json.Marshal(publicValidPosts)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}