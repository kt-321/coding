package main

import (
	"fmt"
)

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


	//スライスこの書き方なら長さも容量も３
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


	//Arrays　スライスと違いサイズは固定。
	//可変長 長さ３になる
	a := [...]string{
		"test1",
		"test2",
		"hoge1",
	}
	fmt.Printf("%T: %v, %v, %v\n", a[0], a[0], len(a), a)

	//①配列をもとにスライスを作成。 [test1 test2]
	//スライスの最初の要素から数えて、元となる配列の要素数
	//startからend-1まで
	//長さ２で容量3となる。→　スライスの最初の要素から数えて、元となる配列の要素数なので容量３
	slice1 := a[0:2]
	fmt.Printf("%T, %v, %v, %v\n", slice1, len(slice1), cap(slice1), slice1)
	slice1[0] = "changed"
	fmt.Printf("%T, %v, %v, %v\n", slice1, len(slice1), cap(slice1), slice1)

	//a（配列）の値、サイズは変わってない
	fmt.Printf("%T: %v, %v\n", a, a, len(a))

	fmt.Println("++++++++++++++++++++")

	//②長さ1で容量2となる  [test2] →　スライスの最初の要素から数えて、元となる配列の要素数なので容量２
	slice2 := a[1:2]
	fmt.Printf("%T, %v, %v, %v\n", slice2, len(slice2), cap(slice2), slice2)


	//a（配列）の値、サイズは変わってない
	fmt.Printf("%T: %v, %v\n", a, a, len(a))

	//③長さ1で容量1となる →　スライスの最初の要素から数えて、元となる配列の要素数なので容量1
	slice3 := a[2:]
	fmt.Printf("%T, %v, %v, %v\n", slice3, len(slice3), cap(slice3), slice3)

	//a（配列）の値、サイズは変わってない
	fmt.Printf("%T: %v, %v\n", a, a, len(a))

	//④長さ2で容量3となる →　スライスの最初の要素から数えて、元となる配列の要素数なので容量3
	slice4 := a[:2]
	fmt.Printf("%T, %v, %v, %v\n", slice4, len(slice4), cap(slice4), slice4)

	//⑤長さ3で容量3となる → 最初から長さ・容量３のスライスと同じ
	slice5 := a[:]
	fmt.Printf("%T, %v, %v, %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Println("++++++++++++++++++++")

	//スライスの代入
	sl1 := []string{"hoge", "fuga"}
	var sl2 []string

	sl2 = sl1
	sl2[0] = "nyaho"

	fmt.Println(sl2)
	fmt.Println(sl1) //sl2の要素を帰るとsl1の値も変化する


	sl3 := []int{2, 4}
	var sl4 []int

	fmt.Printf("%T, %v, %v, %v\n", sl4, len(sl4), cap(sl4), sl4)
	fmt.Println(sl4)

	sl4 = sl3
	fmt.Printf("%T, %v, %v, %v\n", sl4, len(sl4), cap(sl4), sl4)
	fmt.Println(sl4)

	sl4[0] = 3 //sl3の値も変わる
	fmt.Printf("%T, %v, %v, %v\n", sl3, len(sl3), cap(sl3), sl3)
	fmt.Printf("%T, %v, %v, %v\n", sl4, len(sl4), cap(sl4), sl4)

	fmt.Println("====================")

	//スライスのゼロ値はnil 長さと容量は0
	var sli []int
	fmt.Println(sli)
	fmt.Printf("%T: %v, %v, %v\n", sli, sli, len(sli), cap(sli)) //[] 0 0

	if sli == nil {
		fmt.Println("sli == nil !!!")
	}

	fmt.Println("====================")

	//組み込み関数make()を用いたスライス作成.
	//それぞれの要素がゼロ値であるスライスが作成される
	test := make([]int, 3, 3)
	fmt.Println(test)
	// []int: [0 0 0], 3, 3
	fmt.Printf("%T: %v, %v, %v\n", test, test, len(test), cap(test))

	test2 := make([]map[string]int, 3)
	fmt.Println(test2)

	//mapの書き方
	test2[0] = map[string] int{
		"hoge": 1,
		"fuga": 2,
	}
	fmt.Printf("%T: %v, %v, %v\n", test2, test2, len(test2), cap(test2))

	maps := []map[string] int{
		map[string] int {"a": 1, "b": 2},
		map[string] int {"ai": 11, "ue": 22},
		map[string] int {"aii": 111, "uee": 222},
	}

	for k, v := range maps {
		//fmt.Println(k, v)
		test2[k] = v
		fmt.Printf("%T: %v, %v, %v\n", test2, test2, len(test2), cap(test2))
	}


}