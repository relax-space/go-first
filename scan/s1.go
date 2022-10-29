package main

import (
	"fmt"
)

// scan 会等所有的输入完之后, 才会结束
func scan() {
	// 输入 1和空格
	var i1, i2 int
	fmt.Println("请输入:")
	fmt.Scan(&i1, &i2)
	fmt.Printf("scan: %v %v", i1, i2)

}

// scanln 只要按回车, 就结束
func scanln() {
	// 输入 2和空格
	var i3, i4 int
	fmt.Println("请输入:")
	fmt.Scanln(&i3, &i4)
	fmt.Printf("scanln: %v %v", i3, i4)

}
