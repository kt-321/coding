package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (int, error) {
	for i:= range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	//r := strings.NewReader("Hello")
	//fmt.Printf("%T: %v\n", r, r)
	//
	//b := make([]byte, 8)
	//n, _ := r.Read(b)
	//fmt.Printf("%T: %v", n, n)

	//fmt.Printf("%T: %v", MyReader{}, MyReader{})

	reader.Validate(MyReader{})
}
