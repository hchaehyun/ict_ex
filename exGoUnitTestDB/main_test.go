/*
// 초기 테스트 코드, main.go와의 코드중복성이 높음
package main_test


import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"          // http 패키지
	"net/http/httptest" // http test
	"testing"
)

func TestPing(t *testing.T) {
	// 테스트 대상 함수 호출, 테스트용 라우터 생성
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

	// 결과 검증, http 요청 테스트를 위해 요청 생성
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)

	// 해당 요청에 대한 테스트용 라우터 처리
	router.ServeHTTP(response, request)

	// 응답 검증
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}

	if response.Body.String() != `{"message":"pong"}` {
		t.Errorf("Expected message 'pong', got '%s'", response.Body.String())
	}
}
*/

// 초기 테스트 코드는 여러 개념을 하나의 함수에서 처리하고 있음, 보다 좋은 unit test를 위해 테스트 함수 분할 시도
// DB연결 테스트 함수, SQL 쿼리 실행 테스트 함수, 결과 처리 테스트 함수로 분할

package main_test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestPingDBConnect(t *testing.T) {
	// DB연결 테스트 함수
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		t.Errorf("DB 연결 실패: %v", err)
	}
	defer db.Close()
}

func TestPingQuery(t *testing.T) {
	// SQL 쿼리 실행 테스트 함수
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		t.Errorf("DB 연결 실패: %v", err)
	}
	defer db.Close()

	// SELECT 쿼리
	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	if err != nil {
		t.Errorf("SQL 쿼리 실행 실패: %v", err)
	}
}

func TestPingResult(t *testing.T) {
	// 결과 처리 테스트 함수
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		t.Errorf("DB 연결 실패: %v", err)
	}
	defer db.Close()

	// SELECT 쿼리
	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	if err != nil {
		t.Errorf("SQL 쿼리 실행 실패: %v", err)
	}

	// 결과 검증
	if name != "John" {
		t.Errorf("예상 결과와 실제 결과가 일치하지 않습니다: 예상 결과: John, 실제 결과: %s", name)
	}
}

func TestMain(m *testing.M) {
	// 테스트 실행 전에 DB 연결 설정
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 테스트 실행
	m.Run()
}
