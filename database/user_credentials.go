package database

import (
	"../bcrypt"
)

type UserCredential struct {
	Username string
	Password string
}

type ProtectedUserCredential struct {
	ProtectedUsername []byte
	ProtectedPassword []byte
}

func LoginWithCredentials(userCredential UserCredential) *User {
	rows := queryDatabase("SELECT * FROM " + usersTableName)
	var user User
	for rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Credentials.ProtectedUsername,
			&user.Credentials.ProtectedPassword,
			&user.Info.Name,
			&user.Info.AccountNumber,
			&user.Info.IsAdmin)
		if userCredential.IsEqualTo(user.Credentials) {
			rows.Close()
			return &user
		}
	}
	rows.Close()
	return nil
}

func (user *UserCredential) IsEqualTo(protectedUser ProtectedUserCredential) bool {
	err := bcrypt.CompareHashAndPassword(protectedUser.ProtectedUsername, []byte(user.Username))
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(protectedUser.ProtectedPassword, []byte(user.Password))
	return err == nil
}

func (user UserCredential) ProtectUser() ProtectedUserCredential {
	var protectedUser ProtectedUserCredential
	protectedUser.ProtectedUsername, _ = bcrypt.GenerateFromPassword([]byte(user.Username), bcrypt.DefaultCost)
	protectedUser.ProtectedPassword, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return protectedUser
}