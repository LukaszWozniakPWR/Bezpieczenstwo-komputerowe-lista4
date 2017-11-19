package handlers

import (
	"net/http"
	"log"
	"strconv"
	."../transactions"
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
	username := getUsernameFromCookies(request)
	if username != "" {
		executeTemplate(transactionTemplate, responseWriter, nil)
	} else {
		displayNotLoggedError(responseWriter)
	}
}

func processTransaction(responseWriter http.ResponseWriter, request *http.Request) {
	sender := getUsernameFromCookies(request)
	receiver := request.PostFormValue("receiver")
	amount, err := strconv.Atoi(request.PostFormValue("amount"))

	if sender == "" {
		displayNotLoggedError(responseWriter)
	} else if !isUsernameInDatabase(receiver) {
		displayWrongReceiver(responseWriter)
	} else if amount <= 0 || err != nil {
		displayWrongAmount(responseWriter)
	} else {
		transaction := Transaction{"", sender, receiver, amount}
		prepareTransaction(transaction)
		http.Redirect(responseWriter, request, "/confirmation", http.StatusFound)
	}
}


