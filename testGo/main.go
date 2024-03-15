package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
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

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// test
