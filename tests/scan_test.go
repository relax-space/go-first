package tests

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 引用: https://echorand.me/posts/go-test-stdin/

// go test -v -run "^TestInput$"
func TestInput(t *testing.T) {
	exp := "jane"
	r := strings.NewReader(exp)
	scanner := bufio.NewScanner(r)
	msg := "Your name please? Press the Enter key when done"
	fmt.Fprintln(os.Stdout, msg)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	name := scanner.Text()
	if len(name) == 0 {
		t.Log("empty input")
	}
	t.Logf("You entered: %s\n", name)
	assert.Equal(t, exp, name, "scan error")

}

// go test -v -run TestInput2
func TestInput2(t *testing.T) {
	exp := "jane"
	r1 := strings.NewReader(exp)
	r := bufio.NewReader(r1)
	s, _ := r.ReadString('\n')
	assert.Equal(t, exp, s)
	assert.NotEqual(t, "12", s)
}
