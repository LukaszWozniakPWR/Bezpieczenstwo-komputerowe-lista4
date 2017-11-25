package database

type User struct {
	Id int
	Credentials ProtectedUserCredential
	Info UserInfo
}

const usersTableName = "users"

func createUserCredentialsTable() {
	executeStatement("CREATE TABLE IF NOT EXISTS " + usersTableName +
		" (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT, name TEXT, accountNumber TEXT, isAdmin INTEGER)")
}

func deleteUserCredetialsTable() {
	executeStatement("DROP TABLE IF EXISTS " +  usersTableName)
}

func CreateNewUser(userCredential UserCredential, userInfo UserInfo) {
	protectedUser := userCredential.ProtectUser()
	executeStatement("INSERT INTO " + usersTableName + " (username, password, name, accountNumber, isAdmin) VALUES (?,?,?,?,?)",
		string(protectedUser.ProtectedUsername),
			string(protectedUser.ProtectedPassword),
				userInfo.Name,
					userInfo.AccountNumber,
						userInfo.IsAdmin)
}

func IsAccountNumberInDatabase(accountNumber string) bool {
	rows := queryDatabase("SELECT * FROM " + usersTableName + " WHERE accountNumber = " + accountNumber)
	for rows.Next() {
		rows.Close()
		return true
	}
	rows.Close()
	return false
}
