package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	//mapをrangeでイテレーションすると実行するたびに異なる結果。乱択アルゴリズム
	for _, ip := range hosts{
		fmt.Printf("%v\n", ip)
	}
}