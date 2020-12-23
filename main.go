package main

import "fmt"

//クロージャは、それ自身の外部から変数を参照する関数値
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func subtract() func(int) int{
	res := 0
	return func (x int) int {
		res -= x
		return res
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i ++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	hoge, fuga := subtract(), subtract()
	for i := 0; i < 10; i++ {
		fmt.Println(hoge(i), fuga(-2*i))
	}

}
