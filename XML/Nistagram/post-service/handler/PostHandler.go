package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	settingsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type PostHandler struct {
	PostService *service.PostService
	ClassicUserService * userService.ClassicUserService
	ClassicUserFollowingsService * userService.ClassicUserFollowingsService
	ProfileSettings *settingsService.ProfileSettingsService
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

// FALI PROVERA ZA FINDALLPOSTSFORUSER DA LI JE PUBLIC USER (ONDA MOGU SVI DA VIDE) I AKO NIJE PUBLIC USER ONDA MORA DA GA PRATI DA BI VIDEO NJEGOVE POSTOVE

// for selected user (you can only select VALID users)
func (handler *PostHandler) FindAllPostsForUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var checIfValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(id))
	if  checIfValid == false {
		fmt.Println("User NOT valid")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var posts = handler.PostService.FindAllPostsForUser(uuid.MustParse(id))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	postsJson, _ := json.Marshal(posts)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

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