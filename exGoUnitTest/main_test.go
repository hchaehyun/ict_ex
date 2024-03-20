/*
// 코드 1번의 테스트코드
// 기본 testing 이용해 작성한 테스트 코드, if문이 많이 사용된 모습을 확인할 수 있다.

package main

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

type subTest struct {
	arg1, arg2, expected int
}

var subTests = []subTest{
	subTest{5, 2, 3},
	subTest{10, 12, -2},
	subTest{5, 5, 0},
}

func TestSub(t *testing.T) {
	for _, test := range subTests {
		if output := Sub(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
*/

/*
// 코드 1번의 테스트코드
// 위 testing을 사용한코드를 개선한 버전, testify에서 제공한 assert를 사용
package main

import (
	"github.com/stretchr/testify/assert" // go get github.com/stretchr/testify/assert 로 설치 필요
	"testing"
)

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {
	a := assert.New(t)
	for _, test := range addTests {
		a.Equal(test.expected, Add(test.arg1, test.arg2), "they should be equal")
	}
}

type subTest struct {
	arg1, arg2, expected int
}

var subTests = []subTest{
	subTest{5, 2, 3},
	subTest{10, 12, -2},
	subTest{5, 5, 0},
}

func TestSub(t *testing.T) {
	a := assert.New(t)
	for _, test := range subTests {
		a.Equal(test.expected, Sub(test.arg1, test.arg2), "they should be equal")
	}
}
*/

/*
// 코드 2번의 테스트 코드
package table

import (
	"fmt"
	"testing"
)

var table = []struct {
	x    int
	y    int
	want int
}{
	{2, 2, 4},
	{5, 3, 8},
	{8, 4, 12},
	{12, 5, 17},
}

// Table Test (Normal)
func TestSum(t *testing.T) {
	for _, row := range table {
		got := Sum(row.x, row.y)
		if got != row.want {
			t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
		}
	}
}

// Table Test (With Subtest)
func TestSumSubtest(t *testing.T) {
	for _, row := range table {
		testName := fmt.Sprintf("Test %d+%d", row.x, row.y)
		t.Run(testName, func(t *testing.T) {
			got := Sum(row.x, row.y)
			if got != row.want {
				t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
			}
		})
	}
}
*/

// 코드 3번의 직접 작성한 테스트 코드
package main

import "testing" // 테스트 위해서 testing import

func TestUnitEx(t *testing.T) { // test 함수는 prefix로 Test 필요!
	// test case ~
	var testCase = []struct {
		a, b, expected string
	}{
		{"Hello", " Go", "Hello Go"},
		{"Always carry", " a lucky charm!", "Always carry a lucky charm!"},
		{"1999 0", "1 14", "1999 01 14"},
	}

	// 테스트 케이스 반복 실행
	for _, tc := range testCase {
		actual := UnitEx(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("테스트 실패: a=%s, b=%s, 예상 결과=%s, 실제 결과=%s", tc.a, tc.b, tc.expected, actual)
		}
	}
}
