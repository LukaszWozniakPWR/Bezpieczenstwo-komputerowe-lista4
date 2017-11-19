package handlers

import (
	"net/http"
	"log"
	"fmt"
	."../users"
)

func LoginHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		displayLoginPage(responseWriter, request)
	case "POST":
		processLoginRequest(responseWriter, request)
	default:
		log.Print("LoginHandler Unsupported request method", request.Method)
	}
}

func displayLoginPage(responseWriter http.ResponseWriter, request *http.Request) {
	executeTemplate(loginTemplate, responseWriter, nil)
}

func processLoginRequest(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	user := User{request.PostFormValue("username"), request.PostFormValue("password")}
	if isUserInDatabase(user) {
		setSession(user.Username, responseWriter)
		http.Redirect(responseWriter, request, "/menu", http.StatusFound)
	} else {
		fmt.Fprintf(responseWriter, "Wrong username or password")
	}
}
