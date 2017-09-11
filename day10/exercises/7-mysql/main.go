package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Persion struct {
	User_id  int    `db:user_id`
	Username string `db:username`
	Sex      string `db:sex`
	Email    string `db:email`
}

type Place struct {
	Country string `db:country`
	City    string `db:city`
	TelCode int    `db:telcode`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(192.168.11.128:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed", err)
	}
	Db = database
}

func insert() {

	p := Persion{
		Username: "stu001",
		Sex:      "man",
		Email:    "stu001@qq.com",
	}

	sqlcmd := "insert into person(username, sex, email)values(?,?,?)"

	r, err := Db.Exec(sqlcmd, p.Username, p.Sex, p.Email)
	if err != nil {
		fmt.Println("exec failed..")
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec read failed,", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func main() {
	insert()

}
