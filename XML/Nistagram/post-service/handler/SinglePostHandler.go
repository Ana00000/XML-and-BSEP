package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"os"
	"time"
)

type SinglePostHandler struct {
	SinglePostService * service.SinglePostService
	PostService * service.PostService
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

type UserValid struct {
	IsValid bool `json:"is_valid"`
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// for selected user (you can only select VALID users)
func (handler *SinglePostHandler) FindAllPostsForUserNotRegisteredUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	//var checkIfValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(id))
	var userValidity UserValid
	reqUrl := fmt.Sprintf("http://%s:%s/check_if_user_valid/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	err := getJson(reqUrl, &userValidity)
	if err!=nil{
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	var checkIfValid=userValidity.IsValid
	if  checkIfValid == false {
		fmt.Println("User NOT valid")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	fmt.Println("User IS valid")
	//var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	var profileSettings dto.ProfileSettingsDTO
	reqUrl = fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err = getJson(reqUrl, &profileSettings)
	if err!=nil{
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if profileSettings.UserVisibility == "PRIVATE_VISIBILITY"{
		fmt.Println("User IS PRIVATE")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	//finds all posts
	var posts = convertListSinglePostsToSinglePostsDTO(handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id)))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	//finds all conents
	//var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(posts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SinglePostContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//finds all locations
	//var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(posts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	//find all tags
	//var tags = handler.PostTagPostsService.FindAllTagsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(posts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostTagPostsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(posts),contents,locations,tags)


	postsJson, _ := json.Marshal(postsDTOS)
	w.Write(postsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")


}

type ReturnValueBool struct {
	ReturnValue bool `json:"return_value"`
}

func (handler *SinglePostHandler) FindAllPostsForUserRegisteredUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	//var checkIfValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(id))
	var userValidity UserValid
	reqUrl := fmt.Sprintf("http://%s:%s/check_if_user_valid/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	err := getJson(reqUrl, &userValidity)
	if err!=nil{
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	var checkIfValid=userValidity.IsValid

	if  checkIfValid == false {
		fmt.Println("User NOT valid")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	//var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(uuid.MustParse(id))
	var profileSettings dto.ProfileSettingsDTO
	reqUrl = fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err = getJson(reqUrl, &profileSettings)
	if err!=nil{
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if profileSettings.UserVisibility == "PRIVATE_VISIBILITY"{
		fmt.Println("User IS PRIVATE")

		// CHECK IF LOGID FOLLOWING POST USERID
		//var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(logId), uuid.MustParse(id))
		var returnValueFollowing ReturnValueBool
		reqUrl = fmt.Sprintf("http://%s:%s/check_if_following_post_story/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id, logId)
		err = getJson(reqUrl, &returnValueFollowing)
		if err!=nil{
			fmt.Println("Wrong cast response body to ProfileSettingDTO!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		checkIfFollowing := returnValueFollowing.ReturnValue
		if checkIfFollowing == true{
			var posts = convertListSinglePostsToSinglePostsDTO(handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id)))
			//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

			//finds all conents
			//var contents = handler.PostContentService.FindAllContentsForPosts(posts)
			reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
			jsonValidPostsDTO, _ := json.Marshal(posts)
			fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			fmt.Println(string(jsonValidPostsDTO))
			resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
			if err != nil || resp.StatusCode == 400 {
				print("Fail")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var contents []dto.SinglePostContentDTO
			if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
				w.WriteHeader(http.StatusConflict) //400
				return
			}

			//finds all locations
			//var locations = handler.LocationService.FindAllLocationsForPosts(posts)
			reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
			jsonLocationsDTO, _ := json.Marshal(posts)
			fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			fmt.Println(string(jsonLocationsDTO))
			resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
			if err != nil || resp.StatusCode == 400 {
				print("Fail")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var locations []dto.LocationDTO
			if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
				w.WriteHeader(http.StatusConflict) //400
				return
			}

			//find all tags
			//var tags = handler.PostTagPostsService.FindAllTagsForPosts(posts)
			reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
			jsonTagsDTO, _ := json.Marshal(posts)
			fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			fmt.Println(string(jsonTagsDTO))
			resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
			if err != nil || resp.StatusCode == 400 {
				print("Fail")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var tags []dto.PostTagPostsDTO
			if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
				w.WriteHeader(http.StatusConflict) //400
				return
			}

			//creates a list of dtos
			var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(posts),contents,locations,tags)

			postsJson, _ := json.Marshal(postsDTOS)
			w.Write(postsJson)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

		}else{

			fmt.Println("Not following private user")
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}else{
		var posts = convertListSinglePostsToSinglePostsDTO(handler.SinglePostService.FindAllPostsForUser(uuid.MustParse(id)))
		//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

		//finds all conents
		//var contents = handler.PostContentService.FindAllContentsForPosts(posts)
		reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
		jsonValidPostsDTO, _ := json.Marshal(posts)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonValidPostsDTO))
		resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var contents []dto.SinglePostContentDTO
		if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}

		//finds all locations
		//var locations = handler.LocationService.FindAllLocationsForPosts(posts)
		reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
		jsonLocationsDTO, _ := json.Marshal(posts)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonLocationsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var locations []dto.LocationDTO
		if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}
		//find all tags
		//var tags = handler.PostTagPostsService.FindAllTagsForPosts(posts)
		reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
		jsonTagsDTO, _ := json.Marshal(posts)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonTagsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var tags []dto.PostTagPostsDTO
		if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}
		//creates a list of dtos
		var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(posts),contents,locations,tags)

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
	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"),id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	// retuns only valid FOLLOWINGS
	//var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_valid_followings_for_user/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), id)
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var followings []dto.ClassicUserFollowingsFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&followings); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	// returns NOT DELETED POSTS from valid following users
	var posts = convertListSinglePostsToSinglePostsDTO(handler.SinglePostService.FindAllFollowingPosts(followings))

	//finds all conents
	//var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(posts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SinglePostContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//finds all locations
	//var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(posts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//find all tags
	//var tags = handler.PostTagPostsService.FindAllTagsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(posts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostTagPostsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(posts),contents,locations,tags)

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

	//var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	var profileSettings dto.ProfileSettingsDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err := getJson(reqUrl, &profileSettings)
	if err!=nil{
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if profileSettings.UserVisibility == "PUBLIC_VISIBILITY"{
		// EVERYONE CAN SELECT THIS POST
		//finds all conents
		/*
		var contents = handler.PostContentService.FindAllContentsForPost(post)

		//finds all locations
		var locations = handler.LocationService.FindAllLocationsForPost(post)

		//find all tags
		var tags = handler.PostTagPostsService.FindAllTagsForPost(post)
		*/

		reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_post/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
		jsonValidStoriesDTO, _ := json.Marshal(post)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonValidStoriesDTO))
		resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var contents []dto.SinglePostContentDTO
		if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}


		//var locations = handler.LocationService.FindAllLocationsForStories(stories)
		reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
		jsonLocationsDTO, _ := json.Marshal(post)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonLocationsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var locations []dto.LocationDTO
		if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}


		//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
		reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
		jsonTagsDTO, _ := json.Marshal(post)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonTagsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var tags []dto.PostTagPostsDTO
		if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}

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

	//var profileSettings = handler.ProfileSettings.FindProfileSettingByUserId(post.UserID)
	var profileSettings dto.ProfileSettingsDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err := getJson(reqUrl, &profileSettings)
	if err!=nil{
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if profileSettings.UserVisibility == "PUBLIC_VISIBILITY"{
		// EVERYONE CAN SELECT THIS POST
		//finds all conents
		/*var contents = handler.PostContentService.FindAllContentsForPost(post)

		//finds all locations
		var locations = handler.LocationService.FindAllLocationsForPost(post)

		//find all tags
		var tags = handler.PostTagPostsService.FindAllTagsForPost(post)*/
		reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_post/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
		jsonValidStoriesDTO, _ := json.Marshal(post)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonValidStoriesDTO))
		resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var contents []dto.SinglePostContentDTO
		if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}


		//var locations = handler.LocationService.FindAllLocationsForStories(stories)
		reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
		jsonLocationsDTO, _ := json.Marshal(post)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonLocationsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var locations []dto.LocationDTO
		if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}


		//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
		reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
		jsonTagsDTO, _ := json.Marshal(post)
		fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
		fmt.Println(string(jsonTagsDTO))
		resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
		if err != nil || resp.StatusCode == 400 {
			print("Fail")
			w.WriteHeader(http.StatusFailedDependency)
			return
		}
		//defer resp.Body.Close() mozda treba dodati
		var tags []dto.PostTagPostsDTO
		if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
			w.WriteHeader(http.StatusConflict) //400
			return
		}
		//creates a list of dtos
		var postDTO = handler.CreatePostDTO(post,contents,locations,tags)

		postJson, _ := json.Marshal(postDTO)
		w.Write(postJson)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		// CHECK IF LOGID FOLLOWING POST USERID
		//var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(), post.UserID)
		var returnValueFollowing ReturnValueBool
		reqUrl = fmt.Sprintf("http://%s:%s/check_if_following_post_story/%s/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), logId, post.UserID.String())
		err = getJson(reqUrl, &returnValueFollowing)
		if err!=nil{
			fmt.Println("Wrong cast response body to ProfileSettingDTO!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		checkIfFollowing := returnValueFollowing.ReturnValue

		if checkIfFollowing == true{

			//finds all conents
			/*
			var contents = handler.PostContentService.FindAllContentsForPost(post)

			//finds all locations
			var locations = handler.LocationService.FindAllLocationsForPost(post)

			//find all tags
			var tags = handler.PostTagPostsService.FindAllTagsForPost(post)
			*/
			reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_post/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
			jsonValidStoriesDTO, _ := json.Marshal(post)
			fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			fmt.Println(string(jsonValidStoriesDTO))
			resp,err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidStoriesDTO))
			if err != nil || resp.StatusCode == 400 {
				print("Fail")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var contents []dto.SinglePostContentDTO
			if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
				w.WriteHeader(http.StatusConflict) //400
				return
			}


			//var locations = handler.LocationService.FindAllLocationsForStories(stories)
			reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
			jsonLocationsDTO, _ := json.Marshal(post)
			fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			fmt.Println(string(jsonLocationsDTO))
			resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
			if err != nil || resp.StatusCode == 400 {
				print("Fail")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var locations []dto.LocationDTO
			if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
				w.WriteHeader(http.StatusConflict) //400
				return
			}


			//var tags = handler.StoryTagStoriesService.FindAllTagsForStories(stories)
			reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
			jsonTagsDTO, _ := json.Marshal(post)
			fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
			fmt.Println(string(jsonTagsDTO))
			resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
			if err != nil || resp.StatusCode == 400 {
				print("Fail")
				w.WriteHeader(http.StatusFailedDependency)
				return
			}
			//defer resp.Body.Close() mozda treba dodati
			var tags []dto.PostTagPostsDTO
			if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
				w.WriteHeader(http.StatusConflict) //400
				return
			}
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

	// returns only VALID users
	//var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_valid_users/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	// returns all PUBLIC users
	//var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var allPublicUsers []dto.ClassicUserDTO
	if err := json.NewDecoder(resp.Body).Decode(&allPublicUsers); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	// returns all POSTS of public and valid users
	var publicValidPosts = convertListSinglePostsToSinglePostsDTO(handler.SinglePostService.FindAllPublicPostsNotRegisteredUser(allPublicUsers))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST

	//finds all conents
	//var contents = handler.PostContentService.FindAllContentsForPosts(publicValidPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(publicValidPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SinglePostContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//finds all locations
	//var locations = handler.LocationService.FindAllLocationsForPosts(publicValidPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//find all tags
	//var tags = handler.PostTagPostsService.FindAllTagsForPosts(publicValidPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostTagPostsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(publicValidPosts),contents,locations,tags)

	postJson, _ := json.Marshal(postsDTOS)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


//FindAllPublicPostsRegisteredUser
func (handler *SinglePostHandler) FindAllPublicPostsRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	// returns only VALID users but loggedIn user
	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"),id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	// returns all PUBLIC users
	//var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var allPublicUsers []dto.ClassicUserDTO
	if err := json.NewDecoder(resp.Body).Decode(&allPublicUsers); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}


	// returns all POSTS of public and valid users
	var publicValidPosts = convertListSinglePostsToSinglePostsDTO(handler.SinglePostService.FindAllPublicPostsNotRegisteredUser(allPublicUsers))


	//finds all conents
	//var contents = handler.PostContentService.FindAllContentsForPosts(publicValidPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(publicValidPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SinglePostContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//finds all locations
	//var locations = handler.LocationService.FindAllLocationsForPosts(publicValidPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//find all tags
	//var tags = handler.PostTagPostsService.FindAllTagsForPosts(publicValidPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostTagPostsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(publicValidPosts),contents,locations,tags)

	postJson, _ := json.Marshal(postsDTOS)
	w.Write(postJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostHandler) CreatePostsDTOList(posts []model.SinglePost, contents []dto.SinglePostContentDTO, locations []dto.LocationDTO, tags []dto.PostTagPostsDTO) []dto.SelectedPostDTO {
	var listOfPostsDTOs []dto.SelectedPostDTO

	for i := 0; i < len(posts); i++ {
		var postDTO dto.SelectedPostDTO
		postDTO.PostId = posts[i].ID
		postDTO.Description = posts[i].Description
		postDTO.CreationDate = posts[i].CreationDate

		for j := 0; j < len(contents); j++ {
			if contents[j].SinglePostId == posts[i].ID {
				postDTO.Path = contents[j].Path
				postDTO.Type = contents[j].Type
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
				var returnValueTagName ReturnValueString
				reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"),tags[p].PostTagId.String())
				err := getJson(reqUrl, &returnValueTagName)
				if err!=nil{
					fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
					panic(err)
				}
				listOfTags = append(listOfTags, returnValueTagName.ReturnValue)
			}
		}

		postDTO.Tags = listOfTags

		listOfPostsDTOs = append(listOfPostsDTOs, postDTO)

	}

	return listOfPostsDTOs

}

type ReturnValueString struct {
	ReturnValue string `json:"return_value"`
}

func (handler *SinglePostHandler) CreatePostDTO(posts *model.SinglePost, contents []dto.SinglePostContentDTO, locations []dto.LocationDTO, tags []dto.PostTagPostsDTO) dto.SelectedPostDTO {


	var postDTO dto.SelectedPostDTO
	fmt.Println("POSTS")
	postDTO.PostId = posts.ID
	postDTO.Description = posts.Description
	postDTO.CreationDate = posts.CreationDate

	for j := 0; j < len(contents); j++ {
		if contents[j].SinglePostId == posts.ID {
			postDTO.Path = contents[j].Path
			postDTO.Type = contents[j].Type
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
			var returnValueTagName  ReturnValueString
			reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"),tags[p].PostTagId.String())
			err := getJson(reqUrl, &returnValueTagName)
			if err!=nil{
				fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
				panic(err)
			}
			//handler.TagService.FindTagNameById(tags[p].PostTagId)
			listOfTags = append(listOfTags, returnValueTagName.ReturnValue)

		}
	}

	postDTO.Tags = listOfTags


	return postDTO

}

func convertListSinglePostsToSinglePostsDTO(singlePosts []model.SinglePost) []dto.SinglePostFullDTO{
	var singlePostsDTO []dto.SinglePostFullDTO
	for i := 0; i < len(singlePosts); i++ {
		singlePostsDTO=append(singlePostsDTO,convertSinglePostToSinglePostDTO(singlePosts[i]))
	}
	return singlePostsDTO
}

func convertSinglePostToSinglePostDTO(singlePost model.SinglePost) dto.SinglePostFullDTO{
	layout := "2006-01-02T15:04:05.000Z"
	var singlePostDTO= dto.SinglePostFullDTO{
		ID:           singlePost.ID,
		Description:  singlePost.Description,
		CreationDate: singlePost.CreationDate.Format(layout),
		UserID:       singlePost.UserID,
		LocationId:   singlePost.LocationId,
		IsDeleted:    singlePost.IsDeleted,
	}
	return singlePostDTO
}

func convertSinglePostsDTOToListSinglePosts(singlePostsDTO []dto.SinglePostFullDTO) []model.SinglePost{
	var singlePosts []model.SinglePost
	for i := 0; i < len(singlePostsDTO); i++ {
		singlePosts=append(singlePosts,convertSinglePostDTOToSinglePost(singlePostsDTO[i]))
	}
	return singlePosts
}

func convertSinglePostDTOToSinglePost(singlePostDTO dto.SinglePostFullDTO) model.SinglePost{
	layout := "2006-01-02T15:04:05.000Z"
	creationDate,_ := time.Parse(layout,singlePostDTO.CreationDate)
	var singlePost= model.SinglePost{
		Post: model.Post{
			ID:           singlePostDTO.ID,
			Description:  singlePostDTO.Description,
			CreationDate: creationDate,
			UserID:       singlePostDTO.UserID,
			LocationId:   singlePostDTO.LocationId,
			IsDeleted:    singlePostDTO.IsDeleted,
		},
	}
	return singlePost
}