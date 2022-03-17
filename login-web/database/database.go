package database

import (
	"context"
	"log"
	"login-web/ent"
	"login-web/ent/user"

	_ "github.com/go-sql-driver/mysql"
)

type AccountInfo struct {
	Code int
	Id   string
	Name string
	Pwd  string
}

func InitAccount(driver, source string) []AccountInfo {
	info := []AccountInfo{}

	client, oErr := ent.Open(driver, source)
	checkErr(oErr)

	ctx := context.Background()

	account, qErr := client.User.Query().Select(user.FieldID, user.FieldPwd).All(ctx)
	checkErr(qErr)

	for _, a := range account {
		info = append(info, AccountInfo{
			Id:  a.ID,
			Pwd: a.Pwd,
		})
	}

	return info
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
