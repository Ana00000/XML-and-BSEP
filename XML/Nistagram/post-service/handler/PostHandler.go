package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type PostHandler struct {
	PostService *service.PostService
	ClassicUserService * userService.ClassicUserService
	ClassicUserFollowingsService * userService.ClassicUserFollowingsService
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

	// returns POSTS from valid following users
	var posts = handler.PostService.FindAllFollowingPosts(followings)

	postsJson, _ := json.Marshal(posts)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}