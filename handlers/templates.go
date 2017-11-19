package handlers

import (
	"html/template"
	"net/http"
	"log"
)

const loginHtmlPath = "html/login/"
const loginTemplate = "login.html"
const menuHtmlPath = "html/menu/"
const menuTemplate = "menu.html"
const transactionHtmlPath = "html/transaction/"
const transactionTemplate = "transaction.html"
const confirmationHtmlPath = "html/confirmation/"
const confirmationTemplate = "confirmation.html"
const transactionsHtmlPath = "html/transactions/"
const transactionsTemplate = "transactions.html"

var templates = template.Must(template.ParseFiles(
	loginHtmlPath+loginTemplate,
	menuHtmlPath+menuTemplate,
	transactionHtmlPath+transactionTemplate,
	confirmationHtmlPath+confirmationTemplate,
	transactionsHtmlPath+transactionsTemplate))

func executeTemplate(template string, responseWriter http.ResponseWriter, data interface{}) {

	error := templates.ExecuteTemplate(responseWriter, template, data)
	if error != nil {
		log.Print("Error Executing Template: ", template, " failed ", error)
	}
}
