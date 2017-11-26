package handlers

import "net/http"
import (
	"../database"
	"strconv"
)

func ApproveHandler(responseWriter http.ResponseWriter, request *http.Request) {
	user := getLoggedUserFromCookies(request)
	if user == nil {
		displayNotLoggedError(responseWriter)
	} else if !user.Info.IsAdmin {
		displayNotAdminError(responseWriter)
	} else {
		processApprove(responseWriter, request)
	}
}

func processApprove(responseWriter http.ResponseWriter, request *http.Request) {
	id := request.FormValue("transactionid")
	atoiid, err := strconv.Atoi(id)
	if atoiid <=0 || err != nil {
		displayErrorMessage(responseWriter, err)
		return
	}
	database.ApproveTransaction(id)
	http.Redirect(responseWriter, request, "/approves", http.StatusFound)
}
