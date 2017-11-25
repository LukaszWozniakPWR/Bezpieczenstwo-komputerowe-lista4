package handlers

import (
	"html/template"
	"net/http"
	"log"
)

const loginHtmlPath = "html/login/"
const loginTemplate = "login.html"
const menuHtmlPath = "html/menu/"
const menuNoAdminTemplate = "menunoadmin.html"
const menuAdminTemplate = "menuadmin.html"
const transactionHtmlPath = "html/transaction/"
const transactionTemplate = "transaction.html"
const confirmationHtmlPath = "html/confirmation/"
const confirmationTemplate = "confirmation.html"
const transactionsHtmlPath = "html/transactions/"
const transactionsTemplate = "transactions.html"
const mainHtmlPath = "html/main/"
const mainTemplate = "main.html"

var templates = template.Must(template.ParseFiles(
	loginHtmlPath+loginTemplate,
	menuHtmlPath+menuNoAdminTemplate,
	menuHtmlPath+menuAdminTemplate,
	transactionHtmlPath+transactionTemplate,
	confirmationHtmlPath+confirmationTemplate,
	mainHtmlPath+mainTemplate,
	transactionsHtmlPath+transactionsTemplate))

func executeTemplate(template string, responseWriter http.ResponseWriter, data interface{}) {

	error := templates.ExecuteTemplate(responseWriter, template, data)
	if error != nil {
		log.Print("Error Executing Template: ", template, " failed ", error)
	}
}
