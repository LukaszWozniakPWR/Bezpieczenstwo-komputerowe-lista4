package handlers

import (
	"net/http"
)

func MenuHandler(responseWriter http.ResponseWriter, request *http.Request) {
	username := getUserName(request)
	if username != "" {
		executeTemplate(menuTemplate, responseWriter, username)
	} else {
		displayNotLoggedError(responseWriter)
	}
}

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
