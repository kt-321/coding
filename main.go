package main

import "fmt"

func main() {
	//interface{}型の変数を定義
	var i interface{}
	//i = 4
	//fmt.Println(i)
	//
	//i = 4.5
	//fmt.Println(i)

	i = "test"
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	checkType(i)


	//testはポインタ変数
	//var testest *string
	//× iはinterface{}型なので代入できない
	//testest = &i

	//interfaceで定義したものをstring型に戻す(if str, ok := i.(string): ok... のような形で型アサーション)
	//str, ok := i.(string)
	//fmt.Printf("%v :%v\n", str, ok)
	//
	//testest = &str
	//fmt.Println(testest)
	////間接参照モード
	//fmt.Println(*testest)
}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("intです")
	case string:
		fmt.Println("stringです")
	case float64:
		fmt.Println("float64です")
	default:
		fmt.Println("それ以外です")
	}
}