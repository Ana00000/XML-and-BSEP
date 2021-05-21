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

type SinglePostHandler struct {
	SinglePostService * service.SinglePostService
	PostService * service.PostService
	ClassicUserService * userService.ClassicUserService
	ClassicUserFollowingsService * userService.ClassicUserFollowingsService
	ProfileSettings *settingsService.ProfileSettingsService
	PostContentService *contentService.SinglePostContentService
	LocationService *locationService.LocationService
	PostTagPostsService *tagsService.PostTagPostsService
	TagService *tagsService.TagService
}

func (handler *SinglePostHandler) CreateSinglePost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	singlePost := model.SinglePost{
		Post : model.Post{
			ID: uuid.New(),
			Description: singlePostDTO.Description,
			CreationDate: time.Now(),
			UserID: singlePostDTO.UserID,
			LocationId: singlePostDTO.LocationID,
			IsDeleted: false,
		},
	}

	err = handler.SinglePostService.CreateSinglePost(&singlePost)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.PostService.CreatePost(&singlePost.Post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	singlePostIDJson, _ := json.Marshal(singlePost.ID)
	w.Write(singlePostIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

// for selected user (you can only select VALID users)
func (handler *SinglePostHandler) FindAllPostsForUserNotRegisteredUser(w http.ResponseWriter, r *http.Request) {
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
	var posts = handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id))
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

func (handler *SinglePostHandler) FindAllPostsForUserRegisteredUser(w http.ResponseWriter, r *http.Request) {
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
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		if checkIfFollowing == true{
			var posts = handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id))
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

		}else{

			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}else{
		var posts = handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id))
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
}


// returns all VALID posts from FOLLOWING users (FOR HOMEPAGE)
func (handler *SinglePostHandler) FindAllFollowingPosts(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")


	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	// returns only valid users
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	// retuns only valid FOLLOWINGS
	var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)

	// returns NOT DELETED POSTS from valid following users
	var posts = handler.SinglePostService.FindAllFollowingPosts(followings)

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

// FIND SELECTED POST BY ID (ONLY IF NOT DELETED)!
// IF PUBLIC/ IF FOLLOWING PRIVATE PROFILE
func (handler *SinglePostHandler) FindSelectedPostByIdForNotRegisteredUsers(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var post = handler.SinglePostService.FindByID(uuid.MustParse(id))
	if post == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	if profileSettings.UserVisibility == settingsModel.PUBLIC_VISIBILITY{
		// EVERYONE CAN SELECT THIS POST
		//finds all conents
		var contents = handler.PostContentService.FindAllContentsForPost(post)

		//finds all locations
		var locations = handler.LocationService.FindAllLocationsForPost(post)

		//find all tags
		var tags = handler.PostTagPostsService.FindAllTagsForPost(post)

		//creates a list of dtos
		var postDTO = handler.CreatePostDTO(post,contents,locations,tags)

		postJson, _ := json.Marshal(postDTO)
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
func (handler *SinglePostHandler) FindSelectedPostByIdForRegisteredUsers(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var post = handler.SinglePostService.FindByID(uuid.MustParse(id))
	if post == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	if profileSettings.UserVisibility == settingsModel.PUBLIC_VISIBILITY{
		// EVERYONE CAN SELECT THIS POST
		//finds all conents
		var contents = handler.PostContentService.FindAllContentsForPost(post)

		//finds all locations
		var locations = handler.LocationService.FindAllLocationsForPost(post)

		//find all tags
		var tags = handler.PostTagPostsService.FindAllTagsForPost(post)

		//creates a list of dtos
		var postDTO = handler.CreatePostDTO(post,contents,locations,tags)

		postJson, _ := json.Marshal(postDTO)
		w.Write(postJson)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		// CHECK IF LOGID FOLLOWING POST USERID
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), post.UserID)
		if checkIfFollowing == true{

			//finds all conents
			var contents = handler.PostContentService.FindAllContentsForPost(post)

			//finds all locations
			var locations = handler.LocationService.FindAllLocationsForPost(post)

			//find all tags
			var tags = handler.PostTagPostsService.FindAllTagsForPost(post)

			//creates a list of dtos
			var postDTO = handler.CreatePostDTO(post,contents,locations,tags)
			postJson, _ := json.Marshal(postDTO)
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
func (handler *SinglePostHandler) FindAllPublicPostsNotRegisteredUser(w http.ResponseWriter, r *http.Request) {

	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidPosts = handler.SinglePostService.FindAllPublicPostsNotRegisteredUser(allPublicUsers)
	var contents = handler.PostContentService.FindAllContentsForPosts(publicValidPosts)
	var locations = handler.LocationService.FindAllLocationsForPosts(publicValidPosts)
	var tags = handler.PostTagPostsService.FindAllTagsForPosts(publicValidPosts)


	var postsDTOS = handler.CreatePostsDTOList(publicValidPosts,contents,locations,tags)

	postJson, _ := json.Marshal(postsDTOS)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


//FindAllPublicPostsRegisteredUser
func (handler *SinglePostHandler) FindAllPublicPostsRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	// returns only VALID users but loggedIn user
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	// returns all PUBLIC users
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)

	// returns all POSTS of public and valid users
	var publicValidPosts = handler.SinglePostService.FindAllPublicPostsNotRegisteredUser(allPublicUsers)


	//finds all conents
	var contents = handler.PostContentService.FindAllContentsForPosts(publicValidPosts)

	//finds all locations
	var locations = handler.LocationService.FindAllLocationsForPosts(publicValidPosts)

	//find all tags
	var tags = handler.PostTagPostsService.FindAllTagsForPosts(publicValidPosts)

	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(publicValidPosts,contents,locations,tags)

	postJson, _ := json.Marshal(postsDTOS)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostHandler) CreatePostsDTOList(posts []model.SinglePost, contents []contentModel.SinglePostContent, locations []locationModel.Location, tags []tagsModel.PostTagPosts) []dto.SelectedPostDTO {
	var listOfPostsDTOs []dto.SelectedPostDTO

	for i := 0; i < len(posts); i++ {
		var postDTO dto.SelectedPostDTO
		postDTO.PostId = posts[i].ID
		postDTO.Description = posts[i].Description
		postDTO.CreationDate = posts[i].CreationDate

		for j := 0; j < len(contents); j++ {
			if contents[j].SinglePostId == posts[i].ID {
				postDTO.Path = contents[j].Path

				if contents[j].Type == contentModel.VIDEO{
					postDTO.Type = "VIDEO"
				}else if contents[j].Type == contentModel.PICTURE{
					postDTO.Type = "PICTURE"
				}
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
			if tags[p].PostId == posts[i].ID {
				listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].PostTagId))
			}
		}

		postDTO.Tags = listOfTags

		listOfPostsDTOs = append(listOfPostsDTOs, postDTO)

	}

	return listOfPostsDTOs

}

func (handler *SinglePostHandler) CreatePostDTO(posts *model.SinglePost, contents []contentModel.SinglePostContent, locations []locationModel.Location, tags []tagsModel.PostTagPosts) dto.SelectedPostDTO {


	var postDTO dto.SelectedPostDTO
	fmt.Println("POSTS")
	postDTO.PostId = posts.ID
	postDTO.Description = posts.Description
	postDTO.CreationDate = posts.CreationDate

	for j := 0; j < len(contents); j++ {
		if contents[j].SinglePostId == posts.ID {
			postDTO.Path = contents[j].Path

			if contents[j].Type == contentModel.VIDEO{
				postDTO.Type = "VIDEO"
			}else if contents[j].Type == contentModel.PICTURE{
				postDTO.Type = "PICTURE"
			}
		}
	}

	for k := 0; k < len(locations); k++ {
		if locations[k].ID == posts.LocationId {
			postDTO.LocationId = locations[k].ID
			postDTO.City = locations[k].City
			postDTO.Country = locations[k].Country
			postDTO.StreetName = locations[k].StreetName
			postDTO.StreetNumber = locations[k].StreetNumber
		}
	}

	var listOfTags []string
	for p := 0; p < len(tags); p++ {
		if tags[p].PostId == posts.ID {
			listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].PostTagId))

		}
	}

	postDTO.Tags = listOfTags


	return postDTO

}

// SEARCH TAGS FOR NOT REGISTERED USER
func (handler *SinglePostHandler) FindAllTagsForPublicPosts(w http.ResponseWriter, r *http.Request) {

	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidPosts = handler.SinglePostService.FindAllPublicPostsNotRegisteredUser(allPublicUsers)

	var tags = handler.PostTagPostsService.FindAllTagsForPosts(publicValidPosts)

	tagsJson, _ := json.Marshal(tags)
	w.Write(tagsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}