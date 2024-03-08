package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
목표: Golang으로 연결된 DB에 CREATE, INSERT, SELECT, UPDATE, DELETE, DROP문을 이용한 쿼리문 작성
기존(exGoUnitTestDB.go)쿼리 실행은 Query() 함수를 이용해서 실행. (SELECT문만 사용했기 때문)
이번에는 Exec() 함수를 같이 이용해 쿼리문 실행 시도. (CREATE, INSERT, UPDATE, DELETE, DROP문을 사용하기 때문)

1. Exec(), Query() 함수를 이용해 CREATE, INSERT, SELECT, UPDATE, DELETE, DROP 쿼리문 사용
2. 위 쿼리문 모두 사용하기 위해서는 일정 순서 필요
 1. DB 연결
 2. new table CREATE / Exec() -> Lush table 생성하기
 3. new data INSERT to new table / Exec() -> item이 intergalatic인 data 삽입
 4. SELECT for check n print new data / Query()
 5. UPDATE new data / Exec() -> intergalatic을 goddess로 변경
 6. SELECT for check UPDATE data / Query()
 7. DELETE new data / Exec() -> goddess data 삭제
 8. SELECT for check DELETE new data / Query()
 9. DROP new table / Exec() -> Luch table 삭제


*/
// main 시작
func main() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		// 1. sql.DB 객체 생성(연결)
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
		// db 연결 실패한 경우, 실패 사유 에러 log에 찍기
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close() // db 다 사용하면 닫기

		// 2. new table CREATE
		// "Lush" 테이블 생성
		_, err = db.Exec("CREATE TABLE Lush (item VARCHAR(100) NULL);")
		if err != nil {
			log.Fatal(err)
		}

		// 3. new data INSERT to new table
		_, err = db.Exec("INSERT INTO Lush (item) VALUES ('intergalactic')")
		if err != nil {
			log.Fatal(err)
		}
		// 4. SELECT for check n print new data
		var insert_item string
		err = db.QueryRow("SELECT item FROM Lush").Scan(&insert_item)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("INSERT item: ", insert_item)

		// 5. UPDATE new data
		_, err = db.Exec("UPDATE Lush SET item='goddess' WHERE item='intergalactic'")
		if err != nil {
			log.Fatal(err)
		}

		// 6. SELECT for check UPDATE data
		var update_item string
		err = db.QueryRow("SELECT item FROM Lush").Scan(&update_item)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("UPDATE item: ", update_item)

		// 7. DELETE new data
		_, err = db.Exec("DELETE FROM Lush WHERE item='goddess'")
		if err != nil {
			log.Fatal(err)
		}

		// 8. SELECT for check DELETE new data
		var delete_item string
		err = db.QueryRow("SELECT item FROM Lush").Scan(&delete_item)
		if err != sql.ErrNoRows {
			log.Fatal("Deleted item still exists")
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DELETE item: ", delete_item)

		// 9. DROP new table
		_, err = db.Exec("DROP TABLE Lush;")
		if err != nil {
			log.Fatal(err)
		}

		// DB 연결 확인
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
