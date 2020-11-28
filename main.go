package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("ファイル作成")
	f, err := os.Create("hello-world.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("hello!")


	//①ioutil.ReadFile
	//ファイルを閉じるといったことを考える必要なし（そもそも Close メソッドを持つオブジェクトが登場しない）
	content, _ := ioutil.ReadFile("hello-world.txt")
	fmt.Printf("%T: %v\n", content, content)
	fmt.Println(string(content))

	//② os.Open→*os.File ioutil.ReadAll→バイトのスライス
	g, err := os.Open("hello-world.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T: %v\n", g, g) //*os.File os.Fileはインタフェースio.Readerを満たす

	contentg, _ := ioutil.ReadAll(g)
	fmt.Printf("%T: %v\n", contentg, contentg)
	fmt.Println(string(contentg))

	// ファイルを閉じる
	if err := g.Close(); err != nil {
		log.Fatal(err)
	}
	var r io.Reader

	fmt.Printf("%T: %v", r, r)

	r = strings.NewReader("Hello, world")
	// ここで生成されるstrings.Readerはio.Readerというinterfaceの定義を満たし、io.Readerとして扱うことができる

	fmt.Printf("%T: %v\n", r, r)

	content3, err := ioutil.ReadAll(r) //引数にio.Reader(strings.Reader)をとる
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T: %v,\n %T: %v\n", content3, content3, string(content3), string(content3))


	filename := "hello-world.txt"
	m := 10

	test, err := os.Open(filename)
	p := make([]byte, m)
	n, err := test.Read(p)
	if m < n {
		fmt.Printf("%dバイト読もうとしましたが、%dバイトしか読めませんでした\n", n, m)
	}
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(n)
}