package database

import (
	"errors"
	"time"
)

const transactionsTableName = "transactions"

type Transaction struct {
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
	transaction.TransactionTime = time.Now().Format(time.RFC3339)
	err := IsTransactionCorrect(transaction)
	if err != nil {
		return err
	}
	executeStatement("INSERT INTO " + transactionsTableName + " (time, senderAccountNumber, receiverAccountNumber, amount, isApproved) VALUES (?,?,?,?,?)",
		transaction.TransactionTime,
		transaction.SenderAccountNumber,
		transaction.ReceiverAccountNumber,
		transaction.Amount,
		transaction.IsApproved)
	return nil
}

func IsTransactionCorrect(transaction Transaction) error {
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
	var id int
	var transaction Transaction
	var result []Transaction
	for rows.Next() {
		rows.Scan(
			&id,
			&transaction.TransactionTime,
			&transaction.SenderAccountNumber,
			&transaction.ReceiverAccountNumber,
			&transaction.Amount,
			&transaction.IsApproved)
		result = append(result, transaction)
	}
	return result
}
