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
	database, err := sqlx.Open("mysql", "root:root@tcp(192.168.1.201:3306)/golang")
	if err != nil {
		fmt.Println("open mysql failed", err)
	}
	Db = database
}

func insert() {

	// 开始事务
	conn, err := Db.Begin()

	// 一旦出现错误，就会滚回去
	defer func(){
		if err != nil{
			conn.Rollback()
		}
	}()
	
	if err != nil{
		fmt.Println("mysql conn failed, err:", err)
		return
	}


	sqlcmd := "insert into person(username, sex, email)values(?,?,?)"

	r, err := conn.Exec(sqlcmd, "stu02", "man", "stu02@mail.com",)
	if err != nil{
		fmt.Println("exec failed, err:", err)
		// conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil{
		fmt.Println("exec failed, err:", err)
		return
	}

    fmt.Println("insert success", id)
	r, err = conn.Exec(sqlcmd, "stu03", "man", "stu03@mail.com",)
	if err != nil{
		fmt.Println("exec failed, err:", err)
		// conn.Rollback()
		return
	}
	fmt.Println("insert success", id)
	conn.Commit()


}

func main() {
	insert()
}
