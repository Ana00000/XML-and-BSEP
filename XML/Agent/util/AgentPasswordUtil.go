package util

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type AgentPasswordUtil struct {

}

func (util *AgentPasswordUtil)HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (util *AgentPasswordUtil)GeneratePasswordWithSalt(plainPassword string) (string, string) {
	var sb strings.Builder
	salt := uuid.New().String()
	sb.WriteString(plainPassword)
	sb.WriteString(salt)
	passwordWithSalt := sb.String()
	hashedPassword, _ := util.HashPassword(passwordWithSalt)
	return salt, hashedPassword
}

func (util *AgentPasswordUtil)CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

