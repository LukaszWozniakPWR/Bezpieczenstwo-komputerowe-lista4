package main

import (
	"./database"
)

var userCredentials = []database.UserCredential {
	database.UserCredential{"admin", "admin"},
	database.UserCredential{"jankowalski", "jankowalski1"},
	database.UserCredential{"tomasznowak", "tomasznowak1"},
	database.UserCredential{"adwersarz", "adwersarz1"},
}

var userInfo = []database.UserInfo {
	database.UserInfo{"Admin", "1234", true},
	database.UserInfo{"Jan Kowalski","2345", false},
	database.UserInfo{"Tomasz Nowak","3456", false},
	database.UserInfo{"Adwersarz","4567", false},
}

func main() {
	database.Initialize()
	database.RecreateDatabase()
	for index, userCredential := range userCredentials {
		database.CreateNewUser(userCredential, userInfo[index])
	}
}
