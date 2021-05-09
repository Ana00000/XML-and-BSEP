package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"golang.org/x/crypto/bcrypt"
	_ "strconv"
)


type UserHandler struct {
	Service * service.UserService
}

//for login
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}



