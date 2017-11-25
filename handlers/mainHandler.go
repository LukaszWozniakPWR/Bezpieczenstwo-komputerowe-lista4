package handlers

import "net/http"

func MainHandler(responseWriter http.ResponseWriter, request *http.Request) {
	user := getLoggedUserFromCookies(request)
	if user != nil {
		http.Redirect(responseWriter, request, "/menu", http.StatusFound)
	} else {
		executeTemplate(mainTemplate, responseWriter, nil)
	}
}


