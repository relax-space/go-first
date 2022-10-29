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
	fmt.Printf("buf: %v", s)
}
