package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	_ "strconv"
)

type RegisteredUserHandler struct {
	Service * service.RegisteredUserService
}

