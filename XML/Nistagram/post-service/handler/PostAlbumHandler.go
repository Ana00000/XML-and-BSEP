package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"os"
	"time"
)

type PostAlbumHandler struct {
	Service * service.PostAlbumService
	PostService * service.PostService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *PostAlbumHandler) CreatePostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumDTO dto.PostAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "CRPAL580",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postAlbum := model.PostAlbum{
		Post : model.Post{
			ID: id,
			Description: postAlbumDTO.Description,
			CreationDate: time.Now(),
			UserID: postAlbumDTO.UserID,
			LocationId: postAlbumDTO.LocationID,
			IsDeleted: false,
		},
	}

	err = handler.Service.CreatePostAlbum(&postAlbum)
	if err != nil {
		fmt.Println(err)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "CRPAL580",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating post album!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.PostService.CreatePost(&postAlbum.Post)
	if err != nil {
		fmt.Println(err)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "CRPAL580",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating post!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	postAlbumIDJson, _ := json.Marshal(postAlbum.ID)
	w.Write(postAlbumIDJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumHandler",
		"action":   "CRPAL580",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created post album!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) FindAllAlbumPostsForLoggedUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var albumPosts = convertListPostAlbumToListPostAlbumDTO(handler.Service.FindAllAlbumPostsForUser(uuid.MustParse(id)))
	///find_all_contents_for_post_albums/
	//var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(albumPosts)
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_post_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(albumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAAPL581",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all contents for post albums!")
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.PostAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAAPL581",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	///find_locations_for_post_albums/
	//var locations = handler.LocationService.FindAllLocationsForPostAlbums(albumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(albumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAAPL581",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all locations for post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAAPL581",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	///find_all_tags_for_post_album_tag_post_albums/
	//var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(albumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post_album_tag_post_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(albumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAAPL581",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all tags for post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostAlbumTagPostAlbumsFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAAPL581",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumTagPostAlbumsFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var albumPostsDTOS = handler.CreatePostAlbumsDTOList(convertListPostAlbumDTOToListPostAlbum(albumPosts),contents,locations,tags)

	albumPostsJson, _ := json.Marshal(albumPostsDTOS)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumHandler",
		"action":   "FAAPL581",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found all post albums for logged user!")
	w.Write(albumPostsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) CreatePostAlbumsDTOList(albums []model.PostAlbum, contents []dto.PostAlbumContentFullDTO, locations []dto.LocationDTO, tags []dto.PostAlbumTagPostAlbumsFullDTO) []dto.SelectedPostAlbumDTO {
	var listOfPostAlbumsDTOs []dto.SelectedPostAlbumDTO

	for i := 0; i < len(albums); i++ {
		var postAlbumDTO dto.SelectedPostAlbumDTO
		postAlbumDTO.PostAlbumId = albums[i].ID
		postAlbumDTO.Description = albums[i].Description
		postAlbumDTO.CreationDate = albums[i].CreationDate
		postAlbumDTO.UserId = albums[i].UserID

		for j := 0; j < len(contents); j++ {
			if contents[j].PostAlbumId == albums[i].ID {
				postAlbumDTO.Path = append(postAlbumDTO.Path, contents[j].Path)
				postAlbumDTO.Type = append(postAlbumDTO.Type, contents[j].Type)
			}
		}

		for k := 0; k < len(locations); k++ {
			if locations[k].ID == albums[i].LocationId {
				postAlbumDTO.LocationId = locations[k].ID
				postAlbumDTO.City = locations[k].City
				postAlbumDTO.Country = locations[k].Country
				postAlbumDTO.StreetName = locations[k].StreetName
				postAlbumDTO.StreetNumber = locations[k].StreetNumber
			}
		}

		var listOfTags []string
		for p := 0; p < len(tags); p++ {
			if tags[p].PostAlbumId == albums[i].ID {
				//listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))
				var returnValueTagName  ReturnValueString
				reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"),tags[p].TagId.String())
				err := getJson(reqUrl, &returnValueTagName)
				if err!=nil{
					fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
					panic(err)
				}
				listOfTags = append(listOfTags, returnValueTagName.ReturnValue)

			}
		}

		postAlbumDTO.Tags = listOfTags

		listOfPostAlbumsDTOs = append(listOfPostAlbumsDTOs, postAlbumDTO)

	}

	return listOfPostAlbumsDTOs

}

func (handler *PostAlbumHandler) FindSelectedPostAlbumByIdForLoggedUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id") //post album id
	logId := r.URL.Query().Get("logId") //loged user id

	var postAlbum = handler.Service.FindByID(uuid.MustParse(id))

	var postAlbumFullDTO = convertPostAlbumToPostAlbumDTO(*postAlbum)
	if postAlbum == nil {
		fmt.Println("User not found")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if postAlbum.IsDeleted == true{
		fmt.Println("Deleted post album")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Deleted post album!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if postAlbum.UserID != uuid.MustParse(logId){
		fmt.Println("Post album doesnt belong to user")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Post album doesnt belong to user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbum(postAlbum)
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_post_album/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(postAlbumFullDTO)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all contents for post album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.PostAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	///find_locations_for_post_album/
	//var locations = handler.LocationService.FindAllLocationsForPostAlbum(postAlbum)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post_album/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(postAlbumFullDTO)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all location for post album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//
	//var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbum(postAlbum)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post_album/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(postAlbumFullDTO)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all tags for post album!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostAlbumTagPostAlbumsFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FSPAL583",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumTagPostAlbumsFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var postAlbumDTO = handler.CreatePostAlbumDTO(postAlbum,contents,locations,tags)

	postAlbumJson, _ := json.Marshal(postAlbumDTO)
	w.Write(postAlbumJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumHandler",
		"action":   "FSPAL583",
		"timestamp":   time.Now().String(),
	}).Info("Successfuly found selected post album by id for logged user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *PostAlbumHandler) CreatePostAlbumDTO(album *model.PostAlbum, contents []dto.PostAlbumContentFullDTO, locations []dto.LocationDTO, tags []dto.PostAlbumTagPostAlbumsFullDTO) dto.SelectedPostAlbumDTO {
	var postAlbumDTO dto.SelectedPostAlbumDTO

	postAlbumDTO.PostAlbumId = album.ID
	postAlbumDTO.Description = album.Description
	postAlbumDTO.CreationDate = album.CreationDate


	for j := 0; j < len(contents); j++ {
		if contents[j].PostAlbumId == album.ID {
			postAlbumDTO.Path = append(postAlbumDTO.Path, contents[j].Path)
			postAlbumDTO.Type = append(postAlbumDTO.Type, contents[j].Type)
		}
	}

	for k := 0; k < len(locations); k++ {
		if locations[k].ID == album.LocationId {
			postAlbumDTO.LocationId = locations[k].ID
			postAlbumDTO.City = locations[k].City
			postAlbumDTO.Country = locations[k].Country
			postAlbumDTO.StreetName = locations[k].StreetName
			postAlbumDTO.StreetNumber = locations[k].StreetNumber
		}
	}

	var listOfTags []string
	for p := 0; p < len(tags); p++ {
		if tags[p].PostAlbumId == album.ID {
			var returnValueTagName  ReturnValueString
			reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"),tags[p].TagId.String())
			err := getJson(reqUrl, &returnValueTagName)
			if err!=nil{
				fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
				panic(err)

			}
			listOfTags = append(listOfTags, returnValueTagName.ReturnValue)
			//listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))

		}
	}

	postAlbumDTO.Tags = listOfTags

	return postAlbumDTO
}

func (handler *PostAlbumHandler) FindAllPublicAlbumPostsRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	// returns only VALID users but loggedIn user
	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"),id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		//fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all users but logged in!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all public users!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var allPublicUsers []dto.ClassicUserDTO
	if err := json.NewDecoder(resp.Body).Decode(&allPublicUsers); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ClassicUserDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var publicValidAlbumPosts = convertListPostAlbumToListPostAlbumDTO(handler.Service.FindAllPublicAndFriendsPostAlbumsValid(allPublicUsers))
	//var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(publicValidAlbumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_post_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(publicValidAlbumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all collections for post albums")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.PostAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var locations = handler.LocationService.FindAllLocationsForPostAlbums(publicValidAlbumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidAlbumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all locations for post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(publicValidAlbumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post_album_tag_post_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidAlbumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all tags for post album tag post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostAlbumTagPostAlbumsFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAPAP584",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumTagPostAlbumsFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var postAlbumsDTOS = handler.CreatePostAlbumsDTOList(convertListPostAlbumDTOToListPostAlbum(publicValidAlbumPosts),contents,locations,tags)

	postAlbumsJson, _ := json.Marshal(postAlbumsDTOS)
	w.Write(postAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumHandler",
		"action":   "FAPAP584",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found all public post albums for registered user!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) FindAllPublicAlbumPostsNotRegisteredUser(w http.ResponseWriter, r *http.Request) {

	//var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_valid_users/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		//fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all valid users!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	jsonClassicUsersDTO, _ := json.Marshal(allValidUsers)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all public users!")
		panic(err)
		return
	}
	var allPublicUsers []dto.ClassicUserDTO
	if err := json.NewDecoder(resp.Body).Decode(&allPublicUsers); err != nil {
		//w.WriteHeader(http.StatusConflict) //400
		panic(err)
	}

	var publicValidAlbumPosts = convertListPostAlbumToListPostAlbumDTO(handler.Service.FindAllPublicAndFriendsPostAlbumsValid(allPublicUsers))
	//var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(publicValidAlbumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_post_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(publicValidAlbumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all contents for post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.PostAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var locations = handler.LocationService.FindAllLocationsForPostAlbums(publicValidAlbumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(publicValidAlbumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all locations for post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(publicValidAlbumPosts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post_album_tag_post_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(publicValidAlbumPosts)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all tags for post album tag post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostAlbumTagPostAlbumsFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FPAPN585",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumTagPostAlbumsFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var postAlbumsDTOS = handler.CreatePostAlbumsDTOList(convertListPostAlbumDTOToListPostAlbum(publicValidAlbumPosts),contents,locations,tags)

	postAlbumsJson, _ := json.Marshal(postAlbumsDTOS)
	w.Write(postAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumHandler",
		"action":   "FPAPN585",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found all public post albums for not registered user! ")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) FindAllFollowingPostAlbums(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// returns only valid users
	//var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	var  allValidUsers []dto.ClassicUserDTO
	reqUrl := fmt.Sprintf("http://%s:%s/dto/find_all_classic_users_but_logged_in?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"),id)
	err := getJson(reqUrl, &allValidUsers)
	if err!=nil{
		//fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all users but logged in!")
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
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all valid followings for user!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var followings []dto.ClassicUserFollowingsFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&followings); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ClassicUserFollowingsFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	// returns NOT DELETED POST ALBUMS from valid following users
	var postAlbums = convertListPostAlbumToListPostAlbumDTO(handler.Service.FindAllFollowingPostAlbums(followings))
	//var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(postAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_contents_for_post_albums/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(postAlbums)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonValidPostsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all contents for post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.PostAlbumContentFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumContentFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var locations = handler.LocationService.FindAllLocationsForPostAlbums(postAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_post_albums/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(postAlbums)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonLocationsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all locations for post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(postAlbums)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post_album_tag_post_albums/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(postAlbums)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonTagsDTO))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		print("Fail")
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find all tags for post album tag post albums!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostAlbumTagPostAlbumsFullDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumHandler",
			"action":   "FAFPA586",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumTagPostAlbumsFullDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	var postAlbumsDTOS = handler.CreatePostAlbumsDTOList(convertListPostAlbumDTOToListPostAlbum(postAlbums),contents,locations,tags)

	postAlbumsJson, _ := json.Marshal(postAlbumsDTOS)
	w.Write(postAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumHandler",
		"action":   "FAFPA586",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found all following post albums!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func convertPostAlbumToPostAlbumDTO(postAlbum model.PostAlbum) dto.PostAlbumFullDTO{
	layout := "2006-01-02T15:04:05.000Z"
	var postAlbumDTO = dto.PostAlbumFullDTO{
		ID:           postAlbum.ID,
		Description:  postAlbum.Description,
		CreationDate: postAlbum.CreationDate.Format(layout),
		UserID:       postAlbum.UserID,
		LocationId:   postAlbum.LocationId,
		IsDeleted:    postAlbum.IsDeleted,
	}
	return postAlbumDTO
}

func convertListPostAlbumToListPostAlbumDTO(postAlbums []model.PostAlbum) []dto.PostAlbumFullDTO{
	var postAlbumsDTO []dto.PostAlbumFullDTO
	for i := 0; i < len(postAlbums); i++ {
		postAlbumsDTO=append(postAlbumsDTO,convertPostAlbumToPostAlbumDTO(postAlbums[i]))
	}
	return postAlbumsDTO
}

func convertPostAlbumDTOToPostAlbum(postAlbumDTO dto.PostAlbumFullDTO) model.PostAlbum{
	layout := "2006-01-02T15:04:05.000Z"
	date, _ :=time.Parse(layout, postAlbumDTO.CreationDate)
	var postAlbum = model.PostAlbum{
		Post: model.Post{
			ID:           postAlbumDTO.ID,
			Description:  postAlbumDTO.Description,
			CreationDate: date,
			UserID:       postAlbumDTO.UserID,
			LocationId:   postAlbumDTO.LocationId,
			IsDeleted:    postAlbumDTO.IsDeleted,
		},

	}
	return postAlbum
}

func convertListPostAlbumDTOToListPostAlbum(postAlbumsDTO []dto.PostAlbumFullDTO) []model.PostAlbum{
	var postAlbums []model.PostAlbum
	for i := 0; i < len(postAlbumsDTO); i++ {
		postAlbums=append(postAlbums,convertPostAlbumDTOToPostAlbum(postAlbumsDTO[i]))
	}
	return postAlbums
}