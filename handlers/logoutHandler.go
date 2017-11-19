package handlers

import (
	"net/http"
)

func LogoutHandler(responseWriter http.ResponseWriter, request *http.Request) {
	clearSession(responseWriter)
	http.Redirect(responseWriter, request, "/login", http.StatusFound)
}
