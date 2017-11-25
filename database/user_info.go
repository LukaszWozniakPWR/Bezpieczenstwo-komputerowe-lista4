package database

const userInfoTableName = "user_info"

type UserInfo struct {
	Name string
	AccountNumber string
	IsAdmin bool
}