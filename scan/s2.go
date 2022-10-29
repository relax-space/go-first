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
