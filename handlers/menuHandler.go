package handlers

import (
	"net/http"
)

func MenuHandler(responseWriter http.ResponseWriter, request *http.Request) {
	user := getLoggedUserFromCookies(request)
	if user != nil {
		if user.Info.IsAdmin {
			executeTemplate(menuAdminTemplate, responseWriter, user.Info)
		} else {
			executeTemplate(menuNoAdminTemplate, responseWriter, user.Info)
		}
	} else {
		displayNotLoggedError(responseWriter)
	}
}
