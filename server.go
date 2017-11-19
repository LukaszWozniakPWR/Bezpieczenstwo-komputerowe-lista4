package main

import (
	"net/http"
	"log"
	"./handlers"
)

const LOGIN_ENDPOINT = "/login"
const MENU_ENDPOINT = "/menu"
const LOGOUT_ENDPOINT = "/logout"
const TRANSACTION_ENDPOINT = "/transaction"
const TRANSACTIONS_ENDPOINT = "/transactions"
const CONFIRMATION_ENDPOINT = "/confirmation"
const PORT = ":9090"

func main() {
	http.HandleFunc(LOGIN_ENDPOINT, handlers.LoginHandler)
	http.HandleFunc(MENU_ENDPOINT, handlers.MenuHandler)
	http.HandleFunc(LOGOUT_ENDPOINT, handlers.LogoutHandler)
	http.HandleFunc(TRANSACTION_ENDPOINT, handlers.TransactionHandler)
	http.HandleFunc(TRANSACTIONS_ENDPOINT, handlers.TransactionsHandler)
	http.HandleFunc(CONFIRMATION_ENDPOINT, handlers.ConfirmationHandler)

	error := http.ListenAndServe(PORT, nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}
}
