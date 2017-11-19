package users

import (
	"../bcrypt"
	"log"
)

type User struct {
	Username string
	Password string
}

type ProtectedUser struct {
	ProtectedUsername []byte
	ProtectedPassword []byte
}

func (user User) ProtectUser() ProtectedUser {
	var protectedUser ProtectedUser
	var err error
	protectedUser.ProtectedUsername, err = bcrypt.GenerateFromPassword([]byte(user.Username), bcrypt.DefaultCost)
	handleProtectUsernameError(err)
	protectedUser.ProtectedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	handleProtectPasswordError(err)
	return protectedUser
}

func handleProtectUsernameError(err error) {
	if err != nil {
		log.Fatal("Failed to protect username: ", err)
	}
}
func handleProtectPasswordError(err error) {
	if err != nil {
		log.Fatal("Failed to protect password: ", err)
	}
}

func (user *User) IsEqualTo(protectedUser ProtectedUser) bool {
	err := bcrypt.CompareHashAndPassword(protectedUser.ProtectedUsername, []byte(user.Username))
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(protectedUser.ProtectedPassword, []byte(user.Password))
	return err == nil
}

func (protectedUser *ProtectedUser) HasUsername(username string) bool {
	err := bcrypt.CompareHashAndPassword(protectedUser.ProtectedUsername, []byte(username))
	return err == nil
}