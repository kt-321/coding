package main

import "fmt"

func main() {
	fmt.Print("test")

	m := map[string]int{
		"hoge1": 1,
		"hoge2": 2,
		"fuga1": 11,
		"fuga2": 12,
	}

	//map型のforループの書き方
	for k, v := range m {
		fmt.Printf("%T: %v, %T: %v\n", k, k, v, v)
	}

	fmt.Printf("%T: %v, %T: %v\n", m["hoge1"], m["hoge1"], m["hoge2"], m["hoge2"])

	//この書き方だと長さも容量も６になってしまう
	//長さ、容量共に3
	//s := make([]string, 3)
	//fmt.Println(s, len(s), cap(s))
	//s = append(s, "hoge1")
	//s = append(s, "non1")
	//s = append(s, "non2")


	//この書き方なら長さも容量も３
	s := []string{"hoge1", "none2", "none3"}

	fmt.Printf("%T: %v, len:%v, cap:%v\n", s, s, len(s), cap(s))

	for _, v := range s{
		//fmt.Printf("%T: %v, %T: %v\n", i, i, v, v)
		value, ok := m[v]
		if ok{
			fmt.Printf("%T: %v, %T: %v\n", value, value, ok, ok)
		} else {
			fmt.Printf("mapにキー「%v」とバリューのセットは存在しない\n", v)
		}
	}

	//TODO スライスに要素追加すると長さ・容量はどうなるか
}