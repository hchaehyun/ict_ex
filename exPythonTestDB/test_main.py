# 테스트 코드 부분
'''
1. 데이터베이스 연결 테스트
    pymysql.connect() 함수 정상적으로 호출되는지 확인
    연결 올바르게 되었는지 확인
2. 테이블 생성 테스트
    cur.execute() 함수 정상적으로 호출되는지 확인
    테이블 실제로 생성 되었는지 확인, 구조 예상대로 만들어졌는지 확인
3. 데이터 입력 테스트
    cur.execute() 함수 정상적으로 호출되는지 확인
    데이터 실제로 테이블에 입력되었는지 확인, 입력된 데이터 정확한지 확인
4. 데이터 조회 테스트
    cur.execute() 함수 정상적으로 호출되는지 확인
    입력했던 데이터 정상적으로 조회되는지 확인
4. 데이터 삭제 테스트
    cur.execute() 함수 정상적으로 호출되는지 확인
    데이터 실제로 삭제되었는지 확인
5. 테이블 삭제 테스트
    cur.execute() 함수 정상적으로 호출되는지 확인
    테이블 실제로 삭제되었는지 확인
6. 연결 종료 테스트
    conn.close() 함수 정상적으로 호출되는지 확인
    연결 실제로 종료되었는지 확인
'''
import unittest
import pymysql


class TestDatabase(unittest.TestCase):
    def setUp(self):
        # 연결 정보
        self.host = '127.0.0.1'
        self.user = 'hchaehyun'
        self.password = 'password'
        self.db = 'test'
        self.charset = 'utf8'

        # 연결 및 커서 생성, 테이블 생성
        self.conn = pymysql.connect(host=self.host, user=self.user, password=self.password, db=self.db, charset=self.charset)
        self.cur = self.conn.cursor()
        self.cur.execute("CREATE TABLE IF NOT EXISTS Lush (item VARCHAR(100) NULL)")

# 6. 연결 종료 테스트
    def tearDown(self):
        # 연결 종료
        self.cur.close()
        self.conn.close()

# 1. 데이터베이스 연결 테스트
    def testConnect(self):
        # 연결 확인
        self.assertIsNotNone(self.conn)

# 2. 테이블 생성 테스트
    def testCreateTable(self):
        # 테이블 만들기
        self.cur.execute("CREATE TABLE IF NOT EXISTS Lush (item VARCHAR(100) NULL)")

        # 테이블 존재 여부 확인
        self.cur.execute("SHOW TABLES LIKE 'Lush'")
        self.assertEqual(self.cur.fetchone()[0], 'Lush')

# 3. 데이터 입력 테스트
    def testInsertData(self):
        # 데이터 입력
        self.cur.execute("INSERT INTO Lush VALUES('intergalactic')")

        # 입력 데이터 확인
        self.cur.execute("SELECT * FROM Lush")
        for row in self.cur.fetchall():
            self.assertEqual(len(row), 1)

        self.conn.commit()

# 4. 데이터 조회 테스트
    def testSelectData(self):
        # 데이터 조회
        self.cur.execute("SELECT item FROM Lush")
        result = self.cur.fetchone()

        # 조회 결과 확인
        self.assertEqual(result[0], 'intergalactic')

# 5. 데이터 삭제 테스트
    def testDeleteData(self):
        # 데이터 삭제
        self.cur.execute("DELETE FROM Lush WHERE item = 'intergalactic'")

        # 삭제 결과 확인
        self.cur.execute("SELECT * FROM Lush WHERE item = 'intergalactic'")
        self.assertIsNone(self.cur.fetchone())

# 6. 테이블 삭제 테스트
    def testDropTable(self):
        # 테이블 삭제
        self.cur.execute("DROP TABLE IF EXISTS Lush")

        # 테이블 존재 여부 확인
        self.cur.execute("SHOW TABLES LIKE 'Lush'")
        self.assertIsNone(self.cur.fetchone())


if __name__ == '__main__':
    unittest.main()
