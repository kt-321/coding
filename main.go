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

	//絶対値
	test := math.Abs(-100.333)
	fmt.Println(test)

	//累乗
	fmt.Println(math.Pow(2, 3))

	//平方根(square root)
	fmt.Println(math.Sqrt(81))

	var a bool = true
	fmt.Printf("%t", a)
}