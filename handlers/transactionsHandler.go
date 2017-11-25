package handlers

import "net/http"
import "../database"

func TransactionsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	user := getLoggedUserFromCookies(request)
	if user != nil {
		transactions := database.GetUsersTransactions(*user)
		executeTemplate(transactionsTemplate, responseWriter, transactions)
	} else {
		displayNotLoggedError(responseWriter)
	}
}
