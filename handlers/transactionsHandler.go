package handlers

import "net/http"

func TransactionsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	username := getUserName(request)
	if username != "" {
		transactions := getUsersTransactions(username)
		executeTemplate(transactionsTemplate, responseWriter, transactions)
	} else {
		displayNotLoggedError(responseWriter)
	}
}
