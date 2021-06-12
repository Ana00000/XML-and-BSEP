package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"io/ioutil"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type SinglePostContentHandler struct {
	Service * service.SinglePostContentService
	ContentService * service.ContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

var pathPostGlobal = ""

func (handler *SinglePostContentHandler) CreateSinglePostContent(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "CRSIPOCOB123",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-single-post-content-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "CRSIPOCOB123",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}
	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "CRSIPOCOB123",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	var singlePostContentDTO dto.SinglePostContentDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostContentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "CRSIPOCOB123",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SinglePostContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch singlePostContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	singlePostContent := model.SinglePostContent{
		Content: model.Content{
			ID:   id,
			Path: pathPostGlobal,
			Type: contentType,
		},
		SinglePostId: singlePostContentDTO.SinglePostId,
	}

	err = handler.Service.CreateSinglePostContent(&singlePostContent)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "CRSIPOCOB123",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating single post content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.ContentService.CreateContent(&singlePostContent.Content)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "CRSIPOCOB123",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	pathPostGlobal = ""

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SinglePostContentHandler",
		"action":   "CRSIPOCOB123",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created single post content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostContentHandler) FindAllContentsForPosts(w http.ResponseWriter, r *http.Request) {
	var singlePostsDTOs []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTOs)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "FIALCOFOPOQ771",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SinglePostsDTOs!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contents = convertListSinglePostContentsToListSinglePostContentsForSinglePostDTO(handler.Service.FindAllContentsForPosts(singlePostsDTOs))

	pathJson, _ := json.Marshal(contents)
	w.Write(pathJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SinglePostContentHandler",
		"action":   "FIALCOFOPOQ771",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for posts!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostContentHandler) FindAllContentsForPost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "FIALCOFOPOW474",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SinglePostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contents = convertListSinglePostContentsToListSinglePostContentsForSinglePostDTO(handler.Service.FindAllContentsForPost(&singlePostDTO))

	pathJson, _ := json.Marshal(contents)
	w.Write(pathJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SinglePostContentHandler",
		"action":   "FIALCOFOPOW474",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for post!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myPostFile")
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "UPK523",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find the file!")
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"),  "*" + hand.Filename)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "UPK523",
			"timestamp":   time.Now().String(),
		}).Error("Failed to create temporary file!")
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SinglePostContentHandler",
			"action":   "UPK523",
			"timestamp":   time.Now().String(),
		}).Error("Failed to read from file!")
		return
	}
	tempFile.Write(fileBytes)


	pathPostGlobal = tempFile.Name()[20:len(tempFile.Name())]

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SinglePostContentHandler",
		"action":   "UPK523",
		"timestamp":   time.Now().String(),
	}).Info("Successfully uploaded the media!")
	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func convertListSinglePostContentsToListSinglePostContentsForSinglePostDTO(singlePostContents []model.SinglePostContent) []dto.SinglePostContentForSinglePostDTO{
	var listSinglePostContentsDTO []dto.SinglePostContentForSinglePostDTO
	for i := 0; i < len(singlePostContents); i++ {
		listSinglePostContentsDTO = append(listSinglePostContentsDTO,convertSinglePostContentToSinglePostContentForSinglePostDTO(singlePostContents[i]))
	}
	return listSinglePostContentsDTO
}

func convertSinglePostContentToSinglePostContentForSinglePostDTO(singlePostContent model.SinglePostContent) dto.SinglePostContentForSinglePostDTO{
	contentType := ""
	if singlePostContent.Type == model.VIDEO{
		contentType = "VIDEO"
	} else if singlePostContent.Type == model.PICTURE{
		contentType = "PICTURE"
	}

	var singlePostContentDTO = dto.SinglePostContentForSinglePostDTO{
		ID:           singlePostContent.ID,
		Path:         singlePostContent.Path,
		Type:         contentType,
		SinglePostId: singlePostContent.SinglePostId,
	}
	return singlePostContentDTO
}