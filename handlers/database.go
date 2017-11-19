package handlers

import (
	"io/ioutil"
	"log"
	"encoding/json"
	."../users"
	."../transactions"
	"time"
	"os"
)

func isUserInDatabase(user User) bool {
	for  _, protectedUser := range readUsersFromDatabase() {
		if user.IsEqualTo(protectedUser) {
			return true
		}
	}
	return false
}

func isUsernameInDatabase(username string) bool {
	for  _, protectedUser := range readUsersFromDatabase() {
		if protectedUser.HasUsername(username) {
			return true
		}
	}
	return false
}

func readUsersFromDatabase() []ProtectedUser {
	data, err := ioutil.ReadFile("users/users.json")
	if err != nil {
		log.Fatal("Couldnt read users data")
	}
	var protectedUsers []ProtectedUser
	err = json.Unmarshal(data, &protectedUsers)
	if err != nil {
		log.Fatal("Couldnt convert users data to objects")
	}
	return protectedUsers
}

func prepareTransaction(transaction Transaction) {
	jsonTransaction, err := json.Marshal(transaction)
	if err != nil {
		log.Fatal("Error converting transaction to json")
	}
	err = ioutil.WriteFile("transactions/prepared"+transaction.Sender+".json", jsonTransaction, 0644)
	if err != nil {
		log.Fatal("Error writing to file")
	}
}

func getPreparedTransactionFor(username string) (Transaction, error) {
	data, err := ioutil.ReadFile("transactions/prepared"+username+".json")
	if err != nil {
		return Transaction{}, err
	}
	var transaction Transaction
	err = json.Unmarshal(data, &transaction)
	if err != nil {
		return Transaction{}, err
	} else {
		return transaction, nil
	}
}

func makeTransaction(transaction Transaction) {
	transaction.TransactionTime = time.Now().Format(time.RFC3339)
	os.Remove("transactions/prepared"+transaction.Sender+".json")
	jsonTransactions, err := json.Marshal([]Transaction{transaction})
	if err != nil {
		log.Fatal("Error converting transaction to json")
	}
	senderTransactionsFilePath := "transactions/"+transaction.Sender+".json"
	receiverTransactionsFilePath := "transactions/"+transaction.Receiver+".json"
	data, err := ioutil.ReadFile(senderTransactionsFilePath)
	if err != nil {
		err = ioutil.WriteFile(senderTransactionsFilePath, jsonTransactions, 0644)
	} else {
		var transactions []Transaction
		err = json.Unmarshal(data, &transactions)
		if err != nil {
			log.Fatal("Error converting transactions data to objects")
		}
		transactions = append(transactions, transaction)
		jsonTransactions, err := json.Marshal(transactions)
		if err != nil {
			log.Fatal("Error converting transactions to json")
		}
		ioutil.WriteFile(senderTransactionsFilePath, jsonTransactions, 0644)
	}

	data, err = ioutil.ReadFile(receiverTransactionsFilePath)
	if err != nil {
		err = ioutil.WriteFile(receiverTransactionsFilePath, jsonTransactions, 0644)
	} else {
		var transactions []Transaction
		err = json.Unmarshal(data, &transactions)
		if err != nil {
			log.Fatal("Error converting transactions data to objects")
		}
		transactions = append(transactions, transaction)
		jsonTransactions, err := json.Marshal(transactions)
		if err != nil {
			log.Fatal("Error converting transactions to json")
		}
		ioutil.WriteFile(receiverTransactionsFilePath, jsonTransactions, 0644)
	}
}

func getUsersTransactions(username string) []Transaction {
	data, err := ioutil.ReadFile("transactions/"+username+".json")
	if err != nil {
		return []Transaction{}
	} else {
		var transactions []Transaction
		err = json.Unmarshal(data, &transactions)
		if err != nil {
			log.Fatal("Error converting transactions data to objects")
		}
		return transactions
	}
}
