package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type VerificationRequestHandler struct {
	Service   *service.VerificationRequestService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}
var pathPostGlobal = ""

func (handler *VerificationRequestHandler) CreateVerificationRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var verificationRequestDTO dto.VerificationRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&verificationRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "CREVERIFREQ6631",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to VerificationRequestDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}
	if err := handler.Validator.Struct(&verificationRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "CREVERIFREQ6631",
			"timestamp": time.Now().String(),
		}).Error("VerificationRequestDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	category := model.INFLUENCER
	if verificationRequestDTO.RegisteredUserCategory == "INFLUENCER" {
		category = model.INFLUENCER
	}else if verificationRequestDTO.RegisteredUserCategory == "SPORTS" {
		category = model.SPORTS
	}else if verificationRequestDTO.RegisteredUserCategory == "NEW_MEDIA" {
		category = model.NEW_MEDIA
	}else if verificationRequestDTO.RegisteredUserCategory == "BUSINESS" {
		category = model.BUSINESS
	}else if verificationRequestDTO.RegisteredUserCategory == "BRAND" {
		category = model.BRAND
	}else if verificationRequestDTO.RegisteredUserCategory == "ORGANIZATION" {
		category = model.ORGANIZATION
	}

	verificationRequest := model.VerificationRequest{
		ID:                        uuid.UUID{},
		FirstName:                 verificationRequestDTO.FirstName,
		LastName:                  verificationRequestDTO.LastName,
		OfficialDocumentPath:      pathPostGlobal,
		RegisteredUserCategory:    category,
		VerificationRequestStatus: model.PENDING,
		UserId:                    verificationRequestDTO.UserId,
	}

	if err := handler.Service.CreateVerificationRequest(&verificationRequest); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "CREVERIFREQ6631",
			"timestamp": time.Now().String(),
		}).Error("Failed creating verification request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	pathPostGlobal = ""

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "VerificationRequestHandler",
		"action":    "CREVERIFREQ6631",
		"timestamp": time.Now().String(),
	}).Info("Successfully created verification request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *VerificationRequestHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("X-XSS-Protection", "1; mode=block")
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myPostFile")
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "UPK522",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find the file!")
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"), "*"+hand.Filename)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "UPK522",
			"timestamp":   time.Now().String(),
		}).Error("Failed to create temporary file!")
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "UPK522",
			"timestamp":   time.Now().String(),
		}).Error("Failed to read from file!")
		return
	}
	tempFile.Write(fileBytes)

	pathPostGlobal = tempFile.Name()[20:len(tempFile.Name())]

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "VerificationRequestHandler",
		"action":   "UPK522",
		"timestamp":   time.Now().String(),
	}).Info("Successfully uploaded the media!")
	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func (handler *VerificationRequestHandler) FindVerificationRequestById(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "FindVerificationRequestById",
			"timestamp":   time.Now().String(),
		}).Error("User is not logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-verification-request-by-id-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "FindVerificationRequestById",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	var request = handler.Service.FindById(uuid.MustParse(id))
	if request == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "FindVerificationRequestById",
			"timestamp": time.Now().String(),
		}).Error("Request by id not found!")
		fmt.Println("Request by id not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	requestJson, _ := json.Marshal(request)
	w.Write(requestJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "VerificationRequestHandler",
		"action":    "FindVerificationRequestById",
		"timestamp": time.Now().String(),
	}).Info("Successfully found request by id!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *VerificationRequestHandler) RejectVerificationRequest(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "RejectVerificationRequest",
			"timestamp":   time.Now().String(),
		}).Error("User is not logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-update-status-verification-request-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "RejectVerificationRequest",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")
	var request = handler.Service.FindById(uuid.MustParse(id))
	if request == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "RejectVerificationRequest",
			"timestamp": time.Now().String(),
		}).Error("Reject verification request not found!")
		fmt.Println("Reject verification request not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.Service.UpdateVerificationRequestRejected(uuid.MustParse(id))
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "VerificationRequestHandler",
		"action":    "RejectVerificationRequest",
		"timestamp": time.Now().String(),
	}).Info("Successfully created reject verification request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *VerificationRequestHandler) AcceptVerificationRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var verificationRequestAcceptDTO dto.VerificationRequestAcceptDTO
	if err := json.NewDecoder(r.Body).Decode(&verificationRequestAcceptDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "AcceptVerificationRequest",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to VerificationRequestAcceptDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Service.UpdateVerificationRequestAccepted(verificationRequestAcceptDTO.ID); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "AcceptVerificationRequest",
			"timestamp": time.Now().String(),
		}).Error("Fail to update verification request to accept!")
		fmt.Println("Fail to update verification request to accept")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "VerificationRequestHandler",
		"action":    "AcceptVerificationRequest",
		"timestamp": time.Now().String(),
	}).Info("Successfully updated verification request to accepted!")


	reqUrl := fmt.Sprintf("http://%s:%s/update_user_category?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), verificationRequestAcceptDTO.UserId)
	jsonClassicUsersDTO, _ := json.Marshal(verificationRequestAcceptDTO.RegisteredUserCategory)
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "AcceptVerificationRequest",
			"timestamp":   time.Now().String(),
		}).Error("Failed to update user category!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	reqUrl = fmt.Sprintf("http://%s:%s/update_official_document_path/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), verificationRequestAcceptDTO.UserId)
	jsonODP, _ := json.Marshal(verificationRequestAcceptDTO.OfficialDocumentPath)
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonODP))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "AcceptVerificationRequest",
			"timestamp":   time.Now().String(),
		}).Error("Failed to update official document path for user!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *VerificationRequestHandler) FindAllPendingVerificationRequests(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "FindAllPendingVerificationRequests",
			"timestamp":   time.Now().String(),
		}).Error("User is not logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-pending-verification-requests-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "FindAllPendingVerificationRequests",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}


	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var requests = handler.Service.FindAllPendingVerificationRequests()

	requestsJson, _ := json.Marshal(requests)
	w.Write(requestsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "VerificationRequestHandler",
		"action":    "FindAllPendingVerificationRequests",
		"timestamp": time.Now().String(),
	}).Info("Successfully found all pending verification requests!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}