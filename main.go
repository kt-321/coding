package main

import (
	"fmt"
	"math"
)

func main() {
	//浮動小数点数
	//var a float64 = 100
	//fmt.Printf("%T: %v", a, a)

	//円周率
	fmt.Println(math.Pi)

	//関数Dim x-y,マイナスの時は0を返す
	fmt.Println(math.Dim(10, 6)) //4
	fmt.Println(math.Dim(6, 10)) //0

	//絶対値
	test := math.Abs(-100.333)
	fmt.Println(test)

	//累乗
	fmt.Println(math.Pow(2, 3))

	//Pow10 は 10**n (n の 10 を底とする指数) を返します
	fmt.Println(math.Pow10(3)) //1000

	//関数Exp2 Exp2 は x の 2 を底とする指数である 2**x を返します。
	fmt.Println(math.Exp2(6)) //64

	//関数Round ゼロ以下を四捨五入。
	fmt.Println(math.Round(3.4)) //3
	fmt.Println(math.Round(3.5)) //4
	fmt.Println(math.Round(-3.4)) //-3
	fmt.Println(math.Round(-3.5)) //-4

	//関数Floor 引数の数字以下の最大の整数値を返す
	fmt.Println(math.Floor(11.6))
	fmt.Println(math.Floor(11.4))

	//関数Ceil 引数の数字以上の最小の整数値を返す
	fmt.Println(math.Ceil(12.22)) //13
	fmt.Println(math.Ceil(21.222)) //22

	//関数Sqrt 平方根(square root)
	fmt.Println(math.Sqrt(81)) //9

	//関数Cbrt 立方根（cubic root）
	fmt.Println(math.Cbrt(8)) //2

	var a bool = true
	fmt.Printf("%t", a)
}