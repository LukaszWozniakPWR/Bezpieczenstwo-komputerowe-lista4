package handlers

import (
	"fmt"
	"net/http"
)

func displayNotLoggedError(responseWriter http.ResponseWriter) {
	fmt.Fprintf(responseWriter, "Musisz być zalogowany aby przeglądać tę stronę \n")
}

func displayWrongReceiver(responseWriter http.ResponseWriter) {
	fmt.Fprintf(responseWriter, "Nie znaleziono podanego adresata przelewu \n")
}

func displayWrongAmount(responseWriter http.ResponseWriter) {
	fmt.Fprintf(responseWriter, "Pieniądze muszą być dodatnie \n")
}

func displayNotPreparedTransaction(responseWriter http.ResponseWriter) {
	fmt.Fprintf(responseWriter, "Nie znaleziono transakcji \n")
}

func displayErrorMessage(responseWriter http.ResponseWriter, err error) {
	fmt.Fprintf(responseWriter, "%s \n", err.Error())
}