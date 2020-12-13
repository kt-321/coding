package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error){
	n, err := rot.r.Read(b) //書き方
	for i, v := range b {
		switch {
		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
			b[i] += 13
		case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Printf("%T: %v\n", s, s)

	r := rot13Reader{s}
	fmt.Printf("%T: %v\n", r, r)

	io.Copy(os.Stdout, &r) //TODO 復習
}
