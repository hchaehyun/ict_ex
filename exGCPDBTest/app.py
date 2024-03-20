"""
A sample Hello World server.
"""

import pymysql
import os

from flask import Flask, render_template


# pylint: disable=C0103
app = Flask(__name__)



@app.route('/')
def hello():
    # MySQL 기본 설정
    # 환경 변수에서 MySQL 연결 정보 가져오기
    host = 'host.docker.internal'
    port = 3306
    user = os.environ.get('DB_USER', 'default_user')  # 환경 변수가 없을 경우 기본값 'default_user' 사용
    password = os.environ.get('DB_PASS', 'default_password')  # 환경 변수가 없을 경우 기본값 'default_password' 사용
    database = os.environ.get('DB_NAME', 'default_database')  # 환경 변수가 없을 경우 기본값 'default_database' 사용

    # MySQL에 연결
    connection = pymysql.connect(host=host, port=port, user=user, password=password, database=database, charset='utf8')



    """Return a friendly HTTP greeting."""
    message = "It's running!"

    """Get Cloud Run environment variables."""
    service = os.environ.get('K_SERVICE', 'Unknown service')
    revision = os.environ.get('K_REVISION', 'Unknown revision')

    return render_template('index.html',
        message=message,
        Service=service,
        Revision=revision)

if __name__ == '__main__':
    server_port = os.environ.get('PORT', '8080')
    app.run(debug=False, port=server_port, host='0.0.0.0')
