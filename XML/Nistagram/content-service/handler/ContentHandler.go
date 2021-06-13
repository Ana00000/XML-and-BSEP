package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"net/http"
	"os"
	_ "strconv"
	"strings"
	"time"
)

type ContentHandler struct {
	Service * service.ContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func Request(url string, token string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	tokenString := "Bearer "+token
	req.Header.Set("Authorization", tokenString)
	resp, err := http.DefaultClient.Do(req)
	return resp
}

func (handler *ContentHandler) CreateContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var contentDTO dto.ContentDTO
	err := json.NewDecoder(r.Body).Decode(&contentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentHandler",
			"action":   "CRCOU658",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch contentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	content := model.Content{
		ID:   id,
		Path: contentDTO.Path,
		Type: contentType,
	}

	err = handler.Service.CreateContent(&content)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentHandler",
			"action":   "CRCOU658",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ContentHandler",
		"action":   "CRCOU658",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
