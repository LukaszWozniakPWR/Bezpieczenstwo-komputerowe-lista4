package database

import (
	"errors"
	"time"
	"strconv"
)

const transactionsTableName = "transactions"

type Transaction struct {
	Id	string
	TransactionTime       string
	SenderAccountNumber   string
	ReceiverAccountNumber string
	Amount                int
	IsApproved            bool
}

func createTransactionsTable() {
	executeStatement("CREATE TABLE IF NOT EXISTS " + transactionsTableName +
		" (id INTEGER PRIMARY KEY AUTOINCREMENT, time TEXT, senderAccountNumber TEXT, receiverAccountNumber TEXT, amount INTEGER, isApproved INTEGER)")
}

func deleteTransactionsTable() {
	executeStatement("DROP TABLE IF EXISTS " +  transactionsTableName)
}

func MakeTransaction(transaction Transaction) error {
	transaction.TransactionTime = time.Now().Format(time.Stamp)
	executeStatement("INSERT INTO " + transactionsTableName + " (time, senderAccountNumber, receiverAccountNumber, amount, isApproved) VALUES (\"" +
		transaction.TransactionTime + "\",\"" +
		transaction.SenderAccountNumber + "\",\"" +
		transaction.ReceiverAccountNumber + "\"," +
		strconv.Itoa(transaction.Amount) + "," + "0)")

	//executeStatement("INSERT INTO " + transactionsTableName + " (time, senderAccountNumber, receiverAccountNumber, amount, isApproved) VALUES (?,?,?,?,?)",
	//	transaction.TransactionTime,
	//	transaction.SenderAccountNumber,
	//	transaction.ReceiverAccountNumber,
	//	transaction.Amount,
	//	transaction.IsApproved)
	return nil
}

func IsTransactionCorrect(transaction Transaction) error {
	return nil // removed checks
	if !IsAccountNumberInDatabase(transaction.SenderAccountNumber) {
		return errors.New("Sender Account Number Incorrect")
	}
	if !IsAccountNumberInDatabase(transaction.ReceiverAccountNumber) {
		return errors.New("Receiver Account Number Incorrect")
	}
	if transaction.Amount <= 0 {
		return errors.New("Amount must be positive")
	}
	return nil
}

func GetUsersTransactions(user User) []Transaction {
	rows := queryDatabase("SELECT * from " + transactionsTableName+
		" WHERE senderAccountNumber LIKE " + user.Info.AccountNumber +
			" OR receiverAccountNumber LIKE " + user.Info.AccountNumber)
	var transaction Transaction
	var result []Transaction
	for rows.Next() {
		rows.Scan(
			&transaction.Id,
			&transaction.TransactionTime,
			&transaction.SenderAccountNumber,
			&transaction.ReceiverAccountNumber,
			&transaction.Amount,
			&transaction.IsApproved)
		result = append(result, transaction)
	}
	return result
}

func GetAllUnapprovedTransactions() []Transaction {
	rows := queryDatabase("SELECT * from " + transactionsTableName+
		" WHERE isApproved = 0")
	var transaction Transaction
	var result []Transaction
	for rows.Next() {
		rows.Scan(
			&transaction.Id,
			&transaction.TransactionTime,
			&transaction.SenderAccountNumber,
			&transaction.ReceiverAccountNumber,
			&transaction.Amount,
			&transaction.IsApproved)
		result = append(result, transaction)
	}
	return result
}

func ApproveTransaction(id string) {
	executeStatement("UPDATE " + transactionsTableName + " SET isApproved = 1 WHERE id = " + id)
}