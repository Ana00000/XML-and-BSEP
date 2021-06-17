package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

	verificationRequest := model.VerificationRequest{
		ID:                     uuid.UUID{},
		FirstName:              verificationRequestDTO.FirstName,
		LastName:               verificationRequestDTO.LastName,
		OfficialDocumentPath:   pathPostGlobal,
		RegisteredUserCategory: verificationRequestDTO.RegisteredUserCategory,
		VerificationRequestStatus: model.PENDING,
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
			"action":   "UPK523",
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
			"location":   "VerificationRequestHandler",
			"action":   "UPK523",
			"timestamp":   time.Now().String(),
		}).Error("Failed to read from file!")
		return
	}
	tempFile.Write(fileBytes)

	pathPostGlobal = tempFile.Name()[20:len(tempFile.Name())]

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "VerificationRequestHandler",
		"action":   "UPK523",
		"timestamp":   time.Now().String(),
	}).Info("Successfully uploaded the media!")
	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func (handler *VerificationRequestHandler) FindRequestById(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "FIDREQBYID2431",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-request-by-id-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "VerificationRequestHandler",
			"action":   "FIDREQBYID2431",
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
			"action":    "FIDREQBYID2431",
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
		"action":    "FIDREQBYID2431",
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
			"location":   "FollowRequestHandler",
			"action":   "REJFOLLOWREQ4939",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-reject-follow-request-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "FollowRequestHandler",
			"action":   "REJFOLLOWREQ4939",
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
			"location":  "FollowRequestHandler",
			"action":    "REJFOLLOWREQ4939",
			"timestamp": time.Now().String(),
		}).Error("Reject follow request not found!")
		fmt.Println("Reject follow request not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.Service.UpdateFollowRequestRejected(uuid.MustParse(id))
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "REJFOLLOWREQ4939",
		"timestamp": time.Now().String(),
	}).Info("Successfully created reject follow request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *FollowRequestHandler) AcceptVerificationRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	requestId := vars["requestID"]
	if err := handler.Service.UpdateFollowRequestAccepted(uuid.MustParse(requestId)); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "UPDFOLLOWREQTOACCEP7710",
			"timestamp": time.Now().String(),
		}).Error("Fail to update follow request to accept!")
		fmt.Println("Fail to update follow request to accept")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "UPDFOLLOWREQTOACCEP7710",
		"timestamp": time.Now().String(),
	}).Info("Successfully updated follow request to accepted!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
