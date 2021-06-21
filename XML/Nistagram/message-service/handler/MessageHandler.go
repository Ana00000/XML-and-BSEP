package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/service"
	gomail "gopkg.in/mail.v2"
	"net/http"
	"os"
	"time"
)

type MessageHandler struct {
	Service * service.MessageService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var messageDTO dto.MessageDTO
	err := json.NewDecoder(r.Body).Decode(&messageDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageHandler",
			"action":   "CRMEE454",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to MessageDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, messageDTO.CreationDate)
	id := uuid.UUID{}

	message := model.Message{
		ID:                 id,
		MessageSubstanceId: messageDTO.MessageContentID,
		IsDisposable:       messageDTO.IsDisposable,
		CreationDate:       creationDate,
		SenderUserID:       messageDTO.SenderUserID,
		ReceiverUserID:     messageDTO.ReceiverUserID,
		IsDeleted:          messageDTO.IsDeleted,
	}

	err = handler.Service.CreateMessage(&message)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageHandler",
			"action":   "CRMEE454",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating message!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//send email to receiverUserId
	var receiverUser dto.ClassicUserDTO
	reqUrlUser := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), messageDTO.ReceiverUserID)
	err = getJson(reqUrlUser, &receiverUser)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageHandler",
			"action":   "CRMEE454",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find receiver user by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var senderUser dto.ClassicUserDTO
	reqUrlUser = fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), messageDTO.SenderUserID)
	err = getJson(reqUrlUser, &senderUser)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageHandler",
			"action":   "CRMEE454",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find sender user by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//SEND EMAIL NOTIFICATION
	handler.SendNotificationMail(receiverUser.Email, id, senderUser.Username)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "MessageHandler",
		"action":   "CRMEE454",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created message!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *MessageHandler) SendNotificationMail(email string, messageId uuid.UUID, userName string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "xml.ftn.uns@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Confirmation mail")

	// Set E-Mail body. You can set plain text or html with text/html
	text := "New message from "+ userName +"!\n\nhttps://localhost:8081/messageById/" + messageId.String() + "\n\n\nBest regards,\nTim25"
	m.SetBody("text/plain", text)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "xml.ftn.uns@gmail.com", "XMLFTNUNS1")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		//fmt.Println(err)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageHandler",
			"action":   "SendNotificationMail",
			"timestamp":   time.Now().String(),
		}).Error("Failed to send an email with the confirmation token!")
		panic(err)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "MessageHandler",
		"action":   "SEDCONFMAIL227",
		"timestamp":   time.Now().String(),
	}).Info("Successfully sent email with confirmation token!")

}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}