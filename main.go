package main

import (
	"fmt"
	"strconv"
	"time"
)

type Test interface {
	M()
}

//19
type MyError struct {
	When time.Time
	What string
}

func (m *MyError) Error() string{
	return fmt.Sprintf("at %v, %v", m.When, m.What)
}

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

	//var t Test // nilインタフェース
	var t interface{}

	fmt.Printf("%T: %v\n", t, t) // 「<nil>: <nil>」
	//describe(t) //「<nil>: <nil>」
	//t.M() //「呼び出す 具体的な メソッドを示す型がインターフェースのタプル内に存在しないため、 nil インターフェースのメソッドを呼び出すと、ランタイムエラーになります。


	//19 A Tour of Go
	if err := run(); err != nil {
		fmt.Println(err)
	}

	//演習
	exercise()

	//型変換の演習
	exchangeType()
}

func describe(t Test) {
	fmt.Printf("%T: %v\n", t, t)
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

// Errorメソッドがなく関数runだけだと、「Errorメソッドがないためにerrorが実装できない」というエラー
func run() error {
	var hoge = &MyError{
		time.Now(),
		"this is test",
	}
	return hoge

	//return &MyError{
	//	time.Now(),
	//	"this is test.",
	//}
}

func exercise() {
	var i interface{} //空のインタフェース
	fmt.Printf("%T: %v\n", i, i)

	//i = 633
	i = "test"

	//i = MyError{
	//	time.Now(),
	//	"test",
	//}
	fmt.Printf("%T: %v\n", i, i)

	//ポインタ変数
	//var s *string

	//iはinterface型なので代入できない
	//s = &i

	//interface{}型で宣言された変数の値の具体的な型により処理を分岐 Type Switch
	switch i.(type) {
	//cf. uintは使えないっぽい
	case int:
		fmt.Println("iはint")
	case string:
		fmt.Println("iはstr")
		fmt.Printf("%T: %v\n", i, i)
		fmt.Printf("%T: %v\n", i.(string), i.(string))
	case float64:
		fmt.Println("iはfloat64")
	case MyError:
		fmt.Println("iはMyError型")
	//空のインタフェースの具体的な型はnil
	case nil:
		fmt.Println("iはnil")
	default:
		fmt.Println("iはそれ以外")
	}
}

func exchangeType() {
	//intからstringへの型変換 strconv.Itoa
	i := 1
	var str string

	str = strconv.Itoa(i)
	fmt.Printf("%T: %v\n", str, str)

	//△
	//var str2 string
	//str2 = string(i)
	//fmt.Printf("%T: %v\n", str2, str2) // 123がrune(Unicode)として認識され "{" になる



	//stringからintへの型変換 strconv.AtoI
	s := "111"

	//s := "11g" //以下のエラーハンドリングにて「strconv.Atoi: parsing "11g": invalid syntax」と出る。
	//fmt.Printf実行で「int: 0」

	var err error
	it, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T: %v\n", it, it)


	//stringからboolへ変換 strconv.ParseBool
	strtest := "true"
	stringToBool, err := strconv.ParseBool(strtest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T: %v\n", stringToBool, stringToBool)



	//boolからstringへ型変換 strconv.FormatBool
	var b bool = false
	//b := false  //この書き方でも良い

	boolToString := strconv.FormatBool(b)
	fmt.Printf("%T: %v\n", boolToString, boolToString)
}
