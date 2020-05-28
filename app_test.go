package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)


func assert(o bool) {
	if !o {
		fmt.Println(__getRecentLine())
		os.Exit(1)
	}
}
func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}

func TestEncoder(t *testing.T) {

	assert(encode("qwrt") == "qwrt")
	assert(encode("QWRT") == "QWRT")
	assert(encode("aeiou") == "12345")
	assert(encode("AEIOU") == "12345")
	assert(encode("") == "")
	assert(encode("hello") == "h2ll4")
}

func TestDecoder(t *testing.T) {

	assert(decode("qwrt") == "qwrt")
	assert(decode("QWRT") == "QWRT")
	assert(decode("12345") == "aeiou")
	//assert(decode("12345") == "AEIOU") # What should we do with this case?
	assert(decode("") == "")
	assert(decode("h2ll4") == "hello")
}