package handlers

import (
	"net/http"
	"github.com/gorilla/securecookie"
	"../database"
	"encoding/json"
	"log"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func setSession(user *database.User, transaction *database.Transaction, response http.ResponseWriter) {
	userJson, err := json.Marshal(*user)
	if err != nil {
		log.Fatal("error converting user to json")
	}
	var transactionJson []byte
	if transaction != nil {
		transactionJson, err = json.Marshal(*transaction)
	}
	value := map[string]string{
		"user": string(userJson),
		"transaction": string(transactionJson),
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func getLoggedUserFromCookies(request *http.Request) *database.User {
	var user database.User
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			json.Unmarshal([]byte(cookieValue["user"]), &user)
		} else {
			return nil
		}
	} else {
		return nil
	}
	return &user
}

func getTransactionFromCookies(request *http.Request) *database.Transaction {
	var transaction database.Transaction
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			json.Unmarshal([]byte(cookieValue["transaction"]), &transaction)
		} else {
			return nil
		}
	} else {
		return nil
	}
	return &transaction
}