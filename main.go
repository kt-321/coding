package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var r io.Reader

	fmt.Printf("%T: %v", r, r)

	r = strings.NewReader("Hello, world")
	// ここで生成されるstrings.Readerはio.Readerというinterfaceの定義を満たし、io.Readerとして扱うことができる

	fmt.Printf("%T: %v\n", r, r)

	content, err := ioutil.ReadAll(r) //引数にio.Reader(strings.Reader)をとる
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T: %v,\n %T: %v\n", content, content, string(content), string(content))


	filename := "go.mod"
	m := 10

	f, err := os.Open(filename)
	p := make([]byte, m)
	n, err := f.Read(p)
	if m < n {
		fmt.Printf("%dバイト読もうとしましたが、%dバイトしか読めませんでした\n", n, m)
	}
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(n)
}