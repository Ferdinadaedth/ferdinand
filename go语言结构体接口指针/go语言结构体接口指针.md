# lv1 #

联想到了user struct，如下

```go
type User struct {
    ID          int64       
    Username    string
    Age         int64
    sex         string
    Email       string      
    CreatedAt   time.Time   
    UpdatedAt   time.Time   
}
```

# lv2 #

```go
package main
import (
	"fmt"
	"reflect"
)
type Printer interface {
	Print()
}
type User struct {
	UserName string
	Age      int
}
func (U User) Print() {
	fmt.Printf("Type: %s, Value: %v\n", reflect.TypeOf(U), U)
}

type Question struct {
	Id      int64
	content string
}
func (Q Question) Print() {
	fmt.Printf("Type: %s, Value: %v\n", reflect.TypeOf(Q), Q)
}
func main() {
	U := User{UserName: "彭于晏", Age: 41}
	Q := Question{Id: 2333, content: "abcd"}
	var printer Printer
	printer = U
	printer.Print()
	printer = Q
	printer.Print()
}

```

# lv3 #

```go
package main

import "fmt"

type Person struct {
	name       string
	level      int
	experience int
	health     int
	attack     int
}

type Attacker interface {
	Attack(target *Person)
}

func (P *Person) Attack(target *Person) {
	target.health -= P.attack
}
func main() {
	p1 := &Person{name: "顶针", level: 1, experience: 0, health: 100, attack: 10}
	p2 := &Person{name: "王源", level: 1, experience: 0, health: 100, attack: 6}
	var attacker Attacker
	attacker = p1
	fmt.Printf("%s的攻击力为 %d，%s的血量为%d\n", p1.name, p1.attack, p2.name, p2.health)
	attacker.Attack(p2)
	fmt.Printf("%s攻击%s，%s的血量为 %d\n", p1.name, p2.name, p2.name, p2.health)
}

```

