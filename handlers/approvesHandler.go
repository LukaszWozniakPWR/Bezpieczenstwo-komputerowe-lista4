package handlers

import (
	"net/http"
	"../database"
)

func ApprovesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	user := getLoggedUserFromCookies(request)
	if user == nil {
		displayNotLoggedError(responseWriter)
	} else if !user.Info.IsAdmin {
		displayNotAdminError(responseWriter)
	} else {
		displayApprovesPage(responseWriter)
	}
}

func displayApprovesPage(responseWriter http.ResponseWriter) {
	transactions := database.GetAllUnapprovedTransactions()
	executeTemplate(approvesTemplate, responseWriter, transactions)
}