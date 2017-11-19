package main

import (
	. "./users"
	"encoding/json"
	"io/ioutil"
	"log"
)

var userList = []User{
	User{"jankowalski", "jankowalski1"},
	User{"tomasznowak", "tomasznowak1"},
}

func main() {
	var protectedUsers []ProtectedUser
	for _, user := range userList {
		protectedUsers = append(protectedUsers, user.ProtectUser())
	}
	jsonUsers, err := json.Marshal(protectedUsers)
	if err != nil {
		log.Fatal("Failed to convert users to json")
	}
	err = ioutil.WriteFile("users/users.json", jsonUsers, 0644)
	if err != nil {
		log.Fatal("Failed to write users to file")
	}
}
