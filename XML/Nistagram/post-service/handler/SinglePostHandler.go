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
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
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
		return
	}

	err = handler.PostService.CreatePost(&singlePost.Post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
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
		return
	}

	fmt.Println("User IS valid")
	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	if profileSettings.UserVisibility == settingsModel.PRIVATE_VISIBILITY{
		fmt.Println("User IS PRIVATE")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//finds all posts
	var posts = handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id))
	var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

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
		return
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	if profileSettings.UserVisibility == settingsModel.PRIVATE_VISIBILITY{
		fmt.Println("User IS PRIVATE")

		// CHECK IF LOGID FOLLOWING POST USERID
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		if checkIfFollowing == true{
			var posts = handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id))
			var contents = handler.PostContentService.FindAllContentsForPosts(posts)
			var locations = handler.LocationService.FindAllLocationsForPosts(posts)
			var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

			var postsDTOS = handler.CreatePostsDTOList(posts,contents,locations,tags)

			postsJson, _ := json.Marshal(postsDTOS)
			w.Write(postsJson)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

		}else{

			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}else{
		var posts = handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id))
		//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

		var contents = handler.PostContentService.FindAllContentsForPosts(posts)
		var locations = handler.LocationService.FindAllLocationsForPosts(posts)
		var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

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

	// returns only valid users
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	for i := 0; i < len(allValidUsers); i++ {
		fmt.Println("FindAllFollowingPosts FindAllUsersButLoggedIn allValidUsers handler "+allValidUsers[i].Username)
	}


	// retuns only valid FOLLOWINGS
	var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)
	for i := 0; i < len(followings); i++ {
		fmt.Println("FindAllFollowingPosts FindAllUsersButLoggedIn followings handler "+allValidUsers[i].Username)
	}


	// returns NOT DELETED POSTS from valid following users
	var posts = handler.SinglePostService.FindAllFollowingPosts(followings)
	var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

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
		return
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	if profileSettings.UserVisibility == settingsModel.PUBLIC_VISIBILITY{
		// EVERYONE CAN SELECT THIS POST

		var contents = handler.PostContentService.FindAllContentsForPost(post)
		var locations = handler.LocationService.FindAllLocationsForPost(post)
		var tags = handler.PostTagPostsService.FindAllTagsForPost(post)
		var postDTO = handler.CreatePostDTO(post,contents,locations,tags)

		postJson, _ := json.Marshal(postDTO)
		w.Write(postJson)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		// FOR POSTMAN CHECK (should redirect)
		fmt.Println("Profile is private")
		w.WriteHeader(http.StatusExpectationFailed)
		return
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
		return
	}

	var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	if profileSettings.UserVisibility == settingsModel.PUBLIC_VISIBILITY{
		// EVERYONE CAN SELECT THIS POST

		var contents = handler.PostContentService.FindAllContentsForPost(post)
		var locations = handler.LocationService.FindAllLocationsForPost(post)
		var tags = handler.PostTagPostsService.FindAllTagsForPost(post)

		var postDTO = handler.CreatePostDTO(post,contents,locations,tags)

		postJson, _ := json.Marshal(postDTO)
		w.Write(postJson)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		// CHECK IF LOGID FOLLOWING POST USERID
		var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), post.UserID)
		if checkIfFollowing == true{
			var contents = handler.PostContentService.FindAllContentsForPost(post)
			var locations = handler.LocationService.FindAllLocationsForPost(post)
			var tags = handler.PostTagPostsService.FindAllTagsForPost(post)
			var postDTO = handler.CreatePostDTO(post,contents,locations,tags)

			postJson, _ := json.Marshal(postDTO)
			w.Write(postJson)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		}else{
			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}

}

//FIND ALL PUBLIC POSTS (for not registered users)
func (handler *SinglePostHandler) FindAllPublicPostsNotRegisteredUser(w http.ResponseWriter, r *http.Request) {

	// returns only VALID users
	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidPosts = handler.SinglePostService.FindAllPublicAndFriendsPostsValid(allPublicUsers)
	var contents = handler.PostContentService.FindAllContentsForPosts(publicValidPosts)
	var locations = handler.LocationService.FindAllLocationsForPosts(publicValidPosts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(publicValidPosts)
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

	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidPosts = handler.SinglePostService.FindAllPublicAndFriendsPostsValid(allPublicUsers)
	var contents = handler.PostContentService.FindAllContentsForPosts(publicValidPosts)
	var locations = handler.LocationService.FindAllLocationsForPosts(publicValidPosts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(publicValidPosts)

	var postsDTOS = handler.CreatePostsDTOList(publicValidPosts,contents,locations,tags)

	postJson, _ := json.Marshal(postsDTOS)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// all posts (EXCEPT DELETED) for my current logged in user
func (handler *SinglePostHandler) FindAllPostsForLoggedUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var posts = handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id))
	var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

	var postsDTOS = handler.CreatePostsDTOList(posts,contents,locations,tags)

	postsJson, _ := json.Marshal(postsDTOS)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}


// FIND SELECTED POST FROM LOGGEDIN USER BY ID (ONLY IF NOT DELETED)
func (handler *SinglePostHandler) FindSelectedPostByIdForLoggedUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id") //post id
	logId := r.URL.Query().Get("logId") //loged user id

	var post = handler.SinglePostService.FindByID(uuid.MustParse(id))
	if post == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if post.IsDeleted == true{
		fmt.Println("Deleted post")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if post.UserID != uuid.MustParse(logId){
			//POSTMAN CHECK
			fmt.Println("Post doesnt belong to user")
			w.WriteHeader(http.StatusExpectationFailed)
			return
	}

	var contents = handler.PostContentService.FindAllContentsForPost(post)
	var locations = handler.LocationService.FindAllLocationsForPost(post)
	var tags = handler.PostTagPostsService.FindAllTagsForPost(post)

	var postDTO = handler.CreatePostDTO(post,contents,locations,tags)

	postJson, _ := json.Marshal(postDTO)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}



//DTOS

func (handler *SinglePostHandler) CreatePostsDTOList(posts []model.SinglePost, contents []contentModel.SinglePostContent, locations []locationModel.Location, tags []tagsModel.PostTagPosts) []dto.SelectedPostDTO {
	var listOfPostsDTOs []dto.SelectedPostDTO

	for i := 0; i < len(posts); i++ {
		var postDTO dto.SelectedPostDTO
		postDTO.PostId = posts[i].ID
		postDTO.Description = posts[i].Description
		postDTO.CreationDate = posts[i].CreationDate
		postDTO.UserId = posts[i].UserID

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
				listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))
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
			listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))

		}
	}

	postDTO.Tags = listOfTags


	return postDTO

}

// NOT REGISTERED

// SEARCH TAGS FOR NOT REGISTERED USER
func (handler *SinglePostHandler) FindAllTagsForPublicPosts(w http.ResponseWriter, r *http.Request) {

	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidPosts = handler.SinglePostService.FindAllPublicAndFriendsPostsValid(allPublicUsers)

	var tags = handler.PostTagPostsService.FindAllTagsForPosts(publicValidPosts)

	tagsJson, _ := json.Marshal(tags)
	w.Write(tagsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// SEARCH LOCATIONS FOR NOT REGISTERED USER
func (handler *SinglePostHandler) FindAllLocationsForPublicPosts(w http.ResponseWriter, r *http.Request) {

	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidPosts = handler.SinglePostService.FindAllPublicAndFriendsPostsValid(allPublicUsers)
	var locations = handler.LocationService.FindAllLocationsForPosts(publicValidPosts)

	locationsJson, _ := json.Marshal(locations)
	w.Write(locationsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// FIND ALL PUBLIC NOT DELETED POSTS WITH TAG - FOR NOT REG USER S
func (handler *SinglePostHandler) FindAllPostsForTag(w http.ResponseWriter, r *http.Request) {

	tagName := r.URL.Query().Get("tagName") //tag id


	var tag = handler.TagService.FindTagByName(tagName)
	var postIds = handler.PostTagPostsService.FindAllPostIdsWithTagId(tag.ID)
	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var posts = handler.SinglePostService.FindAllPublicPostsByIds(postIds, allValidUsers)

	var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

	var postDTO = handler.CreatePostsDTOList(posts,contents,locations,tags)

	postsJson, _ := json.Marshal(postDTO)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// FIND ALL PUBLIC NOT DELETED POSTS WITH LOCATION - FOR NOT REG USER S
func (handler *SinglePostHandler) FindAllPostsForLocation(w http.ResponseWriter, r *http.Request) {

	//county,city,streetName,streetNumber
	locationString := r.URL.Query().Get("locationString")

	var location = handler.LocationService.FindLocationIdByLocationString(locationString)
	var locationPosts = handler.SinglePostService.FindAllPostIdsWithLocationId(location.ID)
	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var posts = handler.SinglePostService.FindAllPublicAndFriendsPosts(locationPosts, allValidUsers)


	var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

	var postDTO = handler.CreatePostsDTOList(posts,contents,locations,tags)

	postsJson, _ := json.Marshal(postDTO)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


// REGISTERED

func Find(slice []userModel.ClassicUser, val userModel.ClassicUser) (int,bool){
	for i, item := range slice{
		if item.ID == val.ID{
			fmt.Println("Pronasao ga u Find")
			return i, true
		}
	}

	return -1, false
}

func (handler *SinglePostHandler) MergePublicAndFollowingUsers(allPublicUsers []userModel.ClassicUser, allFollowingUsers []userModel.ClassicUser) []userModel.ClassicUser {
	//merge public and following users
	var allPublicAndFollowingUsers []userModel.ClassicUser
	for i := 0; i < len(allPublicUsers); i++ {
		fmt.Println(allPublicUsers[i].Username)
		allPublicAndFollowingUsers = append(allPublicAndFollowingUsers, allPublicUsers[i])
	}
	for i := 0; i < len(allFollowingUsers); i++ {
		_, found := Find(allPublicAndFollowingUsers, allFollowingUsers[i])

		if !found {
			allPublicAndFollowingUsers = append(allPublicAndFollowingUsers, allFollowingUsers[i])
		}
	}
	for i := 0; i < len(allPublicAndFollowingUsers); i++ {
		fmt.Println(allPublicAndFollowingUsers[i].Username)
	}
	fmt.Println()
	return allPublicAndFollowingUsers
}

func (handler *SinglePostHandler) FindAllPublicAndFriendsUsers(id uuid.UUID) []userModel.ClassicUser {

	var allValidUsers = handler.ClassicUserService.FinAllValidUsers() //ok
	var allValidUsersButLoggedIn = handler.FindAllValidUsersButLoggedIn(id, allValidUsers)//ok


	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsersButLoggedIn)//ok

	var allFollowings = handler.ClassicUserFollowingsService.FindAllUserWhoFollowUserId(id, allValidUsersButLoggedIn) //moj user je classic user
	var allFollowingUsers = handler.ClassicUserService.FindAllUsersByFollowingIds(allFollowings)

	// ALL PUBLIC AND FRIENDS USERS EXCEPT LOGGED
	var allUsers = handler.MergePublicAndFollowingUsers(allPublicUsers, allFollowingUsers)
	fmt.Println("Duzina liste")
	fmt.Println(len(allUsers))

	return allUsers
}

func (handler *SinglePostHandler) FindAllValidUsersButLoggedIn(id uuid.UUID, allValidUsers []userModel.ClassicUser) []userModel.ClassicUser {
	var allUsers []userModel.ClassicUser
	myUser := handler.FindMyUserById(id, allValidUsers)

	for i := 0; i < len(allValidUsers); i++ {
		found:= myUser.ID == allValidUsers[i].ID
		if !found {
			fmt.Println(allValidUsers[i].ID.String()+" FindAllValidUsersButLoggedIn")
			allUsers = append(allUsers, allValidUsers[i])
		}
	}

	return allUsers
}

func (handler *SinglePostHandler) FindMyUserById(id uuid.UUID, allValidUsers []userModel.ClassicUser) userModel.ClassicUser {
	//var allUsers []userModel.ClassicUser
	var myUser userModel.ClassicUser
	for i := 0; i < len(allValidUsers); i++ {
		if allValidUsers[i].ID == id {
			myUser = allValidUsers[i]
			return myUser
		}
	}
	return myUser
}

// SEARCH TAGS FOR REGISTERED USER - FIND ALL TAGS ON PUBLIC AND FOLLOWING POSTS
func (handler *SinglePostHandler) FindAllTagsForPublicAndFollowingPosts(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id") //logged in reg user id

	var allUsers = handler.FindAllPublicAndFriendsUsers(uuid.MustParse(id))
	var allValidUsersButLoggedIn = handler.FindAllValidUsersButLoggedIn(uuid.MustParse(id), allUsers)
	var allPosts = handler.SinglePostService.FindAllPostsForUsers(allValidUsersButLoggedIn)
	var tags = handler.PostTagPostsService.FindAllTagsForPosts(allPosts)

	tagsJson, _ := json.Marshal(tags)
	w.Write(tagsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}



// SEARCH LOCATIONS FOR REGISTERED USER - FIND ALL LOCATIONS ON PUBLIC AND FOLLOWING POSTS
func (handler *SinglePostHandler) FindAllLocationsForPublicAndFollowingPosts(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id") //logged in reg user id
	fmt.Println(id)
	var allUsers = handler.FindAllPublicAndFriendsUsers(uuid.MustParse(id))

	var allValidUsersButLoggedIn = handler.FindAllValidUsersButLoggedIn(uuid.MustParse(id), allUsers)
	var allPosts = handler.SinglePostService.FindAllPostsForUsers(allValidUsersButLoggedIn)

	var locations = handler.LocationService.FindAllLocationsForPosts(allPosts)

	locationsJson, _ := json.Marshal(locations)
	w.Write(locationsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// FIND ALL PUBLIC OR FOLLOWING AND NOT DELETED POSTS WITH TAG - FOR REG USER S
func (handler *SinglePostHandler) FindAllPostsForTagRegUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id") //logged in reg user id
	tagName := r.URL.Query().Get("tagName") //tag id


	var tag = handler.TagService.FindTagByName(tagName)
	var postIds = handler.PostTagPostsService.FindAllPostIdsWithTagId(tag.ID)

	var allUsers = handler.FindAllPublicAndFriendsUsers(uuid.MustParse(id))
	var allValidUsersButLoggedIn = handler.FindAllValidUsersButLoggedIn(uuid.MustParse(id), allUsers)
	var posts = handler.SinglePostService.FindAllPublicPostsByIds(postIds, allValidUsersButLoggedIn)

	var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

	var postDTO = handler.CreatePostsDTOList(posts,contents,locations,tags)

	postsJson, _ := json.Marshal(postDTO)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// FIND ALL PUBLIC OR FOLLOWING NOT DELETED POSTS WITH LOCATION - FOR REG USER S
func (handler *SinglePostHandler) FindAllPostsForLocationRegUser(w http.ResponseWriter, r *http.Request) {


	//county,city,streetName,streetNumber
	locationString := r.URL.Query().Get("locationString")
	id := r.URL.Query().Get("id") //logged in reg user id

	var location = handler.LocationService.FindLocationIdByLocationString(locationString)
	var locationPosts = handler.SinglePostService.FindAllPostIdsWithLocationId(location.ID)

	var allUsers = handler.FindAllPublicAndFriendsUsers(uuid.MustParse(id))
	var allValidUsersButLoggedIn = handler.FindAllValidUsersButLoggedIn(uuid.MustParse(id), allUsers)
	var posts = handler.SinglePostService.FindAllPublicAndFriendsPosts(locationPosts, allValidUsersButLoggedIn)


	var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts)

	var postDTO = handler.CreatePostsDTOList(posts,contents,locations,tags)

	postsJson, _ := json.Marshal(postDTO)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
