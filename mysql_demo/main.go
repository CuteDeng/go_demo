package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/study_golang"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return nil
}

func queryOne(id int) (u user) {
	sql := `select id,name,age from user where id = ?;`
	db.QueryRow(sql, id).Scan(&u.id, &u.name, &u.age)
	return
}

func queryMore(id int) (users []user) {
	sql := `select id,name,age from user where id > ?;`
	rows, err := db.Query(sql, id)
	if err != nil {
		fmt.Println("db query err:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("rows scan err:", err)
		}
		users = append(users, u)
	}
	return
}

func insert() (id int64) {
	sql := `insert into user (name,age) values("jack",35);`
	res, err := db.Exec(sql)
	if err != nil {
		fmt.Println("db insert err:", err)
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		fmt.Println("db LastInsertId err:", err)
		return
	}
	return
}

func update() (n int64) {
	sql := `update user set name="tank" where id = 1;`
	res, err := db.Exec(sql)
	if err != nil {
		fmt.Println("db query err:", err)
		return
	}
	n, err = res.RowsAffected()
	if err != nil {
		fmt.Println("db RowsAffected err:", err)
		return
	}
	return
}

func prepareInsert() {
	sql := `insert into user (name,age) values(?,?);`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("db Prepare err:", err)
		return
	}
	var m = map[string]int{
		"jimmy": 25,
		"jason": 19,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init db err:", err)
		return
	}
	// u := queryOne(1)
	// fmt.Println(u)
	// users := queryMore(0)
	// id := insert()
	// n := update()
	// fmt.Println(n)
	prepareInsert()
}
