package handlers

import (
	"net/http"
	"log"
)

func ConfirmationHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		displayConfirmationPage(responseWriter, request)
	case "POST":
		processTransactionConfirmed(responseWriter, request)
	default:
		log.Print("LoginHandler Unsupported request method", request.Method)
	}
}

func displayConfirmationPage(responseWriter http.ResponseWriter, request *http.Request) {
	username := getUsernameFromCookies(request)
	if username == "" {
		displayNotLoggedError(responseWriter)
	} else {
		transaction, err := getPreparedTransactionFor(username)
		if err != nil {
			displayNotPreparedTransaction(responseWriter)
		} else {
			executeTemplate(confirmationTemplate, responseWriter, transaction)
		}
	}
}

func processTransactionConfirmed(responseWriter http.ResponseWriter, request *http.Request) {
	username := getUsernameFromCookies(request)
	if username == "" {
		displayNotLoggedError(responseWriter)
	} else {
		transaction, err := getPreparedTransactionFor(username)
		if err != nil {
			displayNotPreparedTransaction(responseWriter)
		} else {
			makeTransaction(transaction)
			http.Redirect(responseWriter, request, "/transactions", http.StatusFound)
		}
	}
}
