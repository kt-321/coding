package main

import "fmt"

func main() {
	//バイトのスライス
	b := []byte{'a', 'b', 'c'}
	strb := string(b)
	fmt.Printf("%T: %v: %v\n", b, b, string(b)) //[]uint8: [97 98 99]: abc
	fmt.Printf("%T: %v: %v\n", strb, strb, string(strb)) //[]uint8: [97 98 99]: abc

	c := b[1:3] //インデックス1~2　スライスcはスライスbと同じメモリ領域
	fmt.Printf("%T: %v: %v\n", c, c, string(c)) //[]uint8: [98 99]: bc

	d := b[:2] //インデックス1~2　スライスcはスライスbと同じメモリ領域
	fmt.Printf("%T: %v: %v\n", d, d, string(d)) //[]uint8: [97 98]: ab

	e := d[1:] //インデックス1~2　スライスcはスライスbと同じメモリ領域
	fmt.Printf("%T: %v: %v\n", e, e, string(e)) //[]uint8: [98 99]: bc
}