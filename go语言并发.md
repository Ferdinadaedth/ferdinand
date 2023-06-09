# level1 #

### 并发中的资源竞争

#### 原子性

![image-20230508194218086](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230508194218086.png)

### 锁住共享资源

### 原子函数

![image-20230508195334679](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230508195334679.png)

### sync.Mutex(互斥锁) ###

![image-20230509125945870](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230509125945870.png)

### 使用通道做并发同步 ###

![image-20230509130302061](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230509130302061.png)

### 循环接收 ###

![image-20230509131329074](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230509131329074.png)

### 无缓冲的管道 ###

![image-20230509140937460](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230509140937460.png)

### 有缓存的通道 ###

![image-20230509152800054](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230509152800054.png)

### Select监控 ###

![image-20230509152914932](https://gitee.com/ferdinandaedth/ferdinand/raw/master/image-20230509152914932.png)

# level2 #

## time 库 ##

Go 语言中的 time 库提供了一个时间操作对象，其中包含了一些常见的时间操作方法。

**时间的表示形式**

使用 time 库可以方便地表示和操作时间相关的信息。time 库提供了三种时间表达形式：time.Time、time.Duration 和 time.Timer。

静态常量 `time.Now()` 可以获取到当前的时间。可以使用 `time.Parse()` 或 `time.ParseInLocation()` 方法将字符串转换为 time.Time 类型的时间结构体。对于常见的时间格式，可以使用 `time.Format()` 将时间转换为相应的字符串格式。

**常用时间操作**

常见的时间操作包括时间戳转换、时区设置、时间比较、日期计算等。

使用 `time.Unix()` 和 `time.Parse()` 可以将 Unix 时间戳和字符串转换为 time.Time 类型的时间结构体。

对于时区设置，使用 `time.LoadLocation()` 获取所需时区，然后使用 `time.Now().In()` 将时间转换为所需时区。

对于时间比较，可以使用比较运算符（`==`、`!=`、`<`）来比较两个时间。另外，还可以使用 `time.Before()` 和 `time.After()` 方法进行比较。

对于日期计算，可以使用 `time.Add()` 和 `time.Sub()` 方法。

# strings 库

Go 语言中的 strings 标准库提供了一些常见的字符串操作。

**字符串的表示和拼接**

在 Go 中，使用双引号引用的字符串字面量表示原始字符串内容。如果需要编写跨越多行的字符串，可以使用反引号。可以使用加号或 `strings.Join()` 方法将多个字符串拼接。

**字符串的查找和替换**

可以使用 `strings.Contains()` 检查字符串中是否包含子字符串，使用 `strings.Index()` 可以获取子字符串在字符串中的位置。如果需要替换字符串中的内容，可以使用 `strings.Replace()` 方法。

**字符串的分割和拼接**

可以使用 `strings.Split()` 将字符串拆分为字符串数组，使用 `strings.Join()` 可以将字符串数组拼接为一个字符串。

**字符串的处理和转换**

可以使用 `strings.ToLower()` 或 `strings.ToUpper()` 方法将字符串转换为相应的大小写格式；使用 `strings.Trim()`、`strings.TrimLeft()` 和 `strings.TrimRight()` 方法来删除字符串首尾的空格或指定的字符；使用 `strings.Fields()` 方法可以将字符串拆分成一个单词切片；使用 `strconv.ParseXXX()` 方法可以将字符串转换为其他类型的值。

# strconv 库

Go 语言中的 strconv 库用于类型转换和字符串转换。

**类型转换**

可以使用 `strconv.ParseXXX()` 将字符串转换为其他基本类型，其中 XXX 代表要转换的类型名称，如 `strconv.ParseFloat()`。

反过来，可以使用 `strconv.FormatXXX()` 将其他类型转换为字符串，其中 XXX 代表要转换的类型名称，如 `strconv.FormatInt()`。

**错误处理**

如果解析或格式化过程发生错误，则会在第二个结果中返回一个非 `nil` 的 `error` 类型。为了避免程序因此而崩溃，应该在使用返回值之前先检查该错误是否为空，如：

```
result, err := strconv.ParseInt("123", 10, 0)
if err != nil {
    // 处理错误
}
```



# json 库

Go 语言中，json 库提供了对 JSON（JavaScript Object Notation）数据格式的支持。

**JSON 的解析和序列化**

可以使用 `json.Marshal()` 和 `json.Unmarshal()` 将结构体和 JSON 格式之间进行转换。在使用 `json.Marshal()` 进行 JSON 格式的编码时，Go 语言会自动识别结构体中的可导出字段并进行编码。反之，使用 `json.Unmarshal()` 时，Go 语言会自动解码已编码的 JSON 数据为结构体。

**使用标签定制 JSON**

可以通过结构体字段上的标签指定 JSON 数据中的键名和其他编码选项。例如：

```
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

`Name` 字段的 JSON 编码键名将变为 `name`，`Age` 字段的 JSON 编码键名将变为 `age`。
