package main_test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

/*
목표: Golang으로 연결된 DB에 CREATE, INSERT, SELECT, UPDATE, DELETE문을 이용한 쿼리문에 대한 Unit Test 실행

1. main.go에서 사용한 쿼리문이 잘 작동하는지 Unit test code 작성
2. 위 go 파일에서 사용한 쿼리문 모두 테스트 하기 위해서는 일정 순서 필요
 1. DB 연결 test
 2. new table CREATE test
 3. new data INSERT to new table test
 4. SELECT for check n print new data test
 5. UPDATE new data test
 6. SELECT for check UPDATE data test
 7. DELETE new data test
 8. SELECT for check DELETE new data test

*/
// 1. DB 연결 test
func TestPingDBConnect(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		t.Errorf("DB 연결 실패: %v", err)
	}
	defer db.Close()
}

// 2. new table CREATE test

// 3. new data INSERT to new table test

// 4. SELECT for check n print new data test

// 5. UPDATE new data test

// 6. SELECT for check UPDATE data test

// 7. DELETE new data test

// 8. SELECT for check DELETE new data test

// 테스트 실행 전에 DB 연결 설정
func TestMain(m *testing.M) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 테스트 실행
	m.Run()
}
