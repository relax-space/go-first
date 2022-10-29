# 输入

## scan 和 scanln

scan 和scanln 都无法通过一个字段读空格字符串

1. 相同点: 对于1个参数, 都是只读第一个空格前面的内容
``` go
package main

import "fmt"

func s1() {
	var i1 string
	fmt.Println("请输入:")
	fmt.Scan(&i1)
	fmt.Printf("scan: %v", i1)
}

func s2() {
	var i1 string
	fmt.Println("请输入:")
	fmt.Scanln(&i1)
	fmt.Printf("scanln: %v", i1)
}

func s3() {
	var i1 string
	fmt.Println("请输入:")
	fmt.Scanf("%v", &i1)
	fmt.Printf("scanf: %v", i1)
}

func main() {
	s1()
	// s2()
	// s3()
}

```

输出
``` shell
$ go run .
请输入:
1 2
scan: 1
$ go run .
请输入:
1 3
scanln: 1
$ go run .
请输入:
2 4
scanf: 2
```

2. 不同点: 对于多个参数,`scan`会等所有参数都输入结束, `scanln`不会 
```go
package main

import (
	"fmt"
)

// scan 会等所有的输入完之后, 才会结束
func scan() {
	// 输入 1和空格
	var i1, i2 int
	fmt.Scan(&i1, &i2)
	fmt.Printf("scan: %v %v", i1, i2)

}

// scanln 只要按回车, 就结束
func scanln() {
	// 输入 2和空格
	var i3, i4 int
	fmt.Scanln(&i3, &i4)
	fmt.Printf("scanln: %v %v", i3, i4)

}

func main() {
	scan()
	// scanln()
}

```

输出
``` shell
$ go run .
请输入:
1
2
scan: 1 2
$ go run .
请输入:
1
scanln: 1 0

```

## buf

读取空格字符串

``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buf() {
	fmt.Println("请输入:")
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Printf("buf: %v %v", s, len(s))
}

func main() {
	buf()
}

```

输出

``` shell
go run .
请输入:
1 2
buf: 1 2

```