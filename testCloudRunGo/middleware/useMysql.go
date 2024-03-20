package middleware

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
)

var db *sql.DB

func init() {
	// https://cloud.google.com/sql/docs/mysql/samples/cloud-sql-mysql-databasesql-connect-connector?hl=ko
	// 환경 변수에서 데이터베이스 연결 정보 읽기
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
	usePrivate := os.Getenv("PRIVATE_IP")

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		log.Fatalf("cloudsqlconn.NewDialer: %v", err)
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(/%s)/%s?parseTime=true",
		dbUser, dbPassword, instanceConnectionName, dbName)

	db, err = sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	// 연결 풀 설정
	db.SetMaxOpenConns(300)
	db.SetMaxIdleConns(300)

	// 연결 확인
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func UseMySql() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
