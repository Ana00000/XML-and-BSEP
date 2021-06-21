package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"gopkg.in/go-playground/validator.v9"
	gomail "gopkg.in/mail.v2"
	"net/http"
	"os"
	"time"
)

type CommentHandler struct {
	Service   *service.CommentService
	Validator *validator.Validate
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-comment-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var commentDTO dto.CommentDTO
	if err := json.NewDecoder(r.Body).Decode(&commentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to CommentDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&commentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("CommentDTO fields aren't in the valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}
	comment := model.Comment{
		ID:           uuid.UUID{},
		CreationDate: time.Now(),
		UserID:       commentDTO.UserID,
		PostID:       commentDTO.PostID,
		Text:         commentDTO.Text,
	}

	if err := handler.Service.CreateComment(&comment); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating comment!")
		w.WriteHeader(http.StatusExpectationFailed) // 417
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CommentHandler",
		"action":   "CRCOM571",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created comment!")
	
	commentIDJson, _ := json.Marshal(comment.ID)
	w.Write(commentIDJson)

	//GET USERID FROM POSTID FOR COMMENTNOTIFICATION
	var userId uuid.UUID
	reqUrl := fmt.Sprintf("http://%s:%s/find_owner_of_post/%s", os.Getenv("POST_SERVICE_DOMAIN"), os.Getenv("POST_SERVICE_DOMAIN"), commentDTO.PostID)
	err := getJson(reqUrl, &userId)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find owner of the post!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var user dto.ClassicUserDTO
	reqUrlUser := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), userId)
	err = getJson(reqUrlUser, &user)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find user by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var profileSettings dto.ProfileSettingsDTO
	reqUrl = fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), user.ID)
	err = getJson(reqUrl, &profileSettings)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find profile settings by user id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if profileSettings.CommentsNotifications == "ALL_NOTIFICATIONS"{
		//SEND EMAIL NOTIFICATION
		handler.SendNotificationMail(user.Email, commentDTO)
	}else if profileSettings.CommentsNotifications == "FRIENDS_NOTIFICATION"{
		//check if senderUser is friend
		var followings []dto.ClassicUserFollowingsFullDTO
		reqUrl := fmt.Sprintf("http://%s:%s/find_all_valid_followings_for_user/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), user.ID)
		err = getJson(reqUrl, &followings)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "CommentHandler",
				"action":   "CRCOM571",
				"timestamp":   time.Now().String(),
			}).Error("Failed to find followings for user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		for  i:=0; i < len(followings); i++{
			if followings[i].FollowingUserId == commentDTO.UserID {
				//SEND EMAIL NOTIFICATION
				handler.SendNotificationMail(user.Email, commentDTO)
			}
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) SendNotificationMail(email string, comment dto.CommentDTO) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "xml.ftn.uns@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Confirmation mail")

	var user dto.ClassicUserDTO
	reqUrlUser := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), comment.UserID)
	err := getJson(reqUrlUser, &user)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "SEDCONFMAIL100",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find user by id!")
		panic(err)
	}

	// Set E-Mail body. You can set plain text or html with text/html
	text := user.FirstName + " " + user.LastName + " made a new comment on your post!\n\nhttps://localhost:8081/postById/" + comment.PostID.String() + "\n\n\nBest regards,\nTim25"
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
			"location":   "CommentHandler",
			"action":   "SEDCONFMAIL100",
			"timestamp":   time.Now().String(),
		}).Error("Failed sending email with confirmation token!")
		panic(err)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CommentHandler",
		"action":   "SEDCONFMAIL100",
		"timestamp":   time.Now().String(),
	}).Info("Successfully sended email with confirmation token!")
}

func (handler *CommentHandler) FindAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-comments-for-post-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}
	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllCommentsForPost(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all comments for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "CommentHandler",
		"action":   "FACFP572",
		"timestamp":   time.Now().String(),
	}).Error("Comments for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *CommentHandler) FindAllUserComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllUserComments(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "CommentHandler",
			"action":   "FAUCM573",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all user comments!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
		return
	}

	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "CommentHandler",
		"action":   "FAUCM573",
		"timestamp":   time.Now().String(),
	}).Error("All user comments not found!")
	w.WriteHeader(http.StatusBadRequest)
}
