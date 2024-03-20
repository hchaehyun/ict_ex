package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// sql.DB 객체 db 생성
	db, err := sql.Open("mysql", "hchaehyun:password@tcp(127.0.0.1:3306)/docker_root")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SELECT 쿼리
	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	//...(db 사용)....
}
