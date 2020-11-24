package main

import "fmt"

type Hoge struct {
	Label string
	Index int
	Name string
}

func main() {
	name := "hoge"
	fmt.Println(&name)

	var hoge *Hoge

	//Hogeへのポインタである*Hoge型の値を入れている
	hoge = &Hoge{
		Label: "test",
		Index: 0,
		Name: "testだよ",
	}

	fmt.Println(hoge)
	fmt.Println(*hoge)

	//++++++++++++++++++++++++++++++++

	hoge2 := Hoge{
		Label: "test2",
		Index: 1,
		Name: "test2です。",
	}

	fmt.Println(hoge2)

	h := hoge2
	h.Label = "change"
	fmt.Println(h)
	fmt.Println(hoge2)

	h2 := &hoge2
	h2.Label = "change"
	fmt.Println(*h2)
	fmt.Println(hoge2)


	//++++++++++++++++++++++++++++++++++
	//hoge3はポインタ変数
	var hoge3 *int

	fmt.Println(hoge3)

	sample := 2
	hoge3 = &sample

	//アドレスモード
	fmt.Println(hoge3)
	//間接参照モード
	fmt.Println(*hoge3)
}