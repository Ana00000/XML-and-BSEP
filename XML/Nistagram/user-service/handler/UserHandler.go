package handler

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	_ "strconv"
	"strings"
	"time"
)

type UserHandler struct {

	UserService * service.UserService
	AdminService * service.AdminService
	ClassicUserService * service.ClassicUserService
	AgentService * service.AgentService
	Rbac * gorbac.RBAC
	PermissionFindAllUsers *gorbac.Permission
	RegisteredUserService * service.RegisteredUserService
	RecoveryPasswordTokenService *service.RecoveryPasswordTokenService
	PermissionFindUserByID * gorbac.Permission
	PermissionUpdateUserInfo * gorbac.Permission
	Validator                *validator.Validate
	PasswordUtil			 *util.PasswordUtil
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

//FIDALUSRS2330
func (handler *UserHandler) FindAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	err := TokenValid(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDALUSRS2330",
			"timestamp":   time.Now().String(),
		}).Error("User need to be logged in!")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := r.URL.Query().Get("id")
	var loginUser = handler.UserService.FindByID(uuid.MustParse(id))

	userRole := ""
	if loginUser.UserType == model.ADMIN {
		userRole = "role-admin"
	} else if loginUser.UserType == model.AGENT {
		userRole = "role-agent"
	} else {
		userRole = "role-registered-user"
	}
	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllUsers, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDALUSRS2330",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var users []model.User
	users = handler.UserService.FindAllUsers()
	usersJson, _ := json.Marshal(users)
	if usersJson != nil {
		w.WriteHeader(http.StatusOK)
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded all users!")
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJson)
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "UserHandler",
		"action":   "FIDALUSRS2330",
		"timestamp":   time.Now().String(),
	}).Error("Failed founding all users!")
	w.WriteHeader(http.StatusBadRequest)
}

func CreateToken(userName string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userName
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (handler *UserHandler) GetUserIDFromJWTToken(w http.ResponseWriter, r *http.Request){
	token, err := VerifyToken(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "GetUserIDFromJWTToken",
			"timestamp":   time.Now().String(),
		}).Error("Failed verified token!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId:= fmt.Sprintf("%s", claims["user_id"])
		retValJson,_ := json.Marshal(userId)
		w.Write(retValJson)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "UserHandler",
		"action":   "GetUserIDFromJWTToken",
		"timestamp":   time.Now().String(),
	}).Error("Token doesn't valid!")
	w.WriteHeader(http.StatusBadRequest)
}

func getUserNameFromJWT(r *http.Request) (string,error) {
	token, err := VerifyToken(r)
	if err!=nil{
		return "",err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId:= fmt.Sprintf("%s", claims["user_id"])
		return userId, nil
	}
	return "",err
}



//CHUSPASS9112
func (handler *UserHandler) ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var userChangePasswordDTO dto.UserChangePasswordDTO
	err := json.NewDecoder(r.Body).Decode(&userChangePasswordDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to UserChangePasswordDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var recoveryPasswordToken = handler.RecoveryPasswordTokenService.FindByID(userChangePasswordDTO.RecoveryPasswordTokenID)

	if recoveryPasswordToken==nil || recoveryPasswordToken.Status!=model.VERIFIED{
		if recoveryPasswordToken.Status!=model.VERIFIED{
			err = handler.RecoveryPasswordTokenService.UpdateRecoveryPasswordTokenValidity(recoveryPasswordToken.RecoveryPasswordToken, model.INVALID)
			if err != nil {
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "UserHandler",
					"action":   "CHUSPASS9112",
					"timestamp":   time.Now().String(),
				}).Error("Failed updating recovery token to invalid!")
				return
			}
		}
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Recovery password token not found or his status not appropriate!")
		w.WriteHeader(http.StatusConflict)
		return
	}

	var recoveryPasswordTokenUser = handler.UserService.FindByID(recoveryPasswordToken.UserId)
	if recoveryPasswordTokenUser==nil || recoveryPasswordTokenUser.Email!=userChangePasswordDTO.Email{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Recovery password user not exist or user email not match with recovery password token user!")
		w.WriteHeader(http.StatusConflict)
		return
	}

	if userChangePasswordDTO.Password != userChangePasswordDTO.ConfirmedPassword {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Password and his confirmation doesn't match!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = handler.UserService.FindByEmail(userChangePasswordDTO.Email)
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Not founded user with entered email!")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(userChangePasswordDTO.Password)

	if validPassword {
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(userChangePasswordDTO.Password)
	}else {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Password format is invalid!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	err = handler.UserService.UpdateUserPassword(user.ID, salt, password)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating user info!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if user.UserType == model.ADMIN {
		err = handler.AdminService.UpdateAdminPassword(user.ID, salt, password)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "CHUSPASS9112",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating admin info!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else if user.UserType == model.AGENT {
		err = handler.ClassicUserService.UpdateClassicUserPassword(user.ID, salt, password)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "CHUSPASS9112",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating classic user info!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		err = handler.AgentService.UpdateAgentPassword(user.ID, salt, password)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "CHUSPASS9112",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating agent info!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else {
		err = handler.ClassicUserService.UpdateClassicUserPassword(user.ID, salt, password)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "CHUSPASS9112",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating classic user info!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		err = handler.RegisteredUserService.UpdateRegisteredUserPassword(user.ID, salt, password)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "CHUSPASS9112",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating registered user info!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}

	err = handler.RecoveryPasswordTokenService.UpdateRecoveryPasswordTokenValidity(recoveryPasswordToken.RecoveryPasswordToken, model.INVALID)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHUSPASS9112",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating recovery token to invalid!")
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "CHUSPASS9112",
		"timestamp":   time.Now().String(),
	}).Info("Successfully changed user info!")
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
}

//LOG85310
func (handler *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var logInUserDTO dto.LogInUserDTO
	if err := json.NewDecoder(r.Body).Decode(&logInUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "LOG85310",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to LogInUserDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&logInUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "LOG85310",
			"timestamp":   time.Now().String(),
		}).Error("LogInUserDTO fields doesn't entered in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	var user = handler.UserService.FindByUserName(logInUserDTO.Username)

	if user.UserType == model.ADMIN {
		if user == nil || !user.IsConfirmed {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "LOG85310",
				"timestamp":   time.Now().String(),
			}).Error("Admin is not confirmed!")
			w.WriteHeader(http.StatusBadRequest) //400
			return
		}
	} else {
		var classicUser = handler.ClassicUserService.FindClassicUserByUserName(logInUserDTO.Username)
		if user == nil || !classicUser.IsConfirmed || classicUser.IsDeleted {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "LOG85310",
				"timestamp":   time.Now().String(),
			}).Error("Classic user is not valid!")
			w.WriteHeader(http.StatusBadRequest) //400
			return
		}
	}

	validPassword := handler.PasswordUtil.IsValidPassword(logInUserDTO.Password)
	plainPassword := ""

	if validPassword {
		var sb strings.Builder
		salt := user.Salt
		sb.WriteString(logInUserDTO.Password)
		sb.WriteString(salt)
		plainPassword = sb.String()
	}else {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "LOG85310",
			"timestamp":   time.Now().String(),
		}).Error("Password doesn't entered in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if !handler.PasswordUtil.CheckPasswordHash(plainPassword, user.Password) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "LOG85310",
			"timestamp":   time.Now().String(),
		}).Error("Failed sign up!")
		w.WriteHeader(http.StatusConflict)
		return
	}

	//QUESTION CHECK
	if logInUserDTO.Question != user.Question{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "LOG85310",
			"timestamp":   time.Now().String(),
		}).Error("Wrong question for user!")
		w.WriteHeader(http.StatusConflict)
		return
	}

	//ANSWER CHECK

	//ANSWER AND QUESTION
	plainAnswer := ""
	var ab strings.Builder
	answerSalt := user.AnswerSalt
	ab.WriteString(logInUserDTO.Answer)
	ab.WriteString(answerSalt)
	plainAnswer = ab.String()

	if !handler.PasswordUtil.CheckPasswordHash(plainAnswer, user.Answer){
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "LOG85310",
			"timestamp":   time.Now().String(),
		}).Error("Wrong answer to user question!")
		w.WriteHeader(http.StatusConflict)
		return
	}
	//ANSWER CHECK END

	//token
	token, err := CreateToken(user.Username)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "LOG85310",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating AWT token!")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	logInResponse := dto.LogInResponseDTO{
		ID:       user.ID,
		Token:    token,
		UserType: user.UserType,
	}

	logInResponseJson, _ := json.Marshal(logInResponse)
	w.Write(logInResponseJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "LOG85310",
		"timestamp":   time.Now().String(),
	}).Info("Successfully sign up user! User: "+user.Username)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

//UPDUSPROFINF393
func (handler *UserHandler) UpdateUserProfileInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "UPDUSPROFINF393",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	var userDTO dto.UserUpdateProfileInfoDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "UPDUSPROFINF393",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to UserUpdateProfileInfoDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&userDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "UPDUSPROFINF393",
			"timestamp":   time.Now().String(),
		}).Error("UserUpdateProfileInfoDTO fields aren't entered in valid format!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var loginUser = handler.UserService.FindByID(userDTO.ID)
	userRole := ""
	if loginUser.UserType == model.ADMIN {
		userRole = "role-admin"
	} else if loginUser.UserType == model.AGENT {
		userRole = "role-agent"
	} else {
		userRole = "role-registered-user"
	}
	if !handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateUserInfo, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "UPDUSPROFINF393",
			"timestamp":   time.Now().String(),
		}).Error("User aren't authorized to update user information!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := handler.UserService.UpdateUserProfileInfo(&userDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "UPDUSPROFINF393",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating basic user profile information!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	if userDTO.UserType == "ADMIN" {
		err = handler.AdminService.UpdateAdminProfileInfo(&userDTO)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "UPDUSPROFINF393",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating admin profile information!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else if userDTO.UserType == "AGENT" {
		err = handler.AgentService.UpdateAgentProfileInfo(&userDTO)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "UPDUSPROFINF393",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating agent profile information!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		err = handler.ClassicUserService.UpdateClassicUserProfileInfo(&userDTO)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "UPDUSPROFINF393",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating classic user profile information!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else {
		err = handler.RegisteredUserService.UpdateRegisteredUserProfileInfo(&userDTO)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "UPDUSPROFINF393",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating registered user profile information!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		err = handler.ClassicUserService.UpdateClassicUserProfileInfo(&userDTO)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserHandler",
				"action":   "UPDUSPROFINF393",
				"timestamp":   time.Now().String(),
			}).Error("Failed updating classic user profile information!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "UPDUSPROFINF393",
		"timestamp":   time.Now().String(),
	}).Info("Successfully updated user profile info!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDBYID0329
func (handler *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDBYID0329",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDBYID0329",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSignIn = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSignIn)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindUserByID, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDBYID0329",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	var user = handler.UserService.FindByID(uuid.MustParse(id))
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDBYID0329",
			"timestamp":   time.Now().String(),
		}).Error("User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "FIDBYID0329",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded user by id!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//GetByID
func (handler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	var user = handler.UserService.FindByID(uuid.MustParse(id))
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "GetByID",
			"timestamp":   time.Now().String(),
		}).Error("User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	var classicUser = convertClassicUserToClassicUserDTO(*handler.ClassicUserService.FindById(user.ID))
	userJson, _ := json.Marshal(classicUser)
	w.Write(userJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "GetByID",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded user by id!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALUSRSBUTLOGGIN212
func (handler *UserHandler) FindAllUsersButLoggedIn(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	var user = handler.UserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	if  user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDALUSRSBUTLOGGIN212",
			"timestamp":   time.Now().String(),
		}).Error("No user found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "FIDALUSRSBUTLOGGIN212",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all users without logged in!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALPUBUSRS0291
func (handler *UserHandler) FindAllPublicUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var profileSettings []Data
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_for_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	err := getJson(reqUrl, &profileSettings)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDALPUBUSRS0291",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list ids!")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	/*
		var profileSettings = handler.ProfileSettingsService.FindAllProfileSettingsForPublicUsers()
	*/
	var users []model.User
	users = handler.UserService.FindAllPublicUsers(convertListDataToListUUID(profileSettings))

	usersJson, _ := json.Marshal(users)
	if usersJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "RegisteredUserHandler",
			"action":   "FIDALPUBUSRS0291",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded all public users!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "UserHandler",
		"action":   "FIDALPUBUSRS0291",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding all public users!")
	w.WriteHeader(http.StatusBadRequest)
}

//FIDBYUSNAM9482
func (handler *UserHandler) FindByUserName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	username := r.URL.Query().Get("username")

	var user = handler.UserService.FindByUserName(username)
	if  user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "FIDBYUSNAM9482",
			"timestamp":   time.Now().String(),
		}).Error("User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "FIDBYUSNAM9482",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded user by username!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//CHCKIFAUTH9342
func (handler *UserHandler) CheckIfAuthentificated(w http.ResponseWriter, r *http.Request) {
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserHandler",
			"action":   "CHCKIFAUTH9342",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}
	w.WriteHeader(http.StatusOK)
}

type Data struct {
	Uuid uuid.UUID
}

func convertListDataToListUUID(datas []Data) []uuid.UUID {
	var uuids []uuid.UUID
	for i := 0; i < len(datas); i++ {
		uuids = append(uuids, datas[i].Uuid)
	}
	return uuids
}
