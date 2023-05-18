# leve1 #

首先，在 MySQL 中创建一个 database 和一张数据表 student:
```sql
CREATE DATABASE testdb;
USE testdb;
CREATE TABLE student (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50),
  age INT,
  gender VARCHAR(10)
);
```



接下来，使用 Go 语言操作 MySQL，向 student 表中插入十条记录，并全部读出并打印:



```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:yx041110@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
// 插入十条记录
for i := 1; i &lt;= 10; i++ {
    name := fmt.Sprintf("student_%d", i)
    age := i + 18
    gender := "male"
    if i%2 == 0 {
        gender = "female"
    }

    _, err := db.Exec("INSERT INTO student (name, age, gender) VALUES (?, ?, ?)", name, age, gender)
    if err != nil {
        panic(err.Error())
    }
}

// 读取全部记录并打印
rows, err := db.Query("SELECT * FROM student")
if err != nil {
    panic(err.Error())
}
defer rows.Close()

fmt.Println("id\tname\tage\tgender")
for rows.Next() {
    var id int
    var name string
    var age int
    var gender string
    err = rows.Scan(&amp;id, &amp;name, &amp;age, &amp;gender)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("%d\t%s\t%d\t%s\n", id, name, age, gender)
}
```
}

结果

![image-20230516152556053](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230516152556053.png)

# level2 #

数据库中创建数据库`userdb`并插入一条用户名密码

```sql
CREATE DATABASE userdb;
use userdb;
CREATE TABLE user(                                                                username VARCHAR(50),                                                              password VARCHAR(50)                                                              );
INSERT INTO user (username, password) VALUES ('111', '111');
```

对之前作业中`dao/user.go`进行修改，将函数的数据库改为MySQL

```go
package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接信息
const (
	userName = "root"
	password = "yx041110"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "userdb"
)

// SelectUser 根据用户名查询用户是否存在
func SelectUser(username string) bool {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, ip, port, dbName))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 查询用户名是否存在
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM user WHERE username=?", username).Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	return count > 0
}

// AddUser 添加用户
func AddUser(username, password string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, ip, port, dbName))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 插入用户记录
	_, err = db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		panic(err.Error())
	}
}

// SelectPasswordFromUsername 根据用户名查询密码
func SelectPasswordFromUsername(username string) string {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, ip, port, dbName))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 查询密码
	var password string
	err = db.QueryRow("SELECT password FROM user WHERE username=?", username).Scan(&password)
	if err != nil {
		panic(err.Error())
	}
	return password
}

```

成功实现注册功能以及登录功能。

![image-20230516204138803](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230516204138803.png)

![image-20230516203226682](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230516203226682.png)

