package handlers

import (
	"net/http"
	"log"
	"strconv"
	"../database"
)

func TransactionHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		displayTransactionPage(responseWriter, request)
	case "POST":
		processTransaction(responseWriter, request)
	default:
		log.Print("LoginHandler Unsupported request method", request.Method)
	}
}

func displayTransactionPage(responseWriter http.ResponseWriter, request *http.Request) {
	user := getLoggedUserFromCookies(request)
	if user != nil {
		executeTemplate(transactionTemplate, responseWriter, nil)
	} else {
		displayNotLoggedError(responseWriter)
	}
}

func processTransaction(responseWriter http.ResponseWriter, request *http.Request) {
	sender := getLoggedUserFromCookies(request)
	receiverAccountNumber := request.PostFormValue("receiverAccountNumber")
	amount, err := strconv.Atoi(request.PostFormValue("amount"))

	if sender == nil {
		displayNotLoggedError(responseWriter)
	} else if err != nil {
		displayWrongAmount(responseWriter)
	} else {
		transaction := database.Transaction{"", "", sender.Info.AccountNumber, receiverAccountNumber, amount, false}
		err = database.IsTransactionCorrect(transaction)
		if err != nil {
			displayErrorMessage(responseWriter, err)
		} else {
			clearSession(responseWriter)
			setSession(sender, &transaction, responseWriter)
			http.Redirect(responseWriter, request, "/confirmation", http.StatusFound)
		}
	}
}


