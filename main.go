package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//doBuffer()

	//doEqual()

	//doReadFrom()
	//doReader()
	//doReader2()

	stringsReader()
}

func doBuffer() {
	//バイト列からバッファを作る
	buf := bytes.NewBuffer([]byte{1, 2, 3})
	//buf := bytes.NewBuffer([]byte{0x02, 2, 3})
	log.Println(buf.Bytes())

	buf.Write([]byte{4, 5, 6})

	//[]byteでバッファを作成
	b:= make([]byte, 3)
	log.Println(b)

	//bufの内容を作成したバッファに移す
	buf.Read(b)
	log.Println(b, buf.Bytes())

	//空っぽにできる
	buf.Reset()
	log.Println(buf.Bytes())
	log.Println(buf.Len())
	log.Println(buf.Cap())

}

func doEqual() {
	//バイト列からバッファを作る
	buf := bytes.NewBuffer([]byte{1, 2, 3})
	buf2 := bytes.NewBuffer([]byte{0x02, 2, 3})

	log.Println(buf.Len())
	log.Println(buf.Cap())

	//引数はバイトのスライス(Stringにキャストして比較してる)
	log.Println(bytes.Equal(buf.Bytes(), buf2.Bytes()))
}

func doReadFrom() {
	f, _ := os.Open("read.txt")
	//f, _ := os.Open(os.Args[1])

	buf := bytes.Buffer{}
	n, _ := buf.ReadFrom(f)
	log.Println(n)

	log.Println(buf)
	buf.Write([]byte{99, 100})

	buf.WriteTo(os.Stdout)
}

func doReader() {
	buf := bytes.NewReader([]byte{1, 2, 3})

	b := make([]byte, 3)

	buf.Read(b)

	log.Println(b, buf)
}


func doReader2() {
	buf := bytes.NewReader([]byte{1, 2, 3, 4})
	p := make([]byte, 3)

	n, err := buf.Read(p)
	log.Println(n)

	if n < 3 {
		//エラーが発生してもn>0バイト埋められている可能性がある
		log.Println("failed")
	}
	if err != nil {
		log.Println("error")
	}
}

func stringsReader() {
	str := strings.NewReader("testhoge!!!")
	log.Println(str)

	b := make([]byte, 3)

	for {
		n, err := str.Read(b)
		log.Println(string(b), len(b), n, err)
		log.Printf("b[:n] = %q\n", b[:n])
		log.Printf("b = %q\n", b)
		if err == io.EOF {
			log.Println("==========")
			break
		}

	}
}