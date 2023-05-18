# level1 #

在dao中创建一个新文件` file.go`，用于操作数据文件

```go
package dao

import (
	"encoding/csv"
	"os"
)

// 文件路径
const path = "D:\\Edc\\golandprojects\\go1.20.3\\gin-demo\\dao\\data\\users.csv"

// 从文件中加载所有用户数据
func loadUsers() (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = 2
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	users := make(map[string]string)
	for _, record := range records {
		users[record[0]] = record[1]
	}

	return users, nil
}

// 保存所有用户数据到文件
func saveUsers(users map[string]string) error {
	// 创建目录
	err := os.MkdirAll("./data", 0755)
	if err != nil {
		return err
	}

	// 创建或打开文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for username, password := range users {
		err := csvWriter.Write([]string{username, password})
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()

	return nil
}

```

并修改`dao\user.go`

```go
package dao

import "sync"

var (
	mutex    sync.Mutex
	database map[string]string
)

func init() {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	database = users
}

func AddUser(username, password string) error {
	mutex.Lock()
	defer mutex.Unlock()
	database[username] = password
	err := saveUsers(database)
	if err != nil {
		return err
	}

	return nil
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	mutex.Lock()
	defer mutex.Unlock()
	database, _ = Loadpassword(username)
	return database[username]
}
```

从而得以文件操作实现数据库

当注册成功时

![image-20230510231334690](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230510231334690.png)

文件中自动读入并储存数据

![image-20230510231500393](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230510231500393.png)

# level2 #

首先先在`router.go `中新加一条路由

```go
r.POST("/changepassword", changepassword)
```

在file.go中添加` UpdatePassword `函数，以供在` api/user.go`中使用

```go
func UpdatePassword(username string, newPassword string) error {
	// 打开用户数据文件，以读写方式打开
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	// 构造一个 *bufio.Reader 对象，以便读取文件中的数据
	reader := bufio.NewReader(f)

	// 读取文件中的每一行数据
	for {
		// 在文件中查找指定的用户名
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break // 已到文件末尾
		} else if err != nil {
			return err
		}
		fields := strings.Split(line, ",") // 将用户名和密码分离
		if strings.TrimSpace(fields[0]) == username {
			// 找到指定的用户，更新密码字段并写回文件
			fields[1] = newPassword
			_, err = f.Seek(-int64(len(line)), io.SeekCurrent)
			if err != nil {
				return err
			}
			_, err = f.WriteString(strings.Join(fields, ","))
			if err != nil {
				return err
			}
			return nil // 密码更新成功，返回 nil
		}
	}

	return fmt.Errorf("user not found: %s", username)
}
```

在在`api/user.go`中添加` changepassword `函数

```
func changepassword(c *gin.Context) {
	username := c.PostForm("username")
	oldpassword := c.PostForm("oldpassword")
	newpassword := c.PostForm("newpassword")
	// 验证用户是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}
	// 验证旧密码是否正确
	selectPassword := dao.SelectPasswordFromUsername(username)
	if selectPassword != oldpassword {
		utils.RespFail(c, "old password is incorrect")
		return
	}
	// 更新密码
	err := dao.UpdatePassword(username, newpassword)
	if err != nil {
		utils.RespFail(c, "unable to change password")
		return
	}
	// 成功更新密码，返回成功响应
	utils.RespSuccess(c, "password changed")
}
```



![image-20230511142834315](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230511142834315.png)

![image-20230511143251188](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230511143251188.png)

想法是输入用户名找回密码

新增路由`findpassword`

```
r.POST("/findpassword", findpassword)
```



`api/user.go`新增函数

```
func findpassword(c *gin.Context) {
	// 传入用户名和密码
	username := c.PostForm("username")
	selectPassword := dao.SelectPasswordFromUsername(username)
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	utils.RespSuccess(c, fmt.Sprintf("find successful,%s", selectPassword))
}
```



![image-20230511154921047](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230511154921047.png)

# level3 #

在`model/user.go`中添加`Message` 结构体

```go
type Message struct {
	Id        string `form:"Id" json:"id" binding:"required"`
	Text      string `form:"Text" json:"Text" binding:"required"`
	CreatedAt time.Time
}
```

在`api/user.go`中添加`AddMessage`函数接收数据

```go
func AddMessage(c *gin.Context) {
	Id := c.PostForm("id")
	messageText := c.PostForm("text")
	message := model.Message{
		Id:        Id,
		Text:      messageText,
		CreatedAt: time.Now(),
	}
	dao.AddMessage(message)
	utils.RespSuccess(c, "add message successful")
}
```

在`dao\user.go`中添加`AddMessage`函数将数据保存到文件当中

```go
func AddMessage(message model.Message) {
	file, err := os.OpenFile(path1, os.O_APPEND|os.O_WRONLY, 0600)#新建一个保存数据的文件路径为path1
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{message.Id, message.Text, message.CreatedAt.Format("2004-11-11 11:11:11")}); err != nil {
		log.Fatal(err)
	}
}
```

最后在`middleware/router.go`中增加新路由router

```
	r.POST("/message", AddMessage)
```



![image-20230511203037257](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230511203037257.png)

![image-20230511203106040](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230511203106040.png)
