package handlers

import (
	"net/http"
	"log"
	"fmt"
	"../database"
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
	user := database.LoginWithCredentials(
		database.UserCredential{
			request.PostFormValue("username"),
			request.PostFormValue("password")})
	if user != nil {
		setSession(user, nil, responseWriter)
		http.Redirect(responseWriter, request, "/menu", http.StatusFound)
	} else {
		fmt.Fprintf(responseWriter, "Wrong username or password")
	}
}
