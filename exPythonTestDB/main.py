import pymysql
import unittest

# mysql 연결하기
conn = pymysql.connect(host='127.0.0.1', user='hchaehyun', password='password', db='test', charset='utf8')

# mysql 연결 종료하기
conn.close()

