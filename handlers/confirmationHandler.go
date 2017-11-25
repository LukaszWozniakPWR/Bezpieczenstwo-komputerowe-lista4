package handlers

import (
	"net/http"
	"log"
	"../database"
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
	user := getLoggedUserFromCookies(request)
	if user == nil {
		displayNotLoggedError(responseWriter)
	} else {
		transaction := getTransactionFromCookies(request)
		if transaction == nil {
			displayNotPreparedTransaction(responseWriter)
		} else {
			executeTemplate(confirmationTemplate, responseWriter, *transaction)
		}
	}
}

func processTransactionConfirmed(responseWriter http.ResponseWriter, request *http.Request) {
	user := getLoggedUserFromCookies(request)
	if user == nil {
		displayNotLoggedError(responseWriter)
	} else {
		transaction := getTransactionFromCookies(request)
		if transaction == nil {
			displayNotPreparedTransaction(responseWriter)
		} else {
			database.MakeTransaction(*transaction)
			http.Redirect(responseWriter, request, "/transactions", http.StatusFound)
		}
	}
}
