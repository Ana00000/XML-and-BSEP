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
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	tagModel "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	tagsService "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type PostAlbumHandler struct {
	Service * service.PostAlbumService
	PostService * service.PostService
	ClassicUserService * userService.ClassicUserService
	ClassicUserFollowingsService * userService.ClassicUserFollowingsService
	ProfileSettings *settingsService.ProfileSettingsService
	PostAlbumContentService *contentService.PostAlbumContentService
	LocationService *locationService.LocationService
	PostAlbumTagPostAlbumsService *tagsService.PostAlbumTagPostAlbumsService
	TagService *tagsService.TagService
}

func (handler *PostAlbumHandler) CreatePostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumDTO dto.PostAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumDTO)

	if err != nil {
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
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.PostService.CreatePost(&postAlbum.Post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	postAlbumIDJson, _ := json.Marshal(postAlbum.ID)
	w.Write(postAlbumIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) FindAllAlbumPostsForLoggedUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var albumPosts = handler.Service.FindAllAlbumPostsForUser(uuid.MustParse(id))
	var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(albumPosts)
	var locations = handler.LocationService.FindAllLocationsForPostAlbums(albumPosts)
	var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(albumPosts)
	var albumPostsDTOS = handler.CreatePostAlbumsDTOList(albumPosts,contents,locations,tags)

	albumPostsJson, _ := json.Marshal(albumPostsDTOS)
	w.Write(albumPostsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) CreatePostAlbumsDTOList(albums []model.PostAlbum, contents []contentModel.PostAlbumContent, locations []locationModel.Location, tags []tagModel.PostAlbumTagPostAlbums) []dto.SelectedPostAlbumDTO {
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

				if contents[j].Type == contentModel.VIDEO {
					postAlbumDTO.Type = append(postAlbumDTO.Type, "VIDEO")
				} else if contents[j].Type == contentModel.PICTURE {
					postAlbumDTO.Type = append(postAlbumDTO.Type, "PICTURE")
				}
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
				listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))
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
	if postAlbum == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if postAlbum.IsDeleted == true{
		fmt.Println("Deleted post album")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if postAlbum.UserID != uuid.MustParse(logId){
		fmt.Println("Post album doesnt belong to user")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbum(postAlbum)
	var locations = handler.LocationService.FindAllLocationsForPostAlbum(postAlbum)
	var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbum(postAlbum)

	var postAlbumDTO = handler.CreatePostAlbumDTO(postAlbum,contents,locations,tags)

	postAlbumJson, _ := json.Marshal(postAlbumDTO)
	w.Write(postAlbumJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *PostAlbumHandler) CreatePostAlbumDTO(album *model.PostAlbum, contents []contentModel.PostAlbumContent, locations []locationModel.Location, tags []tagModel.PostAlbumTagPostAlbums) dto.SelectedPostAlbumDTO {
	var postAlbumDTO dto.SelectedPostAlbumDTO

	postAlbumDTO.PostAlbumId = album.ID
	postAlbumDTO.Description = album.Description
	postAlbumDTO.CreationDate = album.CreationDate


	for j := 0; j < len(contents); j++ {
		if contents[j].PostAlbumId == album.ID {
			postAlbumDTO.Path = append(postAlbumDTO.Path, contents[j].Path)

			if contents[j].Type == contentModel.VIDEO {
				postAlbumDTO.Type = append(postAlbumDTO.Type, "VIDEO")
			} else if contents[j].Type == contentModel.PICTURE {
				postAlbumDTO.Type = append(postAlbumDTO.Type, "PICTURE")
			}
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
			listOfTags = append(listOfTags, handler.TagService.FindTagNameById(tags[p].TagId))

		}
	}

	postAlbumDTO.Tags = listOfTags

	return postAlbumDTO
}

func (handler *PostAlbumHandler) FindAllPublicAlbumPostsRegisteredUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	// returns only VALID users but loggedIn user
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidAlbumPosts = handler.Service.FindAllPublicAndFriendsPostAlbumsValid(allPublicUsers)
	var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(publicValidAlbumPosts)
	var locations = handler.LocationService.FindAllLocationsForPostAlbums(publicValidAlbumPosts)
	var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(publicValidAlbumPosts)

	var postAlbumsDTOS = handler.CreatePostAlbumsDTOList(publicValidAlbumPosts,contents,locations,tags)

	postAlbumsJson, _ := json.Marshal(postAlbumsDTOS)
	w.Write(postAlbumsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) FindAllPublicAlbumPostsNotRegisteredUser(w http.ResponseWriter, r *http.Request) {

	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var allPublicUsers = handler.ProfileSettings.FindAllPublicUsers(allValidUsers)
	var publicValidAlbumPosts = handler.Service.FindAllPublicAndFriendsPostAlbumsValid(allPublicUsers)
	var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(publicValidAlbumPosts)
	var locations = handler.LocationService.FindAllLocationsForPostAlbums(publicValidAlbumPosts)
	var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(publicValidAlbumPosts)
	var postAlbumsDTOS = handler.CreatePostAlbumsDTOList(publicValidAlbumPosts,contents,locations,tags)

	postAlbumsJson, _ := json.Marshal(postAlbumsDTOS)
	w.Write(postAlbumsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumHandler) FindAllFollowingPostAlbums(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// returns only valid users
	var allValidUsers = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))

	// retuns only valid FOLLOWINGS
	var followings = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), allValidUsers)

	// returns NOT DELETED POST ALBUMS from valid following users
	var postAlbums = handler.Service.FindAllFollowingPostAlbums(followings)
	var contents = handler.PostAlbumContentService.FindAllContentsForPostAlbums(postAlbums)
	var locations = handler.LocationService.FindAllLocationsForPostAlbums(postAlbums)
	var tags = handler.PostAlbumTagPostAlbumsService.FindAllTagsForPostAlbumTagPostAlbums(postAlbums)
	var postAlbumsDTOS = handler.CreatePostAlbumsDTOList(postAlbums,contents,locations,tags)

	postAlbumsJson, _ := json.Marshal(postAlbumsDTOS)
	w.Write(postAlbumsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}