package handler

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"github.com/mikespook/gorbac"
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
	UserService              *service.UserService
	AdminService             *service.AdminService
	ClassicUserService       *service.ClassicUserService
	AgentService             *service.AgentService
	Rbac                     *gorbac.RBAC
	PermissionFindAllUsers   *gorbac.Permission
	RegisteredUserService    *service.RegisteredUserService
	PermissionUpdateUserInfo *gorbac.Permission
	Validator                *validator.Validate
	PasswordUtil             *util.PasswordUtil
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

func (handler *UserHandler) FindAllUsers(w http.ResponseWriter, r *http.Request) {

	err := TokenValid(r)
	if err != nil {
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
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var users []model.User
	users = handler.UserService.FindAllUsers()
	usersJson, _ := json.Marshal(users)
	if usersJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJson)
	}
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

func (handler *UserHandler) ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	var userChangePasswordDTO dto.UserChangePasswordDTO

	err := json.NewDecoder(r.Body).Decode(&userChangePasswordDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userChangePasswordDTO.Password != userChangePasswordDTO.ConfirmedPassword {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = handler.UserService.FindByEmail(userChangePasswordDTO.Email)
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(userChangePasswordDTO.Password)

	if validPassword {
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(userChangePasswordDTO.Password)
	} else {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	err = handler.UserService.UpdateUserPassword(user.ID, salt, password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if user.UserType == model.ADMIN {
		err = handler.AdminService.UpdateAdminPassword(user.ID, salt, password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else if user.UserType == model.AGENT {
		err = handler.ClassicUserService.UpdateClassicUserPassword(user.ID, salt, password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		err = handler.AgentService.UpdateAgentPassword(user.ID, salt, password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else {
		err = handler.ClassicUserService.UpdateClassicUserPassword(user.ID, salt, password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		err = handler.RegisteredUserService.UpdateRegisteredUserPassword(user.ID, salt, password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var logInUserDTO dto.LogInUserDTO
	if err := json.NewDecoder(r.Body).Decode(&logInUserDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&logInUserDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	var user = handler.UserService.FindByUserName(logInUserDTO.Username)

	if user.UserType == model.ADMIN {
		if user == nil || !user.IsConfirmed {
			w.WriteHeader(http.StatusBadRequest) //400
			return
		}
	} else {
		var classicUser = handler.ClassicUserService.FindClassicUserByUserName(logInUserDTO.Username)
		if user == nil || !classicUser.IsConfirmed || classicUser.IsDeleted {
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
	} else {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if !handler.PasswordUtil.CheckPasswordHash(plainPassword, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(user.Username)
	if err != nil {
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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *UserHandler) UpdateUserProfileInfo(w http.ResponseWriter, r *http.Request) {

	if err := TokenValid(r); err != nil {
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	var userDTO dto.UserUpdateProfileInfoDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&userDTO); err != nil {
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
	if !handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateUserInfo, nil) &&
		!handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateUserInfo, nil) &&
		!handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateUserInfo, nil) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := handler.UserService.UpdateUserProfileInfo(&userDTO)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	if userDTO.UserType == "ADMIN" {
		err = handler.AdminService.UpdateAdminProfileInfo(&userDTO)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}
	} else if userDTO.UserType == "AGENT" {
		err = handler.AgentService.UpdateAgentProfileInfo(&userDTO)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}
		err = handler.ClassicUserService.UpdateClassicUserProfileInfo(&userDTO)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}
	} else {
		err = handler.RegisteredUserService.UpdateRegisteredUserProfileInfo(&userDTO)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}
		err = handler.ClassicUserService.UpdateClassicUserProfileInfo(&userDTO)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var user = handler.UserService.FindByID(uuid.MustParse(id))
	if user == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindAllUsersButLoggedIn(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var user = handler.UserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	if user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindAllPublicUsers(w http.ResponseWriter, r *http.Request) {
	var profileSettings []Data
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_for_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	err := getJson(reqUrl, &profileSettings)
	if err != nil {
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	/*
		var profileSettings = handler.ProfileSettingsService.FindAllProfileSettingsForPublicUsers()
	*/
	var users []model.User
	users = handler.UserService.FindAllPublicUsers(convertListDataToListUUID(profileSettings))

	usersJson, _ := json.Marshal(users)
	if usersJson != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJson)
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (handler *UserHandler) FindByUserName(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")

	var user = handler.UserService.FindByUserName(username)
	if user == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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
